package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里什顿RishtonBarony struct {
	feud.BaseBarony
}

var BRishton里什顿 feud.Barony = &里什顿RishtonBarony{}

func init() {
    f := BRishton里什顿.(*里什顿RishtonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rishton",
		TitleName: "里什顿",
		TitleCode: "b_rishton",
	}
}
