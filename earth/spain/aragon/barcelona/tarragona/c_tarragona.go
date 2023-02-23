package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TarragonaCounty interface {
	feud.County
	BAmposta安波斯塔() feud.Barony
	BCambrils坎布里尔斯() feud.Barony
	BMontblanc蒙布朗() feud.Barony
	BReus雷乌斯() feud.Barony
	BSancugat圣库加() feud.Barony
	BSpantortosa托尔托萨() feud.Barony
	BTarragona塔拉戈纳() feud.Barony
	BVendrell本德雷利() feud.Barony
}

type 塔拉戈纳TarragonaCounty struct {
	feud.BaseCounty
	安波斯塔Amposta     feud.Barony
	坎布里尔斯Cambrils   feud.Barony
	蒙布朗Montblanc    feud.Barony
	雷乌斯Reus         feud.Barony
	圣库加Sancugat     feud.Barony
	托尔托萨Spantortosa feud.Barony
	塔拉戈纳Tarragona   feud.Barony
	本德雷利Vendrell    feud.Barony
}

func (c *塔拉戈纳TarragonaCounty) BAmposta安波斯塔() feud.Barony {
	return c.安波斯塔Amposta
}

func (c *塔拉戈纳TarragonaCounty) BCambrils坎布里尔斯() feud.Barony {
	return c.坎布里尔斯Cambrils
}

func (c *塔拉戈纳TarragonaCounty) BMontblanc蒙布朗() feud.Barony {
	return c.蒙布朗Montblanc
}

func (c *塔拉戈纳TarragonaCounty) BReus雷乌斯() feud.Barony {
	return c.雷乌斯Reus
}

func (c *塔拉戈纳TarragonaCounty) BSancugat圣库加() feud.Barony {
	return c.圣库加Sancugat
}

func (c *塔拉戈纳TarragonaCounty) BSpantortosa托尔托萨() feud.Barony {
	return c.托尔托萨Spantortosa
}

func (c *塔拉戈纳TarragonaCounty) BTarragona塔拉戈纳() feud.Barony {
	return c.塔拉戈纳Tarragona
}

func (c *塔拉戈纳TarragonaCounty) BVendrell本德雷利() feud.Barony {
	return c.本德雷利Vendrell
}

var CTarragona塔拉戈纳 TarragonaCounty = &塔拉戈纳TarragonaCounty{}

func init() {
	f := CTarragona塔拉戈纳.(*塔拉戈纳TarragonaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "173",
		Title:     "tarragona",
		TitleName: "塔拉戈纳",
		TitleCode: "c_tarragona",
		Baronies:  map[string]feud.Barony{},
	}

	f.安波斯塔Amposta = BAmposta安波斯塔
	f.安波斯塔Amposta.SetParent(f)

	f.坎布里尔斯Cambrils = BCambrils坎布里尔斯
	f.坎布里尔斯Cambrils.SetParent(f)

	f.蒙布朗Montblanc = BMontblanc蒙布朗
	f.蒙布朗Montblanc.SetParent(f)

	f.雷乌斯Reus = BReus雷乌斯
	f.雷乌斯Reus.SetParent(f)

	f.圣库加Sancugat = BSancugat圣库加
	f.圣库加Sancugat.SetParent(f)

	f.托尔托萨Spantortosa = BSpantortosa托尔托萨
	f.托尔托萨Spantortosa.SetParent(f)

	f.塔拉戈纳Tarragona = BTarragona塔拉戈纳
	f.塔拉戈纳Tarragona.SetParent(f)

	f.本德雷利Vendrell = BVendrell本德雷利
	f.本德雷利Vendrell.SetParent(f)

}
