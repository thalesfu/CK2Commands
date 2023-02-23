package merv

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MervCounty interface {
	feud.County
	BBayramaly拜拉姆阿里() feud.Barony
	BGulanly古兰利() feud.Barony
	BKushka库什卡() feud.Barony
	BMerv木鹿() feud.Barony
	BSakarcage萨卡尔恰加() feud.Barony
	BTagtabazar塔格塔巴扎尔() feud.Barony
	BWekilbazar韦基尔巴扎尔() feud.Barony
	BYoloten约廖坚() feud.Barony
}

type 木鹿MervCounty struct {
	feud.BaseCounty
	拜拉姆阿里Bayramaly   feud.Barony
	古兰利Gulanly       feud.Barony
	库什卡Kushka        feud.Barony
	木鹿Merv           feud.Barony
	萨卡尔恰加Sakarcage   feud.Barony
	塔格塔巴扎尔Tagtabazar feud.Barony
	韦基尔巴扎尔Wekilbazar feud.Barony
	约廖坚Yoloten       feud.Barony
}

func (c *木鹿MervCounty) BBayramaly拜拉姆阿里() feud.Barony {
	return c.拜拉姆阿里Bayramaly
}

func (c *木鹿MervCounty) BGulanly古兰利() feud.Barony {
	return c.古兰利Gulanly
}

func (c *木鹿MervCounty) BKushka库什卡() feud.Barony {
	return c.库什卡Kushka
}

func (c *木鹿MervCounty) BMerv木鹿() feud.Barony {
	return c.木鹿Merv
}

func (c *木鹿MervCounty) BSakarcage萨卡尔恰加() feud.Barony {
	return c.萨卡尔恰加Sakarcage
}

func (c *木鹿MervCounty) BTagtabazar塔格塔巴扎尔() feud.Barony {
	return c.塔格塔巴扎尔Tagtabazar
}

func (c *木鹿MervCounty) BWekilbazar韦基尔巴扎尔() feud.Barony {
	return c.韦基尔巴扎尔Wekilbazar
}

func (c *木鹿MervCounty) BYoloten约廖坚() feud.Barony {
	return c.约廖坚Yoloten
}

var CMerv木鹿 MervCounty = &木鹿MervCounty{}

func init() {
	f := CMerv木鹿.(*木鹿MervCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "630",
		Title:     "merv",
		TitleName: "木鹿",
		TitleCode: "c_merv",
		Baronies:  map[string]feud.Barony{},
	}

	f.拜拉姆阿里Bayramaly = BBayramaly拜拉姆阿里
	f.拜拉姆阿里Bayramaly.SetParent(f)

	f.古兰利Gulanly = BGulanly古兰利
	f.古兰利Gulanly.SetParent(f)

	f.库什卡Kushka = BKushka库什卡
	f.库什卡Kushka.SetParent(f)

	f.木鹿Merv = BMerv木鹿
	f.木鹿Merv.SetParent(f)

	f.萨卡尔恰加Sakarcage = BSakarcage萨卡尔恰加
	f.萨卡尔恰加Sakarcage.SetParent(f)

	f.塔格塔巴扎尔Tagtabazar = BTagtabazar塔格塔巴扎尔
	f.塔格塔巴扎尔Tagtabazar.SetParent(f)

	f.韦基尔巴扎尔Wekilbazar = BWekilbazar韦基尔巴扎尔
	f.韦基尔巴扎尔Wekilbazar.SetParent(f)

	f.约廖坚Yoloten = BYoloten约廖坚
	f.约廖坚Yoloten.SetParent(f)

}
