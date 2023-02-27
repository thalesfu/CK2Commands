package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡瓦巴巴KwababaBarony struct {
	feud.BaseBarony
}

var BKwababa卡瓦巴巴 feud.Barony = &卡瓦巴巴KwababaBarony{}

func init() {
    f := BKwababa卡瓦巴巴.(*卡瓦巴巴KwababaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kwababa",
		TitleName: "卡瓦巴巴",
		TitleCode: "b_kwababa",
	}
}
