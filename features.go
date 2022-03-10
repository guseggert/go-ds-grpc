package grpcds

import (
	"context"
	proto "github.com/guseggert/go-ds-grpc/proto"
	godatastore "github.com/ipfs/go-datastore"
	"time"
)

type ds0 struct {
	godatastore.Datastore
}
type ds1 struct {
	godatastore.Datastore
	batching
}
type ds2 struct {
	godatastore.Datastore
	checked
}
type ds3 struct {
	godatastore.Datastore
	batching
	checked
}
type ds4 struct {
	godatastore.Datastore
	scrubbed
}
type ds5 struct {
	godatastore.Datastore
	batching
	scrubbed
}
type ds6 struct {
	godatastore.Datastore
	checked
	scrubbed
}
type ds7 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
}
type ds8 struct {
	godatastore.Datastore
	gc
}
type ds9 struct {
	godatastore.Datastore
	batching
	gc
}
type ds10 struct {
	godatastore.Datastore
	checked
	gc
}
type ds11 struct {
	godatastore.Datastore
	batching
	checked
	gc
}
type ds12 struct {
	godatastore.Datastore
	scrubbed
	gc
}
type ds13 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
}
type ds14 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
}
type ds15 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
}
type ds16 struct {
	godatastore.Datastore
	persistent
}
type ds17 struct {
	godatastore.Datastore
	batching
	persistent
}
type ds18 struct {
	godatastore.Datastore
	checked
	persistent
}
type ds19 struct {
	godatastore.Datastore
	batching
	checked
	persistent
}
type ds20 struct {
	godatastore.Datastore
	scrubbed
	persistent
}
type ds21 struct {
	godatastore.Datastore
	batching
	scrubbed
	persistent
}
type ds22 struct {
	godatastore.Datastore
	checked
	scrubbed
	persistent
}
type ds23 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	persistent
}
type ds24 struct {
	godatastore.Datastore
	gc
	persistent
}
type ds25 struct {
	godatastore.Datastore
	batching
	gc
	persistent
}
type ds26 struct {
	godatastore.Datastore
	checked
	gc
	persistent
}
type ds27 struct {
	godatastore.Datastore
	batching
	checked
	gc
	persistent
}
type ds28 struct {
	godatastore.Datastore
	scrubbed
	gc
	persistent
}
type ds29 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	persistent
}
type ds30 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	persistent
}
type ds31 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	persistent
}
type ds32 struct {
	godatastore.Datastore
	ttl
}
type ds33 struct {
	godatastore.Datastore
	batching
	ttl
}
type ds34 struct {
	godatastore.Datastore
	checked
	ttl
}
type ds35 struct {
	godatastore.Datastore
	batching
	checked
	ttl
}
type ds36 struct {
	godatastore.Datastore
	scrubbed
	ttl
}
type ds37 struct {
	godatastore.Datastore
	batching
	scrubbed
	ttl
}
type ds38 struct {
	godatastore.Datastore
	checked
	scrubbed
	ttl
}
type ds39 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	ttl
}
type ds40 struct {
	godatastore.Datastore
	gc
	ttl
}
type ds41 struct {
	godatastore.Datastore
	batching
	gc
	ttl
}
type ds42 struct {
	godatastore.Datastore
	checked
	gc
	ttl
}
type ds43 struct {
	godatastore.Datastore
	batching
	checked
	gc
	ttl
}
type ds44 struct {
	godatastore.Datastore
	scrubbed
	gc
	ttl
}
type ds45 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	ttl
}
type ds46 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	ttl
}
type ds47 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	ttl
}
type ds48 struct {
	godatastore.Datastore
	persistent
	ttl
}
type ds49 struct {
	godatastore.Datastore
	batching
	persistent
	ttl
}
type ds50 struct {
	godatastore.Datastore
	checked
	persistent
	ttl
}
type ds51 struct {
	godatastore.Datastore
	batching
	checked
	persistent
	ttl
}
type ds52 struct {
	godatastore.Datastore
	scrubbed
	persistent
	ttl
}
type ds53 struct {
	godatastore.Datastore
	batching
	scrubbed
	persistent
	ttl
}
type ds54 struct {
	godatastore.Datastore
	checked
	scrubbed
	persistent
	ttl
}
type ds55 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	persistent
	ttl
}
type ds56 struct {
	godatastore.Datastore
	gc
	persistent
	ttl
}
type ds57 struct {
	godatastore.Datastore
	batching
	gc
	persistent
	ttl
}
type ds58 struct {
	godatastore.Datastore
	checked
	gc
	persistent
	ttl
}
type ds59 struct {
	godatastore.Datastore
	batching
	checked
	gc
	persistent
	ttl
}
type ds60 struct {
	godatastore.Datastore
	scrubbed
	gc
	persistent
	ttl
}
type ds61 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	persistent
	ttl
}
type ds62 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	persistent
	ttl
}
type ds63 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	persistent
	ttl
}
type ds64 struct {
	godatastore.Datastore
	transaction
}
type ds65 struct {
	godatastore.Datastore
	batching
	transaction
}
type ds66 struct {
	godatastore.Datastore
	checked
	transaction
}
type ds67 struct {
	godatastore.Datastore
	batching
	checked
	transaction
}
type ds68 struct {
	godatastore.Datastore
	scrubbed
	transaction
}
type ds69 struct {
	godatastore.Datastore
	batching
	scrubbed
	transaction
}
type ds70 struct {
	godatastore.Datastore
	checked
	scrubbed
	transaction
}
type ds71 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	transaction
}
type ds72 struct {
	godatastore.Datastore
	gc
	transaction
}
type ds73 struct {
	godatastore.Datastore
	batching
	gc
	transaction
}
type ds74 struct {
	godatastore.Datastore
	checked
	gc
	transaction
}
type ds75 struct {
	godatastore.Datastore
	batching
	checked
	gc
	transaction
}
type ds76 struct {
	godatastore.Datastore
	scrubbed
	gc
	transaction
}
type ds77 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	transaction
}
type ds78 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	transaction
}
type ds79 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	transaction
}
type ds80 struct {
	godatastore.Datastore
	persistent
	transaction
}
type ds81 struct {
	godatastore.Datastore
	batching
	persistent
	transaction
}
type ds82 struct {
	godatastore.Datastore
	checked
	persistent
	transaction
}
type ds83 struct {
	godatastore.Datastore
	batching
	checked
	persistent
	transaction
}
type ds84 struct {
	godatastore.Datastore
	scrubbed
	persistent
	transaction
}
type ds85 struct {
	godatastore.Datastore
	batching
	scrubbed
	persistent
	transaction
}
type ds86 struct {
	godatastore.Datastore
	checked
	scrubbed
	persistent
	transaction
}
type ds87 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	persistent
	transaction
}
type ds88 struct {
	godatastore.Datastore
	gc
	persistent
	transaction
}
type ds89 struct {
	godatastore.Datastore
	batching
	gc
	persistent
	transaction
}
type ds90 struct {
	godatastore.Datastore
	checked
	gc
	persistent
	transaction
}
type ds91 struct {
	godatastore.Datastore
	batching
	checked
	gc
	persistent
	transaction
}
type ds92 struct {
	godatastore.Datastore
	scrubbed
	gc
	persistent
	transaction
}
type ds93 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	persistent
	transaction
}
type ds94 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	persistent
	transaction
}
type ds95 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	persistent
	transaction
}
type ds96 struct {
	godatastore.Datastore
	ttl
	transaction
}
type ds97 struct {
	godatastore.Datastore
	batching
	ttl
	transaction
}
type ds98 struct {
	godatastore.Datastore
	checked
	ttl
	transaction
}
type ds99 struct {
	godatastore.Datastore
	batching
	checked
	ttl
	transaction
}
type ds100 struct {
	godatastore.Datastore
	scrubbed
	ttl
	transaction
}
type ds101 struct {
	godatastore.Datastore
	batching
	scrubbed
	ttl
	transaction
}
type ds102 struct {
	godatastore.Datastore
	checked
	scrubbed
	ttl
	transaction
}
type ds103 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	ttl
	transaction
}
type ds104 struct {
	godatastore.Datastore
	gc
	ttl
	transaction
}
type ds105 struct {
	godatastore.Datastore
	batching
	gc
	ttl
	transaction
}
type ds106 struct {
	godatastore.Datastore
	checked
	gc
	ttl
	transaction
}
type ds107 struct {
	godatastore.Datastore
	batching
	checked
	gc
	ttl
	transaction
}
type ds108 struct {
	godatastore.Datastore
	scrubbed
	gc
	ttl
	transaction
}
type ds109 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	ttl
	transaction
}
type ds110 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	ttl
	transaction
}
type ds111 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	ttl
	transaction
}
type ds112 struct {
	godatastore.Datastore
	persistent
	ttl
	transaction
}
type ds113 struct {
	godatastore.Datastore
	batching
	persistent
	ttl
	transaction
}
type ds114 struct {
	godatastore.Datastore
	checked
	persistent
	ttl
	transaction
}
type ds115 struct {
	godatastore.Datastore
	batching
	checked
	persistent
	ttl
	transaction
}
type ds116 struct {
	godatastore.Datastore
	scrubbed
	persistent
	ttl
	transaction
}
type ds117 struct {
	godatastore.Datastore
	batching
	scrubbed
	persistent
	ttl
	transaction
}
type ds118 struct {
	godatastore.Datastore
	checked
	scrubbed
	persistent
	ttl
	transaction
}
type ds119 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	persistent
	ttl
	transaction
}
type ds120 struct {
	godatastore.Datastore
	gc
	persistent
	ttl
	transaction
}
type ds121 struct {
	godatastore.Datastore
	batching
	gc
	persistent
	ttl
	transaction
}
type ds122 struct {
	godatastore.Datastore
	checked
	gc
	persistent
	ttl
	transaction
}
type ds123 struct {
	godatastore.Datastore
	batching
	checked
	gc
	persistent
	ttl
	transaction
}
type ds124 struct {
	godatastore.Datastore
	scrubbed
	gc
	persistent
	ttl
	transaction
}
type ds125 struct {
	godatastore.Datastore
	batching
	scrubbed
	gc
	persistent
	ttl
	transaction
}
type ds126 struct {
	godatastore.Datastore
	checked
	scrubbed
	gc
	persistent
	ttl
	transaction
}
type ds127 struct {
	godatastore.Datastore
	batching
	checked
	scrubbed
	gc
	persistent
	ttl
	transaction
}

var ctors = map[int]func(*Datastore) godatastore.Datastore{
	0: func(dstore *Datastore) godatastore.Datastore {
		return &ds0{Datastore: dstore}
	},
	1: func(dstore *Datastore) godatastore.Datastore {
		return &ds1{
			Datastore: dstore,
			batching:  dstore,
		}
	},
	10: func(dstore *Datastore) godatastore.Datastore {
		return &ds10{
			Datastore: dstore,
			checked:   dstore,
			gc:        dstore,
		}
	},
	100: func(dstore *Datastore) godatastore.Datastore {
		return &ds100{
			Datastore:   dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	101: func(dstore *Datastore) godatastore.Datastore {
		return &ds101{
			Datastore:   dstore,
			batching:    dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	102: func(dstore *Datastore) godatastore.Datastore {
		return &ds102{
			Datastore:   dstore,
			checked:     dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	103: func(dstore *Datastore) godatastore.Datastore {
		return &ds103{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	104: func(dstore *Datastore) godatastore.Datastore {
		return &ds104{
			Datastore:   dstore,
			gc:          dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	105: func(dstore *Datastore) godatastore.Datastore {
		return &ds105{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	106: func(dstore *Datastore) godatastore.Datastore {
		return &ds106{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	107: func(dstore *Datastore) godatastore.Datastore {
		return &ds107{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	108: func(dstore *Datastore) godatastore.Datastore {
		return &ds108{
			Datastore:   dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	109: func(dstore *Datastore) godatastore.Datastore {
		return &ds109{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	11: func(dstore *Datastore) godatastore.Datastore {
		return &ds11{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			gc:        dstore,
		}
	},
	110: func(dstore *Datastore) godatastore.Datastore {
		return &ds110{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	111: func(dstore *Datastore) godatastore.Datastore {
		return &ds111{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	112: func(dstore *Datastore) godatastore.Datastore {
		return &ds112{
			Datastore:   dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	113: func(dstore *Datastore) godatastore.Datastore {
		return &ds113{
			Datastore:   dstore,
			batching:    dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	114: func(dstore *Datastore) godatastore.Datastore {
		return &ds114{
			Datastore:   dstore,
			checked:     dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	115: func(dstore *Datastore) godatastore.Datastore {
		return &ds115{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	116: func(dstore *Datastore) godatastore.Datastore {
		return &ds116{
			Datastore:   dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	117: func(dstore *Datastore) godatastore.Datastore {
		return &ds117{
			Datastore:   dstore,
			batching:    dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	118: func(dstore *Datastore) godatastore.Datastore {
		return &ds118{
			Datastore:   dstore,
			checked:     dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	119: func(dstore *Datastore) godatastore.Datastore {
		return &ds119{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	12: func(dstore *Datastore) godatastore.Datastore {
		return &ds12{
			Datastore: dstore,
			gc:        dstore,
			scrubbed:  dstore,
		}
	},
	120: func(dstore *Datastore) godatastore.Datastore {
		return &ds120{
			Datastore:   dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	121: func(dstore *Datastore) godatastore.Datastore {
		return &ds121{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	122: func(dstore *Datastore) godatastore.Datastore {
		return &ds122{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	123: func(dstore *Datastore) godatastore.Datastore {
		return &ds123{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	124: func(dstore *Datastore) godatastore.Datastore {
		return &ds124{
			Datastore:   dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	125: func(dstore *Datastore) godatastore.Datastore {
		return &ds125{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	126: func(dstore *Datastore) godatastore.Datastore {
		return &ds126{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	127: func(dstore *Datastore) godatastore.Datastore {
		return &ds127{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	13: func(dstore *Datastore) godatastore.Datastore {
		return &ds13{
			Datastore: dstore,
			batching:  dstore,
			gc:        dstore,
			scrubbed:  dstore,
		}
	},
	14: func(dstore *Datastore) godatastore.Datastore {
		return &ds14{
			Datastore: dstore,
			checked:   dstore,
			gc:        dstore,
			scrubbed:  dstore,
		}
	},
	15: func(dstore *Datastore) godatastore.Datastore {
		return &ds15{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			gc:        dstore,
			scrubbed:  dstore,
		}
	},
	16: func(dstore *Datastore) godatastore.Datastore {
		return &ds16{
			Datastore:  dstore,
			persistent: dstore,
		}
	},
	17: func(dstore *Datastore) godatastore.Datastore {
		return &ds17{
			Datastore:  dstore,
			batching:   dstore,
			persistent: dstore,
		}
	},
	18: func(dstore *Datastore) godatastore.Datastore {
		return &ds18{
			Datastore:  dstore,
			checked:    dstore,
			persistent: dstore,
		}
	},
	19: func(dstore *Datastore) godatastore.Datastore {
		return &ds19{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			persistent: dstore,
		}
	},
	2: func(dstore *Datastore) godatastore.Datastore {
		return &ds2{
			Datastore: dstore,
			checked:   dstore,
		}
	},
	20: func(dstore *Datastore) godatastore.Datastore {
		return &ds20{
			Datastore:  dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	21: func(dstore *Datastore) godatastore.Datastore {
		return &ds21{
			Datastore:  dstore,
			batching:   dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	22: func(dstore *Datastore) godatastore.Datastore {
		return &ds22{
			Datastore:  dstore,
			checked:    dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	23: func(dstore *Datastore) godatastore.Datastore {
		return &ds23{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	24: func(dstore *Datastore) godatastore.Datastore {
		return &ds24{
			Datastore:  dstore,
			gc:         dstore,
			persistent: dstore,
		}
	},
	25: func(dstore *Datastore) godatastore.Datastore {
		return &ds25{
			Datastore:  dstore,
			batching:   dstore,
			gc:         dstore,
			persistent: dstore,
		}
	},
	26: func(dstore *Datastore) godatastore.Datastore {
		return &ds26{
			Datastore:  dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
		}
	},
	27: func(dstore *Datastore) godatastore.Datastore {
		return &ds27{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
		}
	},
	28: func(dstore *Datastore) godatastore.Datastore {
		return &ds28{
			Datastore:  dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	29: func(dstore *Datastore) godatastore.Datastore {
		return &ds29{
			Datastore:  dstore,
			batching:   dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	3: func(dstore *Datastore) godatastore.Datastore {
		return &ds3{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
		}
	},
	30: func(dstore *Datastore) godatastore.Datastore {
		return &ds30{
			Datastore:  dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	31: func(dstore *Datastore) godatastore.Datastore {
		return &ds31{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
		}
	},
	32: func(dstore *Datastore) godatastore.Datastore {
		return &ds32{
			Datastore: dstore,
			ttl:       dstore,
		}
	},
	33: func(dstore *Datastore) godatastore.Datastore {
		return &ds33{
			Datastore: dstore,
			batching:  dstore,
			ttl:       dstore,
		}
	},
	34: func(dstore *Datastore) godatastore.Datastore {
		return &ds34{
			Datastore: dstore,
			checked:   dstore,
			ttl:       dstore,
		}
	},
	35: func(dstore *Datastore) godatastore.Datastore {
		return &ds35{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			ttl:       dstore,
		}
	},
	36: func(dstore *Datastore) godatastore.Datastore {
		return &ds36{
			Datastore: dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	37: func(dstore *Datastore) godatastore.Datastore {
		return &ds37{
			Datastore: dstore,
			batching:  dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	38: func(dstore *Datastore) godatastore.Datastore {
		return &ds38{
			Datastore: dstore,
			checked:   dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	39: func(dstore *Datastore) godatastore.Datastore {
		return &ds39{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	4: func(dstore *Datastore) godatastore.Datastore {
		return &ds4{
			Datastore: dstore,
			scrubbed:  dstore,
		}
	},
	40: func(dstore *Datastore) godatastore.Datastore {
		return &ds40{
			Datastore: dstore,
			gc:        dstore,
			ttl:       dstore,
		}
	},
	41: func(dstore *Datastore) godatastore.Datastore {
		return &ds41{
			Datastore: dstore,
			batching:  dstore,
			gc:        dstore,
			ttl:       dstore,
		}
	},
	42: func(dstore *Datastore) godatastore.Datastore {
		return &ds42{
			Datastore: dstore,
			checked:   dstore,
			gc:        dstore,
			ttl:       dstore,
		}
	},
	43: func(dstore *Datastore) godatastore.Datastore {
		return &ds43{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			gc:        dstore,
			ttl:       dstore,
		}
	},
	44: func(dstore *Datastore) godatastore.Datastore {
		return &ds44{
			Datastore: dstore,
			gc:        dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	45: func(dstore *Datastore) godatastore.Datastore {
		return &ds45{
			Datastore: dstore,
			batching:  dstore,
			gc:        dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	46: func(dstore *Datastore) godatastore.Datastore {
		return &ds46{
			Datastore: dstore,
			checked:   dstore,
			gc:        dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	47: func(dstore *Datastore) godatastore.Datastore {
		return &ds47{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			gc:        dstore,
			scrubbed:  dstore,
			ttl:       dstore,
		}
	},
	48: func(dstore *Datastore) godatastore.Datastore {
		return &ds48{
			Datastore:  dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	49: func(dstore *Datastore) godatastore.Datastore {
		return &ds49{
			Datastore:  dstore,
			batching:   dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	5: func(dstore *Datastore) godatastore.Datastore {
		return &ds5{
			Datastore: dstore,
			batching:  dstore,
			scrubbed:  dstore,
		}
	},
	50: func(dstore *Datastore) godatastore.Datastore {
		return &ds50{
			Datastore:  dstore,
			checked:    dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	51: func(dstore *Datastore) godatastore.Datastore {
		return &ds51{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	52: func(dstore *Datastore) godatastore.Datastore {
		return &ds52{
			Datastore:  dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	53: func(dstore *Datastore) godatastore.Datastore {
		return &ds53{
			Datastore:  dstore,
			batching:   dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	54: func(dstore *Datastore) godatastore.Datastore {
		return &ds54{
			Datastore:  dstore,
			checked:    dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	55: func(dstore *Datastore) godatastore.Datastore {
		return &ds55{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	56: func(dstore *Datastore) godatastore.Datastore {
		return &ds56{
			Datastore:  dstore,
			gc:         dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	57: func(dstore *Datastore) godatastore.Datastore {
		return &ds57{
			Datastore:  dstore,
			batching:   dstore,
			gc:         dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	58: func(dstore *Datastore) godatastore.Datastore {
		return &ds58{
			Datastore:  dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	59: func(dstore *Datastore) godatastore.Datastore {
		return &ds59{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			ttl:        dstore,
		}
	},
	6: func(dstore *Datastore) godatastore.Datastore {
		return &ds6{
			Datastore: dstore,
			checked:   dstore,
			scrubbed:  dstore,
		}
	},
	60: func(dstore *Datastore) godatastore.Datastore {
		return &ds60{
			Datastore:  dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	61: func(dstore *Datastore) godatastore.Datastore {
		return &ds61{
			Datastore:  dstore,
			batching:   dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	62: func(dstore *Datastore) godatastore.Datastore {
		return &ds62{
			Datastore:  dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	63: func(dstore *Datastore) godatastore.Datastore {
		return &ds63{
			Datastore:  dstore,
			batching:   dstore,
			checked:    dstore,
			gc:         dstore,
			persistent: dstore,
			scrubbed:   dstore,
			ttl:        dstore,
		}
	},
	64: func(dstore *Datastore) godatastore.Datastore {
		return &ds64{
			Datastore:   dstore,
			transaction: dstore,
		}
	},
	65: func(dstore *Datastore) godatastore.Datastore {
		return &ds65{
			Datastore:   dstore,
			batching:    dstore,
			transaction: dstore,
		}
	},
	66: func(dstore *Datastore) godatastore.Datastore {
		return &ds66{
			Datastore:   dstore,
			checked:     dstore,
			transaction: dstore,
		}
	},
	67: func(dstore *Datastore) godatastore.Datastore {
		return &ds67{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			transaction: dstore,
		}
	},
	68: func(dstore *Datastore) godatastore.Datastore {
		return &ds68{
			Datastore:   dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	69: func(dstore *Datastore) godatastore.Datastore {
		return &ds69{
			Datastore:   dstore,
			batching:    dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	7: func(dstore *Datastore) godatastore.Datastore {
		return &ds7{
			Datastore: dstore,
			batching:  dstore,
			checked:   dstore,
			scrubbed:  dstore,
		}
	},
	70: func(dstore *Datastore) godatastore.Datastore {
		return &ds70{
			Datastore:   dstore,
			checked:     dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	71: func(dstore *Datastore) godatastore.Datastore {
		return &ds71{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	72: func(dstore *Datastore) godatastore.Datastore {
		return &ds72{
			Datastore:   dstore,
			gc:          dstore,
			transaction: dstore,
		}
	},
	73: func(dstore *Datastore) godatastore.Datastore {
		return &ds73{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			transaction: dstore,
		}
	},
	74: func(dstore *Datastore) godatastore.Datastore {
		return &ds74{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			transaction: dstore,
		}
	},
	75: func(dstore *Datastore) godatastore.Datastore {
		return &ds75{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			transaction: dstore,
		}
	},
	76: func(dstore *Datastore) godatastore.Datastore {
		return &ds76{
			Datastore:   dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	77: func(dstore *Datastore) godatastore.Datastore {
		return &ds77{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	78: func(dstore *Datastore) godatastore.Datastore {
		return &ds78{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	79: func(dstore *Datastore) godatastore.Datastore {
		return &ds79{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	8: func(dstore *Datastore) godatastore.Datastore {
		return &ds8{
			Datastore: dstore,
			gc:        dstore,
		}
	},
	80: func(dstore *Datastore) godatastore.Datastore {
		return &ds80{
			Datastore:   dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	81: func(dstore *Datastore) godatastore.Datastore {
		return &ds81{
			Datastore:   dstore,
			batching:    dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	82: func(dstore *Datastore) godatastore.Datastore {
		return &ds82{
			Datastore:   dstore,
			checked:     dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	83: func(dstore *Datastore) godatastore.Datastore {
		return &ds83{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	84: func(dstore *Datastore) godatastore.Datastore {
		return &ds84{
			Datastore:   dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	85: func(dstore *Datastore) godatastore.Datastore {
		return &ds85{
			Datastore:   dstore,
			batching:    dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	86: func(dstore *Datastore) godatastore.Datastore {
		return &ds86{
			Datastore:   dstore,
			checked:     dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	87: func(dstore *Datastore) godatastore.Datastore {
		return &ds87{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	88: func(dstore *Datastore) godatastore.Datastore {
		return &ds88{
			Datastore:   dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	89: func(dstore *Datastore) godatastore.Datastore {
		return &ds89{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	9: func(dstore *Datastore) godatastore.Datastore {
		return &ds9{
			Datastore: dstore,
			batching:  dstore,
			gc:        dstore,
		}
	},
	90: func(dstore *Datastore) godatastore.Datastore {
		return &ds90{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	91: func(dstore *Datastore) godatastore.Datastore {
		return &ds91{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			transaction: dstore,
		}
	},
	92: func(dstore *Datastore) godatastore.Datastore {
		return &ds92{
			Datastore:   dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	93: func(dstore *Datastore) godatastore.Datastore {
		return &ds93{
			Datastore:   dstore,
			batching:    dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	94: func(dstore *Datastore) godatastore.Datastore {
		return &ds94{
			Datastore:   dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	95: func(dstore *Datastore) godatastore.Datastore {
		return &ds95{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			gc:          dstore,
			persistent:  dstore,
			scrubbed:    dstore,
			transaction: dstore,
		}
	},
	96: func(dstore *Datastore) godatastore.Datastore {
		return &ds96{
			Datastore:   dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	97: func(dstore *Datastore) godatastore.Datastore {
		return &ds97{
			Datastore:   dstore,
			batching:    dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	98: func(dstore *Datastore) godatastore.Datastore {
		return &ds98{
			Datastore:   dstore,
			checked:     dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
	99: func(dstore *Datastore) godatastore.Datastore {
		return &ds99{
			Datastore:   dstore,
			batching:    dstore,
			checked:     dstore,
			transaction: dstore,
			ttl:         dstore,
		}
	},
}

func newWithFeatures(features []proto.FeaturesResponse_Feature, dstore *Datastore) godatastore.Datastore {
	featuresSet := map[proto.FeaturesResponse_Feature]bool{}
	for _, f := range features {
		featuresSet[f] = true
	}
	i := 0
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(1)]; ok {
		i += (1 << 0)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(2)]; ok {
		i += (1 << 1)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(3)]; ok {
		i += (1 << 2)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(4)]; ok {
		i += (1 << 3)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(5)]; ok {
		i += (1 << 4)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(6)]; ok {
		i += (1 << 5)
	}
	if _, ok := featuresSet[proto.FeaturesResponse_Feature(7)]; ok {
		i += (1 << 6)
	}
	return ctors[i](dstore)
}

type batching interface {
	Batch(context.Context) (godatastore.Batch, error)
}
type checked interface {
	Check(context.Context) error
}
type scrubbed interface {
	Scrub(context.Context) error
}
type gc interface {
	CollectGarbage(context.Context) error
}
type persistent interface {
	DiskUsage(context.Context) (uint64, error)
}
type ttl interface {
	PutWithTTL(context.Context, godatastore.Key, []byte, time.Duration) error
	SetTTL(context.Context, godatastore.Key, time.Duration) error
	GetExpiration(context.Context, godatastore.Key) (time.Time, error)
}
type transaction interface {
	NewTransaction(context.Context, bool) (godatastore.Txn, error)
}
