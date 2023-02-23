package korosten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马林MalynBarony struct {
	feud.BaseBarony
}

var BMalyn马林 feud.Barony = &马林MalynBarony{}

func init() {
	f := BMalyn马林.(*马林MalynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malyn",
		TitleName: "马林",
		TitleCode: "b_malyn",
	}
}
