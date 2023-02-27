package zaozerye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎列奇纳亚ZarechnayaBarony struct {
	feud.BaseBarony
}

var BZarechnaya扎列奇纳亚 feud.Barony = &扎列奇纳亚ZarechnayaBarony{}

func init() {
    f := BZarechnaya扎列奇纳亚.(*扎列奇纳亚ZarechnayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zarechnaya",
		TitleName: "扎列奇纳亚",
		TitleCode: "b_zarechnaya",
	}
}
