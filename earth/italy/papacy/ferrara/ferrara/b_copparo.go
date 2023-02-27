package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科帕罗CopparoBarony struct {
	feud.BaseBarony
}

var BCopparo科帕罗 feud.Barony = &科帕罗CopparoBarony{}

func init() {
    f := BCopparo科帕罗.(*科帕罗CopparoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "copparo",
		TitleName: "科帕罗",
		TitleCode: "b_copparo",
	}
}
