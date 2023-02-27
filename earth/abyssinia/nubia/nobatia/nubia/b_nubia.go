package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努比亚NubiaBarony struct {
	feud.BaseBarony
}

var BNubia努比亚 feud.Barony = &努比亚NubiaBarony{}

func init() {
    f := BNubia努比亚.(*努比亚NubiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nubia",
		TitleName: "努比亚",
		TitleCode: "b_nubia",
	}
}
