package nordgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳特堡LauterburgBarony struct {
	feud.BaseBarony
}

var BLauterburg劳特堡 feud.Barony = &劳特堡LauterburgBarony{}

func init() {
	f := BLauterburg劳特堡.(*劳特堡LauterburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lauterburg",
		TitleName: "劳特堡",
		TitleCode: "b_lauterburg",
	}
}
