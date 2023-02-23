package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 郭勒木德GolmudBarony struct {
	feud.BaseBarony
}

var BGolmud郭勒木德 feud.Barony = &郭勒木德GolmudBarony{}

func init() {
	f := BGolmud郭勒木德.(*郭勒木德GolmudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "golmud",
		TitleName: "郭勒木德",
		TitleCode: "b_golmud",
	}
}
