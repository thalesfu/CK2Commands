package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯康AbercornBarony struct {
	feud.BaseBarony
}

var BAbercorn阿伯康 feud.Barony = &阿伯康AbercornBarony{}

func init() {
	f := BAbercorn阿伯康.(*阿伯康AbercornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abercorn",
		TitleName: "阿伯康",
		TitleCode: "b_abercorn",
	}
}
