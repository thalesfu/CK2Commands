package pecs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣勒林茨SzentlorincBarony struct {
	feud.BaseBarony
}

var BSzentlorinc圣勒林茨 feud.Barony = &圣勒林茨SzentlorincBarony{}

func init() {
    f := BSzentlorinc圣勒林茨.(*圣勒林茨SzentlorincBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szentlorinc",
		TitleName: "圣勒林茨",
		TitleCode: "b_szentlorinc",
	}
}
