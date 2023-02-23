package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EmpuriesCounty interface {
	feud.County
	BBanyoles巴尼奥莱斯() feud.Barony
	BBesalu贝萨卢() feud.Barony
	BCastelldaro阿罗堡() feud.Barony
	BCerdana塞尔达尼亚() feud.Barony
	BEmpuries恩普里耶斯() feud.Barony
	BFigueras菲格雷斯() feud.Barony
	BGirona赫罗纳() feud.Barony
	BLabisbaldemporda拉比斯瓦尔登波尔达() feud.Barony
}

type 恩普里耶斯EmpuriesCounty struct {
	feud.BaseCounty
	巴尼奥莱斯Banyoles             feud.Barony
	贝萨卢Besalu                 feud.Barony
	阿罗堡Castelldaro            feud.Barony
	塞尔达尼亚Cerdana              feud.Barony
	恩普里耶斯Empuries             feud.Barony
	菲格雷斯Figueras              feud.Barony
	赫罗纳Girona                 feud.Barony
	拉比斯瓦尔登波尔达Labisbaldemporda feud.Barony
}

func (c *恩普里耶斯EmpuriesCounty) BBanyoles巴尼奥莱斯() feud.Barony {
	return c.巴尼奥莱斯Banyoles
}

func (c *恩普里耶斯EmpuriesCounty) BBesalu贝萨卢() feud.Barony {
	return c.贝萨卢Besalu
}

func (c *恩普里耶斯EmpuriesCounty) BCastelldaro阿罗堡() feud.Barony {
	return c.阿罗堡Castelldaro
}

func (c *恩普里耶斯EmpuriesCounty) BCerdana塞尔达尼亚() feud.Barony {
	return c.塞尔达尼亚Cerdana
}

func (c *恩普里耶斯EmpuriesCounty) BEmpuries恩普里耶斯() feud.Barony {
	return c.恩普里耶斯Empuries
}

func (c *恩普里耶斯EmpuriesCounty) BFigueras菲格雷斯() feud.Barony {
	return c.菲格雷斯Figueras
}

func (c *恩普里耶斯EmpuriesCounty) BGirona赫罗纳() feud.Barony {
	return c.赫罗纳Girona
}

func (c *恩普里耶斯EmpuriesCounty) BLabisbaldemporda拉比斯瓦尔登波尔达() feud.Barony {
	return c.拉比斯瓦尔登波尔达Labisbaldemporda
}

var CEmpuries恩普里耶斯 EmpuriesCounty = &恩普里耶斯EmpuriesCounty{}

func init() {
	f := CEmpuries恩普里耶斯.(*恩普里耶斯EmpuriesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "205",
		Title:     "empuries",
		TitleName: "恩普里耶斯",
		TitleCode: "c_empuries",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尼奥莱斯Banyoles = BBanyoles巴尼奥莱斯
	f.巴尼奥莱斯Banyoles.SetParent(f)

	f.贝萨卢Besalu = BBesalu贝萨卢
	f.贝萨卢Besalu.SetParent(f)

	f.阿罗堡Castelldaro = BCastelldaro阿罗堡
	f.阿罗堡Castelldaro.SetParent(f)

	f.塞尔达尼亚Cerdana = BCerdana塞尔达尼亚
	f.塞尔达尼亚Cerdana.SetParent(f)

	f.恩普里耶斯Empuries = BEmpuries恩普里耶斯
	f.恩普里耶斯Empuries.SetParent(f)

	f.菲格雷斯Figueras = BFigueras菲格雷斯
	f.菲格雷斯Figueras.SetParent(f)

	f.赫罗纳Girona = BGirona赫罗纳
	f.赫罗纳Girona.SetParent(f)

	f.拉比斯瓦尔登波尔达Labisbaldemporda = BLabisbaldemporda拉比斯瓦尔登波尔达
	f.拉比斯瓦尔登波尔达Labisbaldemporda.SetParent(f)

}
