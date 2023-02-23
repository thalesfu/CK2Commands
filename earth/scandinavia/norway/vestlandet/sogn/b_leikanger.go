package sogn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱康厄尔LeikangerBarony struct {
	feud.BaseBarony
}

var BLeikanger莱康厄尔 feud.Barony = &莱康厄尔LeikangerBarony{}

func init() {
	f := BLeikanger莱康厄尔.(*莱康厄尔LeikangerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leikanger",
		TitleName: "莱康厄尔",
		TitleCode: "b_leikanger",
	}
}
