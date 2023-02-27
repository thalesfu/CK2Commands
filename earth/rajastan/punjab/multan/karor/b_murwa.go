package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 牟瓦MurwaBarony struct {
	feud.BaseBarony
}

var BMurwa牟瓦 feud.Barony = &牟瓦MurwaBarony{}

func init() {
    f := BMurwa牟瓦.(*牟瓦MurwaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murwa",
		TitleName: "牟瓦",
		TitleCode: "b_murwa",
	}
}
