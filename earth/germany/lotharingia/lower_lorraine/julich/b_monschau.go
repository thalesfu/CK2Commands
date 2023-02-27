package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙绍MonschauBarony struct {
	feud.BaseBarony
}

var BMonschau蒙绍 feud.Barony = &蒙绍MonschauBarony{}

func init() {
    f := BMonschau蒙绍.(*蒙绍MonschauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monschau",
		TitleName: "蒙绍",
		TitleCode: "b_monschau",
	}
}
