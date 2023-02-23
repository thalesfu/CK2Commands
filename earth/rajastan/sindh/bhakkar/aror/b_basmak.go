package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴斯马克BasmakBarony struct {
	feud.BaseBarony
}

var BBasmak巴斯马克 feud.Barony = &巴斯马克BasmakBarony{}

func init() {
	f := BBasmak巴斯马克.(*巴斯马克BasmakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "basmak",
		TitleName: "巴斯马克",
		TitleCode: "b_basmak",
	}
}
