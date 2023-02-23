package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 岳普湖YopurgaBarony struct {
	feud.BaseBarony
}

var BYopurga岳普湖 feud.Barony = &岳普湖YopurgaBarony{}

func init() {
	f := BYopurga岳普湖.(*岳普湖YopurgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yopurga",
		TitleName: "岳普湖",
		TitleCode: "b_yopurga",
	}
}
