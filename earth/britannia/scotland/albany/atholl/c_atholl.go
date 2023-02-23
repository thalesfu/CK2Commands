package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AthollCounty interface {
	feud.County
	BAnthrachs安施拉克斯() feud.Barony
	BFortingall福廷格尔() feud.Barony
	BGrandtully格兰特利() feud.Barony
	BPitlochry皮特洛赫里() feud.Barony
	BRannoch兰诺赫() feud.Barony
	BStrathardle斯特拉塔杜() feud.Barony
	BStruan斯特鲁恩() feud.Barony
}

type 阿瑟尔AthollCounty struct {
	feud.BaseCounty
	安施拉克斯Anthrachs   feud.Barony
	福廷格尔Fortingall   feud.Barony
	格兰特利Grandtully   feud.Barony
	皮特洛赫里Pitlochry   feud.Barony
	兰诺赫Rannoch       feud.Barony
	斯特拉塔杜Strathardle feud.Barony
	斯特鲁恩Struan       feud.Barony
}

func (c *阿瑟尔AthollCounty) BAnthrachs安施拉克斯() feud.Barony {
	return c.安施拉克斯Anthrachs
}

func (c *阿瑟尔AthollCounty) BFortingall福廷格尔() feud.Barony {
	return c.福廷格尔Fortingall
}

func (c *阿瑟尔AthollCounty) BGrandtully格兰特利() feud.Barony {
	return c.格兰特利Grandtully
}

func (c *阿瑟尔AthollCounty) BPitlochry皮特洛赫里() feud.Barony {
	return c.皮特洛赫里Pitlochry
}

func (c *阿瑟尔AthollCounty) BRannoch兰诺赫() feud.Barony {
	return c.兰诺赫Rannoch
}

func (c *阿瑟尔AthollCounty) BStrathardle斯特拉塔杜() feud.Barony {
	return c.斯特拉塔杜Strathardle
}

func (c *阿瑟尔AthollCounty) BStruan斯特鲁恩() feud.Barony {
	return c.斯特鲁恩Struan
}

var CAtholl阿瑟尔 AthollCounty = &阿瑟尔AthollCounty{}

func init() {
	f := CAtholl阿瑟尔.(*阿瑟尔AthollCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "44",
		Title:     "atholl",
		TitleName: "阿瑟尔",
		TitleCode: "c_atholl",
		Baronies:  map[string]feud.Barony{},
	}

	f.安施拉克斯Anthrachs = BAnthrachs安施拉克斯
	f.安施拉克斯Anthrachs.SetParent(f)

	f.福廷格尔Fortingall = BFortingall福廷格尔
	f.福廷格尔Fortingall.SetParent(f)

	f.格兰特利Grandtully = BGrandtully格兰特利
	f.格兰特利Grandtully.SetParent(f)

	f.皮特洛赫里Pitlochry = BPitlochry皮特洛赫里
	f.皮特洛赫里Pitlochry.SetParent(f)

	f.兰诺赫Rannoch = BRannoch兰诺赫
	f.兰诺赫Rannoch.SetParent(f)

	f.斯特拉塔杜Strathardle = BStrathardle斯特拉塔杜
	f.斯特拉塔杜Strathardle.SetParent(f)

	f.斯特鲁恩Struan = BStruan斯特鲁恩
	f.斯特鲁恩Struan.SetParent(f)

}
