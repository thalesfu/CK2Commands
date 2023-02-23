package kholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚热尔比齐YazhelbitsyBarony struct {
	feud.BaseBarony
}

var BYazhelbitsy亚热尔比齐 feud.Barony = &亚热尔比齐YazhelbitsyBarony{}

func init() {
	f := BYazhelbitsy亚热尔比齐.(*亚热尔比齐YazhelbitsyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yazhelbitsy",
		TitleName: "亚热尔比齐",
		TitleCode: "b_yazhelbitsy",
	}
}
