package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯泰鲁赫SterohBarony struct {
	feud.BaseBarony
}

var BSteroh斯泰鲁赫 feud.Barony = &斯泰鲁赫SterohBarony{}

func init() {
    f := BSteroh斯泰鲁赫.(*斯泰鲁赫SterohBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "steroh",
		TitleName: "斯泰鲁赫",
		TitleCode: "b_steroh",
	}
}
