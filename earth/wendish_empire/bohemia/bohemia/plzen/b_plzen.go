package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普尔曾PlzenBarony struct {
	feud.BaseBarony
}

var BPlzen普尔曾 feud.Barony = &普尔曾PlzenBarony{}

func init() {
    f := BPlzen普尔曾.(*普尔曾PlzenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plzen",
		TitleName: "普尔曾",
		TitleCode: "b_plzen",
	}
}
