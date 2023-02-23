package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaurCounty interface {
	feud.County
	BDair达伊尔() feud.Barony
	BKensay坑萨依() feud.Barony
	BKogeday库库岱() feud.Barony
	BSaur萨吾尔() feud.Barony
	BTughyl吐格勒() feud.Barony
}

type 萨吾尔SaurCounty struct {
	feud.BaseCounty
	达伊尔Dair    feud.Barony
	坑萨依Kensay  feud.Barony
	库库岱Kogeday feud.Barony
	萨吾尔Saur    feud.Barony
	吐格勒Tughyl  feud.Barony
}

func (c *萨吾尔SaurCounty) BDair达伊尔() feud.Barony {
	return c.达伊尔Dair
}

func (c *萨吾尔SaurCounty) BKensay坑萨依() feud.Barony {
	return c.坑萨依Kensay
}

func (c *萨吾尔SaurCounty) BKogeday库库岱() feud.Barony {
	return c.库库岱Kogeday
}

func (c *萨吾尔SaurCounty) BSaur萨吾尔() feud.Barony {
	return c.萨吾尔Saur
}

func (c *萨吾尔SaurCounty) BTughyl吐格勒() feud.Barony {
	return c.吐格勒Tughyl
}

var CSaur萨吾尔 SaurCounty = &萨吾尔SaurCounty{}

func init() {
	f := CSaur萨吾尔.(*萨吾尔SaurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1874",
		Title:     "saur",
		TitleName: "萨吾尔",
		TitleCode: "c_saur",
		Baronies:  map[string]feud.Barony{},
	}

	f.达伊尔Dair = BDair达伊尔
	f.达伊尔Dair.SetParent(f)

	f.坑萨依Kensay = BKensay坑萨依
	f.坑萨依Kensay.SetParent(f)

	f.库库岱Kogeday = BKogeday库库岱
	f.库库岱Kogeday.SetParent(f)

	f.萨吾尔Saur = BSaur萨吾尔
	f.萨吾尔Saur.SetParent(f)

	f.吐格勒Tughyl = BTughyl吐格勒
	f.吐格勒Tughyl.SetParent(f)

}
