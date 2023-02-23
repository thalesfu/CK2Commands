package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏哈尔SuharBarony struct {
	feud.BaseBarony
}

var BSuhar苏哈尔 feud.Barony = &苏哈尔SuharBarony{}

func init() {
	f := BSuhar苏哈尔.(*苏哈尔SuharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suhar",
		TitleName: "苏哈尔",
		TitleCode: "b_suhar",
	}
}
