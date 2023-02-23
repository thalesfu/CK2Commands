package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍伦HoornBarony struct {
	feud.BaseBarony
}

var BHoorn霍伦 feud.Barony = &霍伦HoornBarony{}

func init() {
	f := BHoorn霍伦.(*霍伦HoornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hoorn",
		TitleName: "霍伦",
		TitleCode: "b_hoorn",
	}
}
