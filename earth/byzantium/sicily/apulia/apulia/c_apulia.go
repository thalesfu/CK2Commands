package apulia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ApuliaCounty interface {
	feud.County
	BBarletta巴尔莱塔() feud.Barony
	BCannae坎尼() feud.Barony
	BLavello拉韦洛() feud.Barony
	BLucano卢卡诺() feud.Barony
	BMelfi梅尔菲() feud.Barony
	BMinervo米内尔沃() feud.Barony
	BSalapia萨拉皮亚() feud.Barony
	BTrani特拉尼() feud.Barony
}

type 阿普利亚ApuliaCounty struct {
	feud.BaseCounty
	巴尔莱塔Barletta feud.Barony
	坎尼Cannae     feud.Barony
	拉韦洛Lavello   feud.Barony
	卢卡诺Lucano    feud.Barony
	梅尔菲Melfi     feud.Barony
	米内尔沃Minervo  feud.Barony
	萨拉皮亚Salapia  feud.Barony
	特拉尼Trani     feud.Barony
}

func (c *阿普利亚ApuliaCounty) BBarletta巴尔莱塔() feud.Barony {
	return c.巴尔莱塔Barletta
}

func (c *阿普利亚ApuliaCounty) BCannae坎尼() feud.Barony {
	return c.坎尼Cannae
}

func (c *阿普利亚ApuliaCounty) BLavello拉韦洛() feud.Barony {
	return c.拉韦洛Lavello
}

func (c *阿普利亚ApuliaCounty) BLucano卢卡诺() feud.Barony {
	return c.卢卡诺Lucano
}

func (c *阿普利亚ApuliaCounty) BMelfi梅尔菲() feud.Barony {
	return c.梅尔菲Melfi
}

func (c *阿普利亚ApuliaCounty) BMinervo米内尔沃() feud.Barony {
	return c.米内尔沃Minervo
}

func (c *阿普利亚ApuliaCounty) BSalapia萨拉皮亚() feud.Barony {
	return c.萨拉皮亚Salapia
}

func (c *阿普利亚ApuliaCounty) BTrani特拉尼() feud.Barony {
	return c.特拉尼Trani
}

var CApulia阿普利亚 ApuliaCounty = &阿普利亚ApuliaCounty{}

func init() {
	f := CApulia阿普利亚.(*阿普利亚ApuliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "347",
		Title:     "apulia",
		TitleName: "阿普利亚",
		TitleCode: "c_apulia",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔莱塔Barletta = BBarletta巴尔莱塔
	f.巴尔莱塔Barletta.SetParent(f)

	f.坎尼Cannae = BCannae坎尼
	f.坎尼Cannae.SetParent(f)

	f.拉韦洛Lavello = BLavello拉韦洛
	f.拉韦洛Lavello.SetParent(f)

	f.卢卡诺Lucano = BLucano卢卡诺
	f.卢卡诺Lucano.SetParent(f)

	f.梅尔菲Melfi = BMelfi梅尔菲
	f.梅尔菲Melfi.SetParent(f)

	f.米内尔沃Minervo = BMinervo米内尔沃
	f.米内尔沃Minervo.SetParent(f)

	f.萨拉皮亚Salapia = BSalapia萨拉皮亚
	f.萨拉皮亚Salapia.SetParent(f)

	f.特拉尼Trani = BTrani特拉尼
	f.特拉尼Trani.SetParent(f)

}
