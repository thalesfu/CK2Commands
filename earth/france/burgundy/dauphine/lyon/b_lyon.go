package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里昂LyonBarony struct {
	feud.BaseBarony
}

var BLyon里昂 feud.Barony = &里昂LyonBarony{}

func init() {
    f := BLyon里昂.(*里昂LyonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyon",
		TitleName: "里昂",
		TitleCode: "b_lyon",
	}
}
