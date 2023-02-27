package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维什科夫VyskovBarony struct {
	feud.BaseBarony
}

var BVyskov维什科夫 feud.Barony = &维什科夫VyskovBarony{}

func init() {
    f := BVyskov维什科夫.(*维什科夫VyskovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vyskov",
		TitleName: "维什科夫",
		TitleCode: "b_vyskov",
	}
}
