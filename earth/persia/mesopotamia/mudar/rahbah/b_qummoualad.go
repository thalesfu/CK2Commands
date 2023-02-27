package rahbah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀穆喀拉德QummoualadBarony struct {
	feud.BaseBarony
}

var BQummoualad喀穆喀拉德 feud.Barony = &喀穆喀拉德QummoualadBarony{}

func init() {
    f := BQummoualad喀穆喀拉德.(*喀穆喀拉德QummoualadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qummoualad",
		TitleName: "喀穆喀拉德",
		TitleCode: "b_qummoualad",
	}
}
