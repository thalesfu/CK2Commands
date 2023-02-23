package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩尼斯EnnisBarony struct {
	feud.BaseBarony
}

var BEnnis恩尼斯 feud.Barony = &恩尼斯EnnisBarony{}

func init() {
	f := BEnnis恩尼斯.(*恩尼斯EnnisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ennis",
		TitleName: "恩尼斯",
		TitleCode: "b_ennis",
	}
}
