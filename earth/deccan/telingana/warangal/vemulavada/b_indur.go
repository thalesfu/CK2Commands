package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因杜尔IndurBarony struct {
	feud.BaseBarony
}

var BIndur因杜尔 feud.Barony = &因杜尔IndurBarony{}

func init() {
	f := BIndur因杜尔.(*因杜尔IndurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "indur",
		TitleName: "因杜尔",
		TitleCode: "b_indur",
	}
}
