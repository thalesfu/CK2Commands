package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维尔纽斯VilniusBarony struct {
	feud.BaseBarony
}

var BVilnius维尔纽斯 feud.Barony = &维尔纽斯VilniusBarony{}

func init() {
    f := BVilnius维尔纽斯.(*维尔纽斯VilniusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vilnius",
		TitleName: "维尔纽斯",
		TitleCode: "b_vilnius",
	}
}
