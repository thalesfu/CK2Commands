package om

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯捷普诺伊StepnoyBarony struct {
	feud.BaseBarony
}

var BStepnoy斯捷普诺伊 feud.Barony = &斯捷普诺伊StepnoyBarony{}

func init() {
	f := BStepnoy斯捷普诺伊.(*斯捷普诺伊StepnoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stepnoy",
		TitleName: "斯捷普诺伊",
		TitleCode: "b_stepnoy",
	}
}
