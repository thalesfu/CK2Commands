package zaragoza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉戈萨ZaragozaBarony struct {
	feud.BaseBarony
}

var BZaragoza萨拉戈萨 feud.Barony = &萨拉戈萨ZaragozaBarony{}

func init() {
    f := BZaragoza萨拉戈萨.(*萨拉戈萨ZaragozaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaragoza",
		TitleName: "萨拉戈萨",
		TitleCode: "b_zaragoza",
	}
}
