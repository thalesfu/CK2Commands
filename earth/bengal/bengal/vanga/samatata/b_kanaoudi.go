package samatata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽那奥提KanaoudiBarony struct {
	feud.BaseBarony
}

var BKanaoudi伽那奥提 feud.Barony = &伽那奥提KanaoudiBarony{}

func init() {
    f := BKanaoudi伽那奥提.(*伽那奥提KanaoudiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanaoudi",
		TitleName: "伽那奥提",
		TitleCode: "b_kanaoudi",
	}
}
