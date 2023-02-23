package limisol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LimisolCounty interface {
	feud.County
	BAgridi阿格里迪() feud.Barony
	BArsinoe阿尔西诺伊() feud.Barony
	BDieudamour迪奥达摩() feud.Barony
	BKhirokitia基罗基蒂亚() feud.Barony
	BKolossi科洛西() feud.Barony
	BLimmassol利马索尔() feud.Barony
	BMorphou莫尔富() feud.Barony
	BPaphos帕福斯() feud.Barony
}

type 利马索尔LimisolCounty struct {
	feud.BaseCounty
	阿格里迪Agridi      feud.Barony
	阿尔西诺伊Arsinoe    feud.Barony
	迪奥达摩Dieudamour  feud.Barony
	基罗基蒂亚Khirokitia feud.Barony
	科洛西Kolossi      feud.Barony
	利马索尔Limmassol   feud.Barony
	莫尔富Morphou      feud.Barony
	帕福斯Paphos       feud.Barony
}

func (c *利马索尔LimisolCounty) BAgridi阿格里迪() feud.Barony {
	return c.阿格里迪Agridi
}

func (c *利马索尔LimisolCounty) BArsinoe阿尔西诺伊() feud.Barony {
	return c.阿尔西诺伊Arsinoe
}

func (c *利马索尔LimisolCounty) BDieudamour迪奥达摩() feud.Barony {
	return c.迪奥达摩Dieudamour
}

func (c *利马索尔LimisolCounty) BKhirokitia基罗基蒂亚() feud.Barony {
	return c.基罗基蒂亚Khirokitia
}

func (c *利马索尔LimisolCounty) BKolossi科洛西() feud.Barony {
	return c.科洛西Kolossi
}

func (c *利马索尔LimisolCounty) BLimmassol利马索尔() feud.Barony {
	return c.利马索尔Limmassol
}

func (c *利马索尔LimisolCounty) BMorphou莫尔富() feud.Barony {
	return c.莫尔富Morphou
}

func (c *利马索尔LimisolCounty) BPaphos帕福斯() feud.Barony {
	return c.帕福斯Paphos
}

var CLimisol利马索尔 LimisolCounty = &利马索尔LimisolCounty{}

func init() {
	f := CLimisol利马索尔.(*利马索尔LimisolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "756",
		Title:     "limisol",
		TitleName: "利马索尔",
		TitleCode: "c_limisol",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格里迪Agridi = BAgridi阿格里迪
	f.阿格里迪Agridi.SetParent(f)

	f.阿尔西诺伊Arsinoe = BArsinoe阿尔西诺伊
	f.阿尔西诺伊Arsinoe.SetParent(f)

	f.迪奥达摩Dieudamour = BDieudamour迪奥达摩
	f.迪奥达摩Dieudamour.SetParent(f)

	f.基罗基蒂亚Khirokitia = BKhirokitia基罗基蒂亚
	f.基罗基蒂亚Khirokitia.SetParent(f)

	f.科洛西Kolossi = BKolossi科洛西
	f.科洛西Kolossi.SetParent(f)

	f.利马索尔Limmassol = BLimmassol利马索尔
	f.利马索尔Limmassol.SetParent(f)

	f.莫尔富Morphou = BMorphou莫尔富
	f.莫尔富Morphou.SetParent(f)

	f.帕福斯Paphos = BPaphos帕福斯
	f.帕福斯Paphos.SetParent(f)

}
