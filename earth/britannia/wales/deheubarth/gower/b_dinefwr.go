package gower

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪内弗尔DinefwrBarony struct {
	feud.BaseBarony
}

var BDinefwr迪内弗尔 feud.Barony = &迪内弗尔DinefwrBarony{}

func init() {
	f := BDinefwr迪内弗尔.(*迪内弗尔DinefwrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dinefwr",
		TitleName: "迪内弗尔",
		TitleCode: "b_dinefwr",
	}
}
