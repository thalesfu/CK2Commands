package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 首波罗SurparakaBarony struct {
	feud.BaseBarony
}

var BSurparaka首波罗 feud.Barony = &首波罗SurparakaBarony{}

func init() {
	f := BSurparaka首波罗.(*首波罗SurparakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "surparaka",
		TitleName: "首波罗",
		TitleCode: "b_surparaka",
	}
}
