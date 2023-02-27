package khinjali_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那耶姞利呬NayagarhBarony struct {
	feud.BaseBarony
}

var BNayagarh那耶姞利呬 feud.Barony = &那耶姞利呬NayagarhBarony{}

func init() {
    f := BNayagarh那耶姞利呬.(*那耶姞利呬NayagarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nayagarh",
		TitleName: "那耶姞利呬",
		TitleCode: "b_nayagarh",
	}
}
