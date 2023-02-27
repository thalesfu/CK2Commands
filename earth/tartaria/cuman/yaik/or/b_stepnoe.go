package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯捷普诺耶StepnoeBarony struct {
	feud.BaseBarony
}

var BStepnoe斯捷普诺耶 feud.Barony = &斯捷普诺耶StepnoeBarony{}

func init() {
    f := BStepnoe斯捷普诺耶.(*斯捷普诺耶StepnoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stepnoe",
		TitleName: "斯捷普诺耶",
		TitleCode: "b_stepnoe",
	}
}
