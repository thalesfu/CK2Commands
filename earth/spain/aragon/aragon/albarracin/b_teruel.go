package albarracin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁埃尔TeruelBarony struct {
	feud.BaseBarony
}

var BTeruel特鲁埃尔 feud.Barony = &特鲁埃尔TeruelBarony{}

func init() {
	f := BTeruel特鲁埃尔.(*特鲁埃尔TeruelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "teruel",
		TitleName: "特鲁埃尔",
		TitleCode: "b_teruel",
	}
}
