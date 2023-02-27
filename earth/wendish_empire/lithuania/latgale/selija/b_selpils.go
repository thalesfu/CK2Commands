package selija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尔皮尔斯SelpilsBarony struct {
	feud.BaseBarony
}

var BSelpils塞尔皮尔斯 feud.Barony = &塞尔皮尔斯SelpilsBarony{}

func init() {
    f := BSelpils塞尔皮尔斯.(*塞尔皮尔斯SelpilsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "selpils",
		TitleName: "塞尔皮尔斯",
		TitleCode: "b_selpils",
	}
}
