package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔梅索斯TelmissosBarony struct {
	feud.BaseBarony
}

var BTelmissos特尔梅索斯 feud.Barony = &特尔梅索斯TelmissosBarony{}

func init() {
	f := BTelmissos特尔梅索斯.(*特尔梅索斯TelmissosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "telmissos",
		TitleName: "特尔梅索斯",
		TitleCode: "b_telmissos",
	}
}
