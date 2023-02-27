package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜伊诺DuinoBarony struct {
	feud.BaseBarony
}

var BDuino杜伊诺 feud.Barony = &杜伊诺DuinoBarony{}

func init() {
    f := BDuino杜伊诺.(*杜伊诺DuinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duino",
		TitleName: "杜伊诺",
		TitleCode: "b_duino",
	}
}
