package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨加拉索斯SagalassosBarony struct {
	feud.BaseBarony
}

var BSagalassos萨加拉索斯 feud.Barony = &萨加拉索斯SagalassosBarony{}

func init() {
    f := BSagalassos萨加拉索斯.(*萨加拉索斯SagalassosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sagalassos",
		TitleName: "萨加拉索斯",
		TitleCode: "b_sagalassos",
	}
}
