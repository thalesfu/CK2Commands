package vastmanland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雪平斯胡斯KopingshusBarony struct {
	feud.BaseBarony
}

var BKopingshus雪平斯胡斯 feud.Barony = &雪平斯胡斯KopingshusBarony{}

func init() {
    f := BKopingshus雪平斯胡斯.(*雪平斯胡斯KopingshusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kopingshus",
		TitleName: "雪平斯胡斯",
		TitleCode: "b_kopingshus",
	}
}
