package syr_darya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈扎雷KazalyBarony struct {
	feud.BaseBarony
}

var BKazaly哈扎雷 feud.Barony = &哈扎雷KazalyBarony{}

func init() {
    f := BKazaly哈扎雷.(*哈扎雷KazalyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazaly",
		TitleName: "哈扎雷",
		TitleCode: "b_kazaly",
	}
}
