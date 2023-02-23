package abyssinia

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbyssiniaEmpire interface {
	feud.Empire
	KAbyssinia阿比西尼亚() abyssinia.AbyssiniaKingdom
	KNubia努比亚() nubia.NubiaKingdom
}

type 阿比西尼亚AbyssiniaEmpire struct {
	feud.BaseEmpire
	阿比西尼亚Abyssinia abyssinia.AbyssiniaKingdom
	努比亚Nubia       nubia.NubiaKingdom
}

func (e *阿比西尼亚AbyssiniaEmpire) KAbyssinia阿比西尼亚() abyssinia.AbyssiniaKingdom {
	return e.阿比西尼亚Abyssinia
}

func (e *阿比西尼亚AbyssiniaEmpire) KNubia努比亚() nubia.NubiaKingdom {
	return e.努比亚Nubia
}

var EAbyssinia阿比西尼亚 AbyssiniaEmpire = &阿比西尼亚AbyssiniaEmpire{}

func init() {
	f := EAbyssinia阿比西尼亚.(*阿比西尼亚AbyssiniaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "abyssinia",
		TitleName: "阿比西尼亚",
		TitleCode: "e_abyssinia",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.阿比西尼亚Abyssinia = abyssinia.KAbyssinia阿比西尼亚
	f.阿比西尼亚Abyssinia.SetParent(f)
	f.努比亚Nubia = nubia.KNubia努比亚
	f.努比亚Nubia.SetParent(f)
}
