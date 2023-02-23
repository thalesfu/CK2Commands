package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩加扎尔加姆NgazargamuBarony struct {
	feud.BaseBarony
}

var BNgazargamu恩加扎尔加姆 feud.Barony = &恩加扎尔加姆NgazargamuBarony{}

func init() {
	f := BNgazargamu恩加扎尔加姆.(*恩加扎尔加姆NgazargamuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ngazargamu",
		TitleName: "恩加扎尔加姆",
		TitleCode: "b_ngazargamu",
	}
}
