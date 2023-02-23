package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍克尔ShekerBarony struct {
	feud.BaseBarony
}

var BSheker舍克尔 feud.Barony = &舍克尔ShekerBarony{}

func init() {
	f := BSheker舍克尔.(*舍克尔ShekerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sheker",
		TitleName: "舍克尔",
		TitleCode: "b_sheker",
	}
}
