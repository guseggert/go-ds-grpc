// +build ignore
//go:build ignore

package main

import (
	"fmt"
	"math"
	"os"

	. "github.com/dave/jennifer/jen"
	pb "github.com/guseggert/go-ds-grpc/proto"
)

type feature struct {
	pbVal       pb.FeaturesResponse_Feature
	typeName    string
	emitTypeDef func(*File)
}

var features = []feature{
	{
		pbVal:    pb.FeaturesResponse_BATCHING,
		typeName: "batching",
		emitTypeDef: func(f *File) {
			f.Type().Id("batching").Interface(
				Id("Batch").Params(Qual("context", "Context")).Parens(List(Qual(dsPkg, "Batch"), Error())),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_CHECKED,
		typeName: "checked",
		emitTypeDef: func(f *File) {
			f.Type().Id("checked").Interface(
				Id("Check").Params(Qual("context", "Context")).Error(),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_SCRUBBED,
		typeName: "scrubbed",
		emitTypeDef: func(f *File) {
			f.Type().Id("scrubbed").Interface(
				Id("Scrub").Params(Qual("context", "Context")).Error(),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_GC,
		typeName: "gc",
		emitTypeDef: func(f *File) {
			f.Type().Id("gc").Interface(
				Id("CollectGarbage").Params(Qual("context", "Context")).Error(),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_PERSISTENT,
		typeName: "persistent",
		emitTypeDef: func(f *File) {
			f.Type().Id("persistent").Interface(
				Id("DiskUsage").Params(Qual("context", "Context")).Parens(List(Uint64(), Error())),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_TTL,
		typeName: "ttl",
		emitTypeDef: func(f *File) {
			f.Type().Id("ttl").Interface(
				Id("PutWithTTL").Params(Qual("context", "Context"), Qual(dsPkg, "Key"), Index().Id("byte"), Qual("time", "Duration")).Error(),
				Id("SetTTL").Params(Qual("context", "Context"), Qual(dsPkg, "Key"), Qual("time", "Duration")).Error(),
				Id("GetExpiration").Params(Qual("context", "Context"), Qual(dsPkg, "Key")).Parens(List(Qual("time", "Time"), Error())),
			)
		},
	},
	{
		pbVal:    pb.FeaturesResponse_TRANSACTION,
		typeName: "transaction",
		emitTypeDef: func(f *File) {
			f.Type().Id("transaction").Interface(
				Id("NewTransaction").Params(Qual("context", "Context"), Bool()).Parens(List(Qual(dsPkg, "Txn"), Error())),
			)
		},
	},
}

// maps from the protobuf enum to the corresponding interface name
var pbFeatureToTypeName = map[pb.FeaturesResponse_Feature]string{}

func init() {
	for _, feat := range features {
		pbFeatureToTypeName[feat.pbVal] = feat.typeName
	}
}

var (
	dsPkg = "github.com/ipfs/go-datastore"
	pbPkg = "github.com/guseggert/go-ds-grpc/proto"
)

func main() {
	// iterate though all 128 combinations of features and generate a struct for each combination
	// the name of each struct is deterministic, by using a consistent sorting of features and
	// mapping those bitwise into an integer.
	//
	// then when a client is started, it fetches the list of features from the server and picks
	// the correct struct that exactl implements those set of features, so that interface assertions
	// match
	f := NewFile("grpcds")

	numStructs := int(math.Pow(2, float64(len(features))))

	ctorEntries := Dict{}
	for i := 0; i < numStructs; i++ {
		num := i
		fields := []Code{Qual(dsPkg, "Datastore")}
		assignments := Dict{Id("Datastore"): Id("dstore")}
		for bit := 0; bit < len(features)+1; bit++ {
			if ((num >> bit) & 1) == 1 {
				// we add 1 to the bit here b/c we're ignoring the first feature which is the datastore interface
				// so that we can just reuse the canonical datastore interface, which is quite large
				fields = append(fields, Id(pbFeatureToTypeName[pb.FeaturesResponse_Feature(bit+1)]))
				assignments[Id(pbFeatureToTypeName[pb.FeaturesResponse_Feature(bit+1)])] = Id("dstore")
			}
		}
		structName := fmt.Sprintf("ds%d", i)
		f.Type().Id(structName).Struct(fields...)
		ctorEntries[Lit(i)] = Func().
			Params(Id("dstore").Op("*").Id("Datastore")).
			Qual(dsPkg, "Datastore").
			Block(Return().Op("&").Id(structName).Values(assignments))
	}

	f.Var().Id("ctors").Op("=").Map(Int()).Func().Params(Op("*").Id("Datastore")).Qual(dsPkg, "Datastore").Values(ctorEntries)

	makeNewFunc(f)

	for _, feat := range features {
		feat.emitTypeDef(f)
	}

	fi, err := os.Create("features.go")
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		panic(err)
	}

	err = f.Render(fi)
	if err != nil {
		fi.Close()
		fmt.Printf("error rendering: %v\n", err)
		panic(err)
	}

	err = fi.Close()
	if err != nil {
		fmt.Printf("error closing file: %v\n", err)
		panic(err)

	}
}

func makeNewFunc(f *File) {
	newBlock := []Code{
		Id("featuresSet").Op(":=").Map(Qual(pbPkg, "FeaturesResponse_Feature")).Bool().Values(),
		For(List(Id("_"), Id("f")).Op(":=").Range().Id("features").Block(
			Id("featuresSet").Index(Id("f")).Op("=").True(),
		)),
		Id("i").Op(":=").Lit(0),
	}

	var ifs []Code
	// note that the order of features here must be deterministic
	// we just use the order that is defined by the protobuf
	for i := 1; i < len(pb.FeaturesResponse_Feature_name); i++ {
		ifs = append(ifs,
			If(
				List(Id("_"), Id("ok")).Op(":=").Id("featuresSet").Index(Qual(pbPkg, "FeaturesResponse_Feature").Call(Lit(int(i)))),
				Id("ok"),
			).Block(
				Id("i").Op("+=").Parens(Lit(1).Op("<<").Lit(i-1)),
			),
		)
	}

	newBlock = append(newBlock, ifs...)
	newBlock = append(newBlock, Return(Id("ctors").Index(Id("i")).Call(Id("dstore"))))

	f.Func().Id("newWithFeatures").Params(
		Id("features").Index().Qual(pbPkg, "FeaturesResponse_Feature"),
		Id("dstore").Op("*").Id("Datastore"),
	).Qual(dsPkg, "Datastore").Block(newBlock...)
}
