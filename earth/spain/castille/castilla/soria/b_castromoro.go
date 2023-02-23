package soria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特罗莫罗CastromoroBarony struct {
	feud.BaseBarony
}

var BCastromoro卡斯特罗莫罗 feud.Barony = &卡斯特罗莫罗CastromoroBarony{}

func init() {
	f := BCastromoro卡斯特罗莫罗.(*卡斯特罗莫罗CastromoroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castromoro",
		TitleName: "卡斯特罗莫罗",
		TitleCode: "b_castromoro",
	}
}
