package trinacria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrinacriaKingdom interface {
	feud.Kingdom
}

type 特里纳克里亚TrinacriaKingdom struct {
	feud.BaseKingdom
}

var KTrinacria特里纳克里亚 TrinacriaKingdom = &特里纳克里亚TrinacriaKingdom{}

func init() {
	f := KTrinacria特里纳克里亚.(*特里纳克里亚TrinacriaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "trinacria",
		TitleName: "特里纳克里亚",
		TitleCode: "k_trinacria",
		Dukes:     map[string]feud.Duke{},
	}

}
