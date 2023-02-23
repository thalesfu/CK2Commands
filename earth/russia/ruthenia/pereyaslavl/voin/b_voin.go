package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃因VoinBarony struct {
	feud.BaseBarony
}

var BVoin沃因 feud.Barony = &沃因VoinBarony{}

func init() {
	f := BVoin沃因.(*沃因VoinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voin",
		TitleName: "沃因",
		TitleCode: "b_voin",
	}
}
