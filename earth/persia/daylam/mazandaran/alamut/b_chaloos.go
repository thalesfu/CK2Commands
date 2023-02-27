package alamut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰卢斯ChaloosBarony struct {
	feud.BaseBarony
}

var BChaloos恰卢斯 feud.Barony = &恰卢斯ChaloosBarony{}

func init() {
    f := BChaloos恰卢斯.(*恰卢斯ChaloosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaloos",
		TitleName: "恰卢斯",
		TitleCode: "b_chaloos",
	}
}
