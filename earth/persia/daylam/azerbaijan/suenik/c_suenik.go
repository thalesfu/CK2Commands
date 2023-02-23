package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuenikCounty interface {
	feud.County
	BAreni阿雷尼() feud.Barony
	BGandzassar甘扎萨尔() feud.Barony
	BGhapan卡拉恩() feud.Barony
	BGoris戈里斯() feud.Barony
	BNoravank诺拉万克() feud.Barony
	BProchabert拉洛斯比尔特() feud.Barony
	BTatev塔特夫() feud.Barony
	BVorotnavank沃罗特纳万克() feud.Barony
}

type 休尼克SuenikCounty struct {
	feud.BaseCounty
	阿雷尼Areni          feud.Barony
	甘扎萨尔Gandzassar    feud.Barony
	卡拉恩Ghapan         feud.Barony
	戈里斯Goris          feud.Barony
	诺拉万克Noravank      feud.Barony
	拉洛斯比尔特Prochabert  feud.Barony
	塔特夫Tatev          feud.Barony
	沃罗特纳万克Vorotnavank feud.Barony
}

func (c *休尼克SuenikCounty) BAreni阿雷尼() feud.Barony {
	return c.阿雷尼Areni
}

func (c *休尼克SuenikCounty) BGandzassar甘扎萨尔() feud.Barony {
	return c.甘扎萨尔Gandzassar
}

func (c *休尼克SuenikCounty) BGhapan卡拉恩() feud.Barony {
	return c.卡拉恩Ghapan
}

func (c *休尼克SuenikCounty) BGoris戈里斯() feud.Barony {
	return c.戈里斯Goris
}

func (c *休尼克SuenikCounty) BNoravank诺拉万克() feud.Barony {
	return c.诺拉万克Noravank
}

func (c *休尼克SuenikCounty) BProchabert拉洛斯比尔特() feud.Barony {
	return c.拉洛斯比尔特Prochabert
}

func (c *休尼克SuenikCounty) BTatev塔特夫() feud.Barony {
	return c.塔特夫Tatev
}

func (c *休尼克SuenikCounty) BVorotnavank沃罗特纳万克() feud.Barony {
	return c.沃罗特纳万克Vorotnavank
}

var CSuenik休尼克 SuenikCounty = &休尼克SuenikCounty{}

func init() {
	f := CSuenik休尼克.(*休尼克SuenikCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "671",
		Title:     "suenik",
		TitleName: "休尼克",
		TitleCode: "c_suenik",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿雷尼Areni = BAreni阿雷尼
	f.阿雷尼Areni.SetParent(f)

	f.甘扎萨尔Gandzassar = BGandzassar甘扎萨尔
	f.甘扎萨尔Gandzassar.SetParent(f)

	f.卡拉恩Ghapan = BGhapan卡拉恩
	f.卡拉恩Ghapan.SetParent(f)

	f.戈里斯Goris = BGoris戈里斯
	f.戈里斯Goris.SetParent(f)

	f.诺拉万克Noravank = BNoravank诺拉万克
	f.诺拉万克Noravank.SetParent(f)

	f.拉洛斯比尔特Prochabert = BProchabert拉洛斯比尔特
	f.拉洛斯比尔特Prochabert.SetParent(f)

	f.塔特夫Tatev = BTatev塔特夫
	f.塔特夫Tatev.SetParent(f)

	f.沃罗特纳万克Vorotnavank = BVorotnavank沃罗特纳万克
	f.沃罗特纳万克Vorotnavank.SetParent(f)

}
