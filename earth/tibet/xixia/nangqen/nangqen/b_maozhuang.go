package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛庄MaozhuangBarony struct {
	feud.BaseBarony
}

var BMaozhuang毛庄 feud.Barony = &毛庄MaozhuangBarony{}

func init() {
    f := BMaozhuang毛庄.(*毛庄MaozhuangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maozhuang",
		TitleName: "毛庄",
		TitleCode: "b_maozhuang",
	}
}
