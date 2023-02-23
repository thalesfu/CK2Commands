package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米特福德MitfordBarony struct {
	feud.BaseBarony
}

var BMitford米特福德 feud.Barony = &米特福德MitfordBarony{}

func init() {
	f := BMitford米特福德.(*米特福德MitfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mitford",
		TitleName: "米特福德",
		TitleCode: "b_mitford",
	}
}
