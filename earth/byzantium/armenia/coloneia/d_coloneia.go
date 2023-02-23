package coloneia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/coloneia/koloneia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/coloneia/melitene"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ColoneiaDuke interface {
	feud.Duke
	CKoloneia科洛尼亚() koloneia.KoloneiaCounty
	CMelitene梅利蒂尼() melitene.MeliteneCounty
}

type 科洛尼亚ColoneiaDuke struct {
	feud.BaseDuke
	科洛尼亚Koloneia koloneia.KoloneiaCounty
	梅利蒂尼Melitene melitene.MeliteneCounty
}

func (d *科洛尼亚ColoneiaDuke) CKoloneia科洛尼亚() koloneia.KoloneiaCounty {
	return d.科洛尼亚Koloneia
}

func (d *科洛尼亚ColoneiaDuke) CMelitene梅利蒂尼() melitene.MeliteneCounty {
	return d.梅利蒂尼Melitene
}

var DColoneia科洛尼亚 ColoneiaDuke = &科洛尼亚ColoneiaDuke{}

func init() {
	f := DColoneia科洛尼亚.(*科洛尼亚ColoneiaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "coloneia",
		TitleName: "科洛尼亚",
		TitleCode: "d_coloneia",
		Counties:  map[string]feud.County{},
	}

	f.科洛尼亚Koloneia = koloneia.CKoloneia科洛尼亚
	f.科洛尼亚Koloneia.SetParent(f)

	f.梅利蒂尼Melitene = melitene.CMelitene梅利蒂尼
	f.梅利蒂尼Melitene.SetParent(f)

}
