package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里策比特尔RitzebuttelBarony struct {
	feud.BaseBarony
}

var BRitzebuttel里策比特尔 feud.Barony = &里策比特尔RitzebuttelBarony{}

func init() {
	f := BRitzebuttel里策比特尔.(*里策比特尔RitzebuttelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ritzebuttel",
		TitleName: "里策比特尔",
		TitleCode: "b_ritzebuttel",
	}
}
