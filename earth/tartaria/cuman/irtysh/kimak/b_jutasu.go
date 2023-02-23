package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 珠塔苏JutasuBarony struct {
	feud.BaseBarony
}

var BJutasu珠塔苏 feud.Barony = &珠塔苏JutasuBarony{}

func init() {
	f := BJutasu珠塔苏.(*珠塔苏JutasuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jutasu",
		TitleName: "珠塔苏",
		TitleCode: "b_jutasu",
	}
}
