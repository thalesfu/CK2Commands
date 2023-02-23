package nevers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旺德内斯VandenesseBarony struct {
	feud.BaseBarony
}

var BVandenesse旺德内斯 feud.Barony = &旺德内斯VandenesseBarony{}

func init() {
	f := BVandenesse旺德内斯.(*旺德内斯VandenesseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vandenesse",
		TitleName: "旺德内斯",
		TitleCode: "b_vandenesse",
	}
}
