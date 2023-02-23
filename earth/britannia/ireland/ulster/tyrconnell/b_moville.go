package tyrconnell

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫维尔MovilleBarony struct {
	feud.BaseBarony
}

var BMoville莫维尔 feud.Barony = &莫维尔MovilleBarony{}

func init() {
	f := BMoville莫维尔.(*莫维尔MovilleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moville",
		TitleName: "莫维尔",
		TitleCode: "b_moville",
	}
}
