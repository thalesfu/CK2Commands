package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 劳瑟LowtherBarony struct {
	feud.BaseBarony
}

var BLowther劳瑟 feud.Barony = &劳瑟LowtherBarony{}

func init() {
	f := BLowther劳瑟.(*劳瑟LowtherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lowther",
		TitleName: "劳瑟",
		TitleCode: "b_lowther",
	}
}
