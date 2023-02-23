package epeiros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EpeirosCounty interface {
	feud.County
	BButrint布特林特() feud.Barony
	BGjirokaster吉罗卡斯特() feud.Barony
	BIgoumenitsa伊古迈尼察() feud.Barony
	BIoannina约阿尼纳() feud.Barony
	BParamythia帕拉米西亚() feud.Barony
	BPogonia波戈尼亚() feud.Barony
	BSagiada萨基亚达() feud.Barony
	BSopot索波特() feud.Barony
}

type 伊庇鲁斯EpeirosCounty struct {
	feud.BaseCounty
	布特林特Butrint      feud.Barony
	吉罗卡斯特Gjirokaster feud.Barony
	伊古迈尼察Igoumenitsa feud.Barony
	约阿尼纳Ioannina     feud.Barony
	帕拉米西亚Paramythia  feud.Barony
	波戈尼亚Pogonia      feud.Barony
	萨基亚达Sagiada      feud.Barony
	索波特Sopot         feud.Barony
}

func (c *伊庇鲁斯EpeirosCounty) BButrint布特林特() feud.Barony {
	return c.布特林特Butrint
}

func (c *伊庇鲁斯EpeirosCounty) BGjirokaster吉罗卡斯特() feud.Barony {
	return c.吉罗卡斯特Gjirokaster
}

func (c *伊庇鲁斯EpeirosCounty) BIgoumenitsa伊古迈尼察() feud.Barony {
	return c.伊古迈尼察Igoumenitsa
}

func (c *伊庇鲁斯EpeirosCounty) BIoannina约阿尼纳() feud.Barony {
	return c.约阿尼纳Ioannina
}

func (c *伊庇鲁斯EpeirosCounty) BParamythia帕拉米西亚() feud.Barony {
	return c.帕拉米西亚Paramythia
}

func (c *伊庇鲁斯EpeirosCounty) BPogonia波戈尼亚() feud.Barony {
	return c.波戈尼亚Pogonia
}

func (c *伊庇鲁斯EpeirosCounty) BSagiada萨基亚达() feud.Barony {
	return c.萨基亚达Sagiada
}

func (c *伊庇鲁斯EpeirosCounty) BSopot索波特() feud.Barony {
	return c.索波特Sopot
}

var CEpeiros伊庇鲁斯 EpeirosCounty = &伊庇鲁斯EpeirosCounty{}

func init() {
	f := CEpeiros伊庇鲁斯.(*伊庇鲁斯EpeirosCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "472",
		Title:     "epeiros",
		TitleName: "伊庇鲁斯",
		TitleCode: "c_epeiros",
		Baronies:  map[string]feud.Barony{},
	}

	f.布特林特Butrint = BButrint布特林特
	f.布特林特Butrint.SetParent(f)

	f.吉罗卡斯特Gjirokaster = BGjirokaster吉罗卡斯特
	f.吉罗卡斯特Gjirokaster.SetParent(f)

	f.伊古迈尼察Igoumenitsa = BIgoumenitsa伊古迈尼察
	f.伊古迈尼察Igoumenitsa.SetParent(f)

	f.约阿尼纳Ioannina = BIoannina约阿尼纳
	f.约阿尼纳Ioannina.SetParent(f)

	f.帕拉米西亚Paramythia = BParamythia帕拉米西亚
	f.帕拉米西亚Paramythia.SetParent(f)

	f.波戈尼亚Pogonia = BPogonia波戈尼亚
	f.波戈尼亚Pogonia.SetParent(f)

	f.萨基亚达Sagiada = BSagiada萨基亚达
	f.萨基亚达Sagiada.SetParent(f)

	f.索波特Sopot = BSopot索波特
	f.索波特Sopot.SetParent(f)

}
