package bjarmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库兹涅佐瓦KusnezowaBarony struct {
	feud.BaseBarony
}

var BKusnezowa库兹涅佐瓦 feud.Barony = &库兹涅佐瓦KusnezowaBarony{}

func init() {
    f := BKusnezowa库兹涅佐瓦.(*库兹涅佐瓦KusnezowaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kusnezowa",
		TitleName: "库兹涅佐瓦",
		TitleCode: "b_kusnezowa",
	}
}
