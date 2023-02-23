package kol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索龙SoronBarony struct {
	feud.BaseBarony
}

var BSoron索龙 feud.Barony = &索龙SoronBarony{}

func init() {
	f := BSoron索龙.(*索龙SoronBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soron",
		TitleName: "索龙",
		TitleCode: "b_soron",
	}
}
