package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀多瓦DadrewaBarony struct {
	feud.BaseBarony
}

var BDadrewa陀多瓦 feud.Barony = &陀多瓦DadrewaBarony{}

func init() {
	f := BDadrewa陀多瓦.(*陀多瓦DadrewaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dadrewa",
		TitleName: "陀多瓦",
		TitleCode: "b_dadrewa",
	}
}
