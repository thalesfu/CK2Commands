package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏库尔SukkurBarony struct {
	feud.BaseBarony
}

var BSukkur苏库尔 feud.Barony = &苏库尔SukkurBarony{}

func init() {
    f := BSukkur苏库尔.(*苏库尔SukkurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukkur",
		TitleName: "苏库尔",
		TitleCode: "b_sukkur",
	}
}
