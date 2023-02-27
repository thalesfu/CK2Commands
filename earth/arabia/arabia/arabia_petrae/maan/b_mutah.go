package maan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆阿泰MutahBarony struct {
	feud.BaseBarony
}

var BMutah穆阿泰 feud.Barony = &穆阿泰MutahBarony{}

func init() {
    f := BMutah穆阿泰.(*穆阿泰MutahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mutah",
		TitleName: "穆阿泰",
		TitleCode: "b_mutah",
	}
}
