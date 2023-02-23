package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏沃罗瓦SuvorovaBarony struct {
	feud.BaseBarony
}

var BSuvorova苏沃罗瓦 feud.Barony = &苏沃罗瓦SuvorovaBarony{}

func init() {
	f := BSuvorova苏沃罗瓦.(*苏沃罗瓦SuvorovaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suvorova",
		TitleName: "苏沃罗瓦",
		TitleCode: "b_suvorova",
	}
}
