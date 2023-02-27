package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利亚克拉KaliakraBarony struct {
	feud.BaseBarony
}

var BKaliakra卡利亚克拉 feud.Barony = &卡利亚克拉KaliakraBarony{}

func init() {
    f := BKaliakra卡利亚克拉.(*卡利亚克拉KaliakraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaliakra",
		TitleName: "卡利亚克拉",
		TitleCode: "b_kaliakra",
	}
}
