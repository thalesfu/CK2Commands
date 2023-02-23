package koln

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马克MarkBarony struct {
	feud.BaseBarony
}

var BMark马克 feud.Barony = &马克MarkBarony{}

func init() {
	f := BMark马克.(*马克MarkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mark",
		TitleName: "马克",
		TitleCode: "b_mark",
	}
}
