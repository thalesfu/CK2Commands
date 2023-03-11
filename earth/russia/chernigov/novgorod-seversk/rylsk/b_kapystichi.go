package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡佩斯季奇KapystichiBarony struct {
	feud.BaseBarony
}

var BKapystichi卡佩斯季奇 feud.Barony = &卡佩斯季奇KapystichiBarony{}

func init() {
    f := BKapystichi卡佩斯季奇.(*卡佩斯季奇KapystichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapystichi",
		TitleName: "卡佩斯季奇",
		TitleCode: "b_kapystichi",
	}
}
