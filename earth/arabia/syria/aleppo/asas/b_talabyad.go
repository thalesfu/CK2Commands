package asas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰勒艾卜耶德TalabyadBarony struct {
	feud.BaseBarony
}

var BTalabyad泰勒艾卜耶德 feud.Barony = &泰勒艾卜耶德TalabyadBarony{}

func init() {
    f := BTalabyad泰勒艾卜耶德.(*泰勒艾卜耶德TalabyadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talabyad",
		TitleName: "泰勒艾卜耶德",
		TitleCode: "b_talabyad",
	}
}
