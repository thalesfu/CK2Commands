package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎伊雷克ZhailykBarony struct {
	feud.BaseBarony
}

var BZhailyk扎伊雷克 feud.Barony = &扎伊雷克ZhailykBarony{}

func init() {
	f := BZhailyk扎伊雷克.(*扎伊雷克ZhailykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhailyk",
		TitleName: "扎伊雷克",
		TitleCode: "b_zhailyk",
	}
}
