package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪平斯维克KopingsvikBarony struct {
	feud.BaseBarony
}

var BKopingsvik雪平斯维克 feud.Barony = &雪平斯维克KopingsvikBarony{}

func init() {
	f := BKopingsvik雪平斯维克.(*雪平斯维克KopingsvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kopingsvik",
		TitleName: "雪平斯维克",
		TitleCode: "b_kopingsvik",
	}
}
