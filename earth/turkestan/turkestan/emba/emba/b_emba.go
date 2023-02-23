package emba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩巴EmbaBarony struct {
	feud.BaseBarony
}

var BEmba恩巴 feud.Barony = &恩巴EmbaBarony{}

func init() {
	f := BEmba恩巴.(*恩巴EmbaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "emba",
		TitleName: "恩巴",
		TitleCode: "b_emba",
	}
}
