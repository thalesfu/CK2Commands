package rama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RamaCounty interface {
	feud.County
	BBanovici巴诺维契() feud.Barony
	BGradacac格拉达查茨() feud.Barony
	BKladanj克拉达尼() feud.Barony
	BOlovo奥洛沃() feud.Barony
	BSoli索利() feud.Barony
	BSrebrenica斯雷布雷尼察() feud.Barony
	BZvonik兹沃尼克() feud.Barony
}

type 索利RamaCounty struct {
	feud.BaseCounty
	巴诺维契Banovici     feud.Barony
	格拉达查茨Gradacac    feud.Barony
	克拉达尼Kladanj      feud.Barony
	奥洛沃Olovo         feud.Barony
	索利Soli           feud.Barony
	斯雷布雷尼察Srebrenica feud.Barony
	兹沃尼克Zvonik       feud.Barony
}

func (c *索利RamaCounty) BBanovici巴诺维契() feud.Barony {
	return c.巴诺维契Banovici
}

func (c *索利RamaCounty) BGradacac格拉达查茨() feud.Barony {
	return c.格拉达查茨Gradacac
}

func (c *索利RamaCounty) BKladanj克拉达尼() feud.Barony {
	return c.克拉达尼Kladanj
}

func (c *索利RamaCounty) BOlovo奥洛沃() feud.Barony {
	return c.奥洛沃Olovo
}

func (c *索利RamaCounty) BSoli索利() feud.Barony {
	return c.索利Soli
}

func (c *索利RamaCounty) BSrebrenica斯雷布雷尼察() feud.Barony {
	return c.斯雷布雷尼察Srebrenica
}

func (c *索利RamaCounty) BZvonik兹沃尼克() feud.Barony {
	return c.兹沃尼克Zvonik
}

var CRama索利 RamaCounty = &索利RamaCounty{}

func init() {
	f := CRama索利.(*索利RamaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "504",
		Title:     "rama",
		TitleName: "索利",
		TitleCode: "c_rama",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴诺维契Banovici = BBanovici巴诺维契
	f.巴诺维契Banovici.SetParent(f)

	f.格拉达查茨Gradacac = BGradacac格拉达查茨
	f.格拉达查茨Gradacac.SetParent(f)

	f.克拉达尼Kladanj = BKladanj克拉达尼
	f.克拉达尼Kladanj.SetParent(f)

	f.奥洛沃Olovo = BOlovo奥洛沃
	f.奥洛沃Olovo.SetParent(f)

	f.索利Soli = BSoli索利
	f.索利Soli.SetParent(f)

	f.斯雷布雷尼察Srebrenica = BSrebrenica斯雷布雷尼察
	f.斯雷布雷尼察Srebrenica.SetParent(f)

	f.兹沃尼克Zvonik = BZvonik兹沃尼克
	f.兹沃尼克Zvonik.SetParent(f)

}
