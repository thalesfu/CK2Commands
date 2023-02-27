package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马哈奇MakhachBarony struct {
	feud.BaseBarony
}

var BMakhach马哈奇 feud.Barony = &马哈奇MakhachBarony{}

func init() {
    f := BMakhach马哈奇.(*马哈奇MakhachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "makhach",
		TitleName: "马哈奇",
		TitleCode: "b_makhach",
	}
}
