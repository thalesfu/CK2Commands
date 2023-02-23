package nordland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NordlandCounty interface {
	feud.County
	BAndenes安德内斯() feud.Barony
	BBeiarn贝亚恩() feud.Barony
	BBodo博德() feud.Barony
	BHarstad哈尔斯塔() feud.Barony
	BKabelvag卡伯尔沃格() feud.Barony
	BNarvik纳尔维克() feud.Barony
	BRodoy勒德于() feud.Barony
	BRost勒斯特() feud.Barony
}

type 诺德兰NordlandCounty struct {
	feud.BaseCounty
	安德内斯Andenes   feud.Barony
	贝亚恩Beiarn     feud.Barony
	博德Bodo        feud.Barony
	哈尔斯塔Harstad   feud.Barony
	卡伯尔沃格Kabelvag feud.Barony
	纳尔维克Narvik    feud.Barony
	勒德于Rodoy      feud.Barony
	勒斯特Rost       feud.Barony
}

func (c *诺德兰NordlandCounty) BAndenes安德内斯() feud.Barony {
	return c.安德内斯Andenes
}

func (c *诺德兰NordlandCounty) BBeiarn贝亚恩() feud.Barony {
	return c.贝亚恩Beiarn
}

func (c *诺德兰NordlandCounty) BBodo博德() feud.Barony {
	return c.博德Bodo
}

func (c *诺德兰NordlandCounty) BHarstad哈尔斯塔() feud.Barony {
	return c.哈尔斯塔Harstad
}

func (c *诺德兰NordlandCounty) BKabelvag卡伯尔沃格() feud.Barony {
	return c.卡伯尔沃格Kabelvag
}

func (c *诺德兰NordlandCounty) BNarvik纳尔维克() feud.Barony {
	return c.纳尔维克Narvik
}

func (c *诺德兰NordlandCounty) BRodoy勒德于() feud.Barony {
	return c.勒德于Rodoy
}

func (c *诺德兰NordlandCounty) BRost勒斯特() feud.Barony {
	return c.勒斯特Rost
}

var CNordland诺德兰 NordlandCounty = &诺德兰NordlandCounty{}

func init() {
	f := CNordland诺德兰.(*诺德兰NordlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "391",
		Title:     "nordland",
		TitleName: "诺德兰",
		TitleCode: "c_nordland",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德内斯Andenes = BAndenes安德内斯
	f.安德内斯Andenes.SetParent(f)

	f.贝亚恩Beiarn = BBeiarn贝亚恩
	f.贝亚恩Beiarn.SetParent(f)

	f.博德Bodo = BBodo博德
	f.博德Bodo.SetParent(f)

	f.哈尔斯塔Harstad = BHarstad哈尔斯塔
	f.哈尔斯塔Harstad.SetParent(f)

	f.卡伯尔沃格Kabelvag = BKabelvag卡伯尔沃格
	f.卡伯尔沃格Kabelvag.SetParent(f)

	f.纳尔维克Narvik = BNarvik纳尔维克
	f.纳尔维克Narvik.SetParent(f)

	f.勒德于Rodoy = BRodoy勒德于
	f.勒德于Rodoy.SetParent(f)

	f.勒斯特Rost = BRost勒斯特
	f.勒斯特Rost.SetParent(f)

}
