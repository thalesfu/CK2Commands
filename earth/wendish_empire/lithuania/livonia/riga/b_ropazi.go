package riga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗帕日RopaziBarony struct {
	feud.BaseBarony
}

var BRopazi罗帕日 feud.Barony = &罗帕日RopaziBarony{}

func init() {
    f := BRopazi罗帕日.(*罗帕日RopaziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ropazi",
		TitleName: "罗帕日",
		TitleCode: "b_ropazi",
	}
}
