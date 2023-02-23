package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舒巴尔ShubarBarony struct {
	feud.BaseBarony
}

var BShubar舒巴尔 feud.Barony = &舒巴尔ShubarBarony{}

func init() {
	f := BShubar舒巴尔.(*舒巴尔ShubarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shubar",
		TitleName: "舒巴尔",
		TitleCode: "b_shubar",
	}
}
