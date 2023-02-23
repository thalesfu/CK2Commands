package fes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 非斯FesBarony struct {
	feud.BaseBarony
}

var BFes非斯 feud.Barony = &非斯FesBarony{}

func init() {
	f := BFes非斯.(*非斯FesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fes",
		TitleName: "非斯",
		TitleCode: "b_fes",
	}
}
