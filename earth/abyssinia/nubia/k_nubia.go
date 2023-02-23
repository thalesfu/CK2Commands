package nubia

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/hayya"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nobatia"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/nubia"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/nubia/sennar"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NubiaKingdom interface {
	feud.Kingdom
	DHayya比林美伊亚() hayya.HayyaDuke
	DNobatia诺巴提亚() nobatia.NobatiaDuke
	DNubia马库里亚() nubia.NubiaDuke
	DSennar阿洛地亚() sennar.SennarDuke
}

type 努比亚NubiaKingdom struct {
	feud.BaseKingdom
	比林美伊亚Hayya  hayya.HayyaDuke
	诺巴提亚Nobatia nobatia.NobatiaDuke
	马库里亚Nubia   nubia.NubiaDuke
	阿洛地亚Sennar  sennar.SennarDuke
}

func (k *努比亚NubiaKingdom) DHayya比林美伊亚() hayya.HayyaDuke {
	return k.比林美伊亚Hayya
}

func (k *努比亚NubiaKingdom) DNobatia诺巴提亚() nobatia.NobatiaDuke {
	return k.诺巴提亚Nobatia
}

func (k *努比亚NubiaKingdom) DNubia马库里亚() nubia.NubiaDuke {
	return k.马库里亚Nubia
}

func (k *努比亚NubiaKingdom) DSennar阿洛地亚() sennar.SennarDuke {
	return k.阿洛地亚Sennar
}

var KNubia努比亚 NubiaKingdom = &努比亚NubiaKingdom{}

func init() {
	f := KNubia努比亚.(*努比亚NubiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "nubia",
		TitleName: "努比亚",
		TitleCode: "k_nubia",
		Dukes:     map[string]feud.Duke{},
	}

	f.比林美伊亚Hayya = hayya.DHayya比林美伊亚
	f.比林美伊亚Hayya.SetParent(f)

	f.诺巴提亚Nobatia = nobatia.DNobatia诺巴提亚
	f.诺巴提亚Nobatia.SetParent(f)

	f.马库里亚Nubia = nubia.DNubia马库里亚
	f.马库里亚Nubia.SetParent(f)

	f.阿洛地亚Sennar = sennar.DSennar阿洛地亚
	f.阿洛地亚Sennar.SetParent(f)

}
