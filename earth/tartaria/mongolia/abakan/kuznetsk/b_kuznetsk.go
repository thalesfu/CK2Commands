package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库兹涅茨克KuznetskBarony struct {
	feud.BaseBarony
}

var BKuznetsk库兹涅茨克 feud.Barony = &库兹涅茨克KuznetskBarony{}

func init() {
	f := BKuznetsk库兹涅茨克.(*库兹涅茨克KuznetskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuznetsk",
		TitleName: "库兹涅茨克",
		TitleCode: "b_kuznetsk",
	}
}
