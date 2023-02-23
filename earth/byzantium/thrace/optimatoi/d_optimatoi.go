package optimatoi

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/optimatoi/herakleia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/optimatoi/nikomedeia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OptimatoiDuke interface {
	feud.Duke
	CHerakleia赫拉克利亚() herakleia.HerakleiaCounty
	CNikomedeia尼科米底亚() nikomedeia.NikomedeiaCounty
}

type 奥普提马通OptimatoiDuke struct {
	feud.BaseDuke
	赫拉克利亚Herakleia  herakleia.HerakleiaCounty
	尼科米底亚Nikomedeia nikomedeia.NikomedeiaCounty
}

func (d *奥普提马通OptimatoiDuke) CHerakleia赫拉克利亚() herakleia.HerakleiaCounty {
	return d.赫拉克利亚Herakleia
}

func (d *奥普提马通OptimatoiDuke) CNikomedeia尼科米底亚() nikomedeia.NikomedeiaCounty {
	return d.尼科米底亚Nikomedeia
}

var DOptimatoi奥普提马通 OptimatoiDuke = &奥普提马通OptimatoiDuke{}

func init() {
	f := DOptimatoi奥普提马通.(*奥普提马通OptimatoiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "optimatoi",
		TitleName: "奥普提马通",
		TitleCode: "d_optimatoi",
		Counties:  map[string]feud.County{},
	}

	f.赫拉克利亚Herakleia = herakleia.CHerakleia赫拉克利亚
	f.赫拉克利亚Herakleia.SetParent(f)

	f.尼科米底亚Nikomedeia = nikomedeia.CNikomedeia尼科米底亚
	f.尼科米底亚Nikomedeia.SetParent(f)

}
