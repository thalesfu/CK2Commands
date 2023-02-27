package korinthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕萨瓦PassavaBarony struct {
	feud.BaseBarony
}

var BPassava帕萨瓦 feud.Barony = &帕萨瓦PassavaBarony{}

func init() {
    f := BPassava帕萨瓦.(*帕萨瓦PassavaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "passava",
		TitleName: "帕萨瓦",
		TitleCode: "b_passava",
	}
}
