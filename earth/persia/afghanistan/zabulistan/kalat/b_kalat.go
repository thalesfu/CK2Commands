package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉特KalatBarony struct {
	feud.BaseBarony
}

var BKalat卡拉特 feud.Barony = &卡拉特KalatBarony{}

func init() {
	f := BKalat卡拉特.(*卡拉特KalatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalat",
		TitleName: "卡拉特",
		TitleCode: "b_kalat",
	}
}
