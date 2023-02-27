package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克莱普KleppBarony struct {
	feud.BaseBarony
}

var BKlepp克莱普 feud.Barony = &克莱普KleppBarony{}

func init() {
    f := BKlepp克莱普.(*克莱普KleppBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klepp",
		TitleName: "克莱普",
		TitleCode: "b_klepp",
	}
}
