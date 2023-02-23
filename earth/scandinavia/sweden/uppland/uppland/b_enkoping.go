package uppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩雪平EnkopingBarony struct {
	feud.BaseBarony
}

var BEnkoping恩雪平 feud.Barony = &恩雪平EnkopingBarony{}

func init() {
	f := BEnkoping恩雪平.(*恩雪平EnkopingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enkoping",
		TitleName: "恩雪平",
		TitleCode: "b_enkoping",
	}
}
