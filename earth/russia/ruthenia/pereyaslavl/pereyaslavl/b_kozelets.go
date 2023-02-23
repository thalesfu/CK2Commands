package pereyaslavl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科泽列齐KozeletsBarony struct {
	feud.BaseBarony
}

var BKozelets科泽列齐 feud.Barony = &科泽列齐KozeletsBarony{}

func init() {
	f := BKozelets科泽列齐.(*科泽列齐KozeletsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozelets",
		TitleName: "科泽列齐",
		TitleCode: "b_kozelets",
	}
}
