package manyakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩若契吒ManyakhetaBarony struct {
	feud.BaseBarony
}

var BManyakheta摩若契吒 feud.Barony = &摩若契吒ManyakhetaBarony{}

func init() {
	f := BManyakheta摩若契吒.(*摩若契吒ManyakhetaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manyakheta",
		TitleName: "摩若契吒",
		TitleCode: "b_manyakheta",
	}
}
