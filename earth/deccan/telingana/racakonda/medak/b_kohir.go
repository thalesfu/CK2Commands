package medak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科赫尔KohirBarony struct {
	feud.BaseBarony
}

var BKohir科赫尔 feud.Barony = &科赫尔KohirBarony{}

func init() {
	f := BKohir科赫尔.(*科赫尔KohirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kohir",
		TitleName: "科赫尔",
		TitleCode: "b_kohir",
	}
}
