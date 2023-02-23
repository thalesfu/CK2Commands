package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦地那MedinaBarony struct {
	feud.BaseBarony
}

var BMedina麦地那 feud.Barony = &麦地那MedinaBarony{}

func init() {
	f := BMedina麦地那.(*麦地那MedinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medina",
		TitleName: "麦地那",
		TitleCode: "b_medina",
	}
}
