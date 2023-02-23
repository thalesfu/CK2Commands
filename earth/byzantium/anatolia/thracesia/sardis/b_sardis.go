package sardis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨第斯SardisBarony struct {
	feud.BaseBarony
}

var BSardis萨第斯 feud.Barony = &萨第斯SardisBarony{}

func init() {
	f := BSardis萨第斯.(*萨第斯SardisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sardis",
		TitleName: "萨第斯",
		TitleCode: "b_sardis",
	}
}
