package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔布纳DelbnaBarony struct {
	feud.BaseBarony
}

var BDelbna德尔布纳 feud.Barony = &德尔布纳DelbnaBarony{}

func init() {
    f := BDelbna德尔布纳.(*德尔布纳DelbnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "delbna",
		TitleName: "德尔布纳",
		TitleCode: "b_delbna",
	}
}
