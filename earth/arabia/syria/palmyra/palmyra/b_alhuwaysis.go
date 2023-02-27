package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡韦西斯AlhuwaysisBarony struct {
	feud.BaseBarony
}

var BAlhuwaysis胡韦西斯 feud.Barony = &胡韦西斯AlhuwaysisBarony{}

func init() {
    f := BAlhuwaysis胡韦西斯.(*胡韦西斯AlhuwaysisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhuwaysis",
		TitleName: "胡韦西斯",
		TitleCode: "b_alhuwaysis",
	}
}
