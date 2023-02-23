package denia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeniaCounty interface {
	feud.County
	BAlbatera阿尔瓦特拉() feud.Barony
	BAlicante阿拉坎特() feud.Barony
	BBenissa贝尼萨() feud.Barony
	BCastalla卡斯塔利亚() feud.Barony
	BDenia德尼亚() feud.Barony
	BElche埃利奇() feud.Barony
	BOrihuela奥里韦拉() feud.Barony
	BVillena比列纳() feud.Barony
}

type 德尼亚DeniaCounty struct {
	feud.BaseCounty
	阿尔瓦特拉Albatera feud.Barony
	阿拉坎特Alicante  feud.Barony
	贝尼萨Benissa    feud.Barony
	卡斯塔利亚Castalla feud.Barony
	德尼亚Denia      feud.Barony
	埃利奇Elche      feud.Barony
	奥里韦拉Orihuela  feud.Barony
	比列纳Villena    feud.Barony
}

func (c *德尼亚DeniaCounty) BAlbatera阿尔瓦特拉() feud.Barony {
	return c.阿尔瓦特拉Albatera
}

func (c *德尼亚DeniaCounty) BAlicante阿拉坎特() feud.Barony {
	return c.阿拉坎特Alicante
}

func (c *德尼亚DeniaCounty) BBenissa贝尼萨() feud.Barony {
	return c.贝尼萨Benissa
}

func (c *德尼亚DeniaCounty) BCastalla卡斯塔利亚() feud.Barony {
	return c.卡斯塔利亚Castalla
}

func (c *德尼亚DeniaCounty) BDenia德尼亚() feud.Barony {
	return c.德尼亚Denia
}

func (c *德尼亚DeniaCounty) BElche埃利奇() feud.Barony {
	return c.埃利奇Elche
}

func (c *德尼亚DeniaCounty) BOrihuela奥里韦拉() feud.Barony {
	return c.奥里韦拉Orihuela
}

func (c *德尼亚DeniaCounty) BVillena比列纳() feud.Barony {
	return c.比列纳Villena
}

var CDenia德尼亚 DeniaCounty = &德尼亚DeniaCounty{}

func init() {
	f := CDenia德尼亚.(*德尼亚DeniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "170",
		Title:     "denia",
		TitleName: "德尼亚",
		TitleCode: "c_denia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔瓦特拉Albatera = BAlbatera阿尔瓦特拉
	f.阿尔瓦特拉Albatera.SetParent(f)

	f.阿拉坎特Alicante = BAlicante阿拉坎特
	f.阿拉坎特Alicante.SetParent(f)

	f.贝尼萨Benissa = BBenissa贝尼萨
	f.贝尼萨Benissa.SetParent(f)

	f.卡斯塔利亚Castalla = BCastalla卡斯塔利亚
	f.卡斯塔利亚Castalla.SetParent(f)

	f.德尼亚Denia = BDenia德尼亚
	f.德尼亚Denia.SetParent(f)

	f.埃利奇Elche = BElche埃利奇
	f.埃利奇Elche.SetParent(f)

	f.奥里韦拉Orihuela = BOrihuela奥里韦拉
	f.奥里韦拉Orihuela.SetParent(f)

	f.比列纳Villena = BVillena比列纳
	f.比列纳Villena.SetParent(f)

}
