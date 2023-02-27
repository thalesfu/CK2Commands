package terek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯捷普诺耶StepnoyeBarony struct {
	feud.BaseBarony
}

var BStepnoye斯捷普诺耶 feud.Barony = &斯捷普诺耶StepnoyeBarony{}

func init() {
    f := BStepnoye斯捷普诺耶.(*斯捷普诺耶StepnoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stepnoye",
		TitleName: "斯捷普诺耶",
		TitleCode: "b_stepnoye",
	}
}
