package dihistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DihistanCounty interface {
	feud.County
	BAkhur阿胡尔() feud.Barony
	BBayatkhadzi巴亚特_哈济() feud.Barony
	BGasankuli埃先古雷() feud.Barony
	BKemir克米尔() feud.Barony
	BKumdag古姆达格() feud.Barony
	BTorskhali图尔沙克利() feud.Barony
	BYarymtyk亚勒姆特克() feud.Barony
	BYasydepe亚瑟杰佩() feud.Barony
}

type 大益斯坦DihistanCounty struct {
	feud.BaseCounty
	阿胡尔Akhur          feud.Barony
	巴亚特_哈济Bayatkhadzi feud.Barony
	埃先古雷Gasankuli     feud.Barony
	克米尔Kemir          feud.Barony
	古姆达格Kumdag        feud.Barony
	图尔沙克利Torskhali    feud.Barony
	亚勒姆特克Yarymtyk     feud.Barony
	亚瑟杰佩Yasydepe      feud.Barony
}

func (c *大益斯坦DihistanCounty) BAkhur阿胡尔() feud.Barony {
	return c.阿胡尔Akhur
}

func (c *大益斯坦DihistanCounty) BBayatkhadzi巴亚特_哈济() feud.Barony {
	return c.巴亚特_哈济Bayatkhadzi
}

func (c *大益斯坦DihistanCounty) BGasankuli埃先古雷() feud.Barony {
	return c.埃先古雷Gasankuli
}

func (c *大益斯坦DihistanCounty) BKemir克米尔() feud.Barony {
	return c.克米尔Kemir
}

func (c *大益斯坦DihistanCounty) BKumdag古姆达格() feud.Barony {
	return c.古姆达格Kumdag
}

func (c *大益斯坦DihistanCounty) BTorskhali图尔沙克利() feud.Barony {
	return c.图尔沙克利Torskhali
}

func (c *大益斯坦DihistanCounty) BYarymtyk亚勒姆特克() feud.Barony {
	return c.亚勒姆特克Yarymtyk
}

func (c *大益斯坦DihistanCounty) BYasydepe亚瑟杰佩() feud.Barony {
	return c.亚瑟杰佩Yasydepe
}

var CDihistan大益斯坦 DihistanCounty = &大益斯坦DihistanCounty{}

func init() {
	f := CDihistan大益斯坦.(*大益斯坦DihistanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "631",
		Title:     "dihistan",
		TitleName: "大益斯坦",
		TitleCode: "c_dihistan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿胡尔Akhur = BAkhur阿胡尔
	f.阿胡尔Akhur.SetParent(f)

	f.巴亚特_哈济Bayatkhadzi = BBayatkhadzi巴亚特_哈济
	f.巴亚特_哈济Bayatkhadzi.SetParent(f)

	f.埃先古雷Gasankuli = BGasankuli埃先古雷
	f.埃先古雷Gasankuli.SetParent(f)

	f.克米尔Kemir = BKemir克米尔
	f.克米尔Kemir.SetParent(f)

	f.古姆达格Kumdag = BKumdag古姆达格
	f.古姆达格Kumdag.SetParent(f)

	f.图尔沙克利Torskhali = BTorskhali图尔沙克利
	f.图尔沙克利Torskhali.SetParent(f)

	f.亚勒姆特克Yarymtyk = BYarymtyk亚勒姆特克
	f.亚勒姆特克Yarymtyk.SetParent(f)

	f.亚瑟杰佩Yasydepe = BYasydepe亚瑟杰佩
	f.亚瑟杰佩Yasydepe.SetParent(f)

}
