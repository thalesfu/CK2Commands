package annaba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnnabaCounty interface {
	feud.County
	BAnnabah安纳巴() feud.Barony
	BBoudjama布乌贾马() feud.Barony
	BDrean德雷安() feud.Barony
	BElbouni布乌尼() feud.Barony
	BEltarf塔里夫() feud.Barony
	BMeroudj梅罗贾() feud.Barony
	BSkikda斯基克达() feud.Barony
	BSoukahras苏克阿赫拉斯() feud.Barony
}

type 安纳巴AnnabaCounty struct {
	feud.BaseCounty
	安纳巴Annabah      feud.Barony
	布乌贾马Boudjama    feud.Barony
	德雷安Drean        feud.Barony
	布乌尼Elbouni      feud.Barony
	塔里夫Eltarf       feud.Barony
	梅罗贾Meroudj      feud.Barony
	斯基克达Skikda      feud.Barony
	苏克阿赫拉斯Soukahras feud.Barony
}

func (c *安纳巴AnnabaCounty) BAnnabah安纳巴() feud.Barony {
	return c.安纳巴Annabah
}

func (c *安纳巴AnnabaCounty) BBoudjama布乌贾马() feud.Barony {
	return c.布乌贾马Boudjama
}

func (c *安纳巴AnnabaCounty) BDrean德雷安() feud.Barony {
	return c.德雷安Drean
}

func (c *安纳巴AnnabaCounty) BElbouni布乌尼() feud.Barony {
	return c.布乌尼Elbouni
}

func (c *安纳巴AnnabaCounty) BEltarf塔里夫() feud.Barony {
	return c.塔里夫Eltarf
}

func (c *安纳巴AnnabaCounty) BMeroudj梅罗贾() feud.Barony {
	return c.梅罗贾Meroudj
}

func (c *安纳巴AnnabaCounty) BSkikda斯基克达() feud.Barony {
	return c.斯基克达Skikda
}

func (c *安纳巴AnnabaCounty) BSoukahras苏克阿赫拉斯() feud.Barony {
	return c.苏克阿赫拉斯Soukahras
}

var CAnnaba安纳巴 AnnabaCounty = &安纳巴AnnabaCounty{}

func init() {
	f := CAnnaba安纳巴.(*安纳巴AnnabaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "820",
		Title:     "annaba",
		TitleName: "安纳巴",
		TitleCode: "c_annaba",
		Baronies:  map[string]feud.Barony{},
	}

	f.安纳巴Annabah = BAnnabah安纳巴
	f.安纳巴Annabah.SetParent(f)

	f.布乌贾马Boudjama = BBoudjama布乌贾马
	f.布乌贾马Boudjama.SetParent(f)

	f.德雷安Drean = BDrean德雷安
	f.德雷安Drean.SetParent(f)

	f.布乌尼Elbouni = BElbouni布乌尼
	f.布乌尼Elbouni.SetParent(f)

	f.塔里夫Eltarf = BEltarf塔里夫
	f.塔里夫Eltarf.SetParent(f)

	f.梅罗贾Meroudj = BMeroudj梅罗贾
	f.梅罗贾Meroudj.SetParent(f)

	f.斯基克达Skikda = BSkikda斯基克达
	f.斯基克达Skikda.SetParent(f)

	f.苏克阿赫拉斯Soukahras = BSoukahras苏克阿赫拉斯
	f.苏克阿赫拉斯Soukahras.SetParent(f)

}
