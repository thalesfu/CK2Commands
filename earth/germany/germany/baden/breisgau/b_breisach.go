package breisgau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赖萨赫BreisachBarony struct {
	feud.BaseBarony
}

var BBreisach布赖萨赫 feud.Barony = &布赖萨赫BreisachBarony{}

func init() {
    f := BBreisach布赖萨赫.(*布赖萨赫BreisachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "breisach",
		TitleName: "布赖萨赫",
		TitleCode: "b_breisach",
	}
}
