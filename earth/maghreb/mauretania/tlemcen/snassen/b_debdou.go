package snassen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德卜杜DebdouBarony struct {
	feud.BaseBarony
}

var BDebdou德卜杜 feud.Barony = &德卜杜DebdouBarony{}

func init() {
	f := BDebdou德卜杜.(*德卜杜DebdouBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "debdou",
		TitleName: "德卜杜",
		TitleCode: "b_debdou",
	}
}
