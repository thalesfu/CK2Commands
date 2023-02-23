package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海堡CastellamareBarony struct {
	feud.BaseBarony
}

var BCastellamare海堡 feud.Barony = &海堡CastellamareBarony{}

func init() {
	f := BCastellamare海堡.(*海堡CastellamareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castellamare",
		TitleName: "海堡",
		TitleCode: "b_castellamare",
	}
}
