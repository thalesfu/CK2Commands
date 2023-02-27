package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉阿尔加瓦LaalgabaBarony struct {
	feud.BaseBarony
}

var BLaalgaba拉阿尔加瓦 feud.Barony = &拉阿尔加瓦LaalgabaBarony{}

func init() {
    f := BLaalgaba拉阿尔加瓦.(*拉阿尔加瓦LaalgabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "laalgaba",
		TitleName: "拉阿尔加瓦",
		TitleCode: "b_laalgaba",
	}
}
