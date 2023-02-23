package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波谛罗PateraBarony struct {
	feud.BaseBarony
}

var BPatera波谛罗 feud.Barony = &波谛罗PateraBarony{}

func init() {
	f := BPatera波谛罗.(*波谛罗PateraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "patera",
		TitleName: "波谛罗",
		TitleCode: "b_patera",
	}
}
