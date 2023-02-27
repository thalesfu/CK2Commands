package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡扎科沃KazakovoBarony struct {
	feud.BaseBarony
}

var BKazakovo卡扎科沃 feud.Barony = &卡扎科沃KazakovoBarony{}

func init() {
    f := BKazakovo卡扎科沃.(*卡扎科沃KazakovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kazakovo",
		TitleName: "卡扎科沃",
		TitleCode: "b_kazakovo",
	}
}
