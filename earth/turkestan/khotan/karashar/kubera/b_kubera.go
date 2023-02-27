package kubera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱毗罗KuberaBarony struct {
	feud.BaseBarony
}

var BKubera俱毗罗 feud.Barony = &俱毗罗KuberaBarony{}

func init() {
    f := BKubera俱毗罗.(*俱毗罗KuberaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kubera",
		TitleName: "俱毗罗",
		TitleCode: "b_kubera",
	}
}
