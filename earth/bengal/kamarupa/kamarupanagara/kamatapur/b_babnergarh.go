package kamatapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乏婆涅姞利呬BabnergarhBarony struct {
	feud.BaseBarony
}

var BBabnergarh乏婆涅姞利呬 feud.Barony = &乏婆涅姞利呬BabnergarhBarony{}

func init() {
	f := BBabnergarh乏婆涅姞利呬.(*乏婆涅姞利呬BabnergarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babnergarh",
		TitleName: "乏婆涅姞利呬",
		TitleCode: "b_babnergarh",
	}
}
