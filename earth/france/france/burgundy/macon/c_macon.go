package macon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaconCounty interface {
	feud.County
	BBeaujeu博热() feud.Barony
	BBerze贝尔泽() feud.Barony
	BCluny克吕尼() feud.Barony
	BDavaye达瓦耶() feud.Barony
	BFuisse菲塞() feud.Barony
	BLugny吕尼() feud.Barony
	BMacon马孔() feud.Barony
	BVillefranchesursaone索恩河畔自由城() feud.Barony
}

type 马孔MaconCounty struct {
	feud.BaseCounty
	博热Beaujeu                   feud.Barony
	贝尔泽Berze                    feud.Barony
	克吕尼Cluny                    feud.Barony
	达瓦耶Davaye                   feud.Barony
	菲塞Fuisse                    feud.Barony
	吕尼Lugny                     feud.Barony
	马孔Macon                     feud.Barony
	索恩河畔自由城Villefranchesursaone feud.Barony
}

func (c *马孔MaconCounty) BBeaujeu博热() feud.Barony {
	return c.博热Beaujeu
}

func (c *马孔MaconCounty) BBerze贝尔泽() feud.Barony {
	return c.贝尔泽Berze
}

func (c *马孔MaconCounty) BCluny克吕尼() feud.Barony {
	return c.克吕尼Cluny
}

func (c *马孔MaconCounty) BDavaye达瓦耶() feud.Barony {
	return c.达瓦耶Davaye
}

func (c *马孔MaconCounty) BFuisse菲塞() feud.Barony {
	return c.菲塞Fuisse
}

func (c *马孔MaconCounty) BLugny吕尼() feud.Barony {
	return c.吕尼Lugny
}

func (c *马孔MaconCounty) BMacon马孔() feud.Barony {
	return c.马孔Macon
}

func (c *马孔MaconCounty) BVillefranchesursaone索恩河畔自由城() feud.Barony {
	return c.索恩河畔自由城Villefranchesursaone
}

var CMacon马孔 MaconCounty = &马孔MaconCounty{}

func init() {
	f := CMacon马孔.(*马孔MaconCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "225",
		Title:     "macon",
		TitleName: "马孔",
		TitleCode: "c_macon",
		Baronies:  map[string]feud.Barony{},
	}

	f.博热Beaujeu = BBeaujeu博热
	f.博热Beaujeu.SetParent(f)

	f.贝尔泽Berze = BBerze贝尔泽
	f.贝尔泽Berze.SetParent(f)

	f.克吕尼Cluny = BCluny克吕尼
	f.克吕尼Cluny.SetParent(f)

	f.达瓦耶Davaye = BDavaye达瓦耶
	f.达瓦耶Davaye.SetParent(f)

	f.菲塞Fuisse = BFuisse菲塞
	f.菲塞Fuisse.SetParent(f)

	f.吕尼Lugny = BLugny吕尼
	f.吕尼Lugny.SetParent(f)

	f.马孔Macon = BMacon马孔
	f.马孔Macon.SetParent(f)

	f.索恩河畔自由城Villefranchesursaone = BVillefranchesursaone索恩河畔自由城
	f.索恩河畔自由城Villefranchesursaone.SetParent(f)

}
