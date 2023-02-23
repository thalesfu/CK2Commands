package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OravaCounty interface {
	feud.County
	BArvavara阿尔沃堡() feud.Barony
	BLiptovskymikulas利普托夫斯基米库拉什() feud.Barony
	BNameszto纳迈斯托() feud.Barony
	BNemetlipcse内迈特利普切() feud.Barony
	BRozsahegy罗饶海吉() feud.Barony
	BTrsztena特尔斯泰纳() feud.Barony
	BTurdossin图尔多欣() feud.Barony
	BZolyom佐约姆() feud.Barony
}

type 兹沃伦OravaCounty struct {
	feud.BaseCounty
	阿尔沃堡Arvavara               feud.Barony
	利普托夫斯基米库拉什Liptovskymikulas feud.Barony
	纳迈斯托Nameszto               feud.Barony
	内迈特利普切Nemetlipcse          feud.Barony
	罗饶海吉Rozsahegy              feud.Barony
	特尔斯泰纳Trsztena              feud.Barony
	图尔多欣Turdossin              feud.Barony
	佐约姆Zolyom                  feud.Barony
}

func (c *兹沃伦OravaCounty) BArvavara阿尔沃堡() feud.Barony {
	return c.阿尔沃堡Arvavara
}

func (c *兹沃伦OravaCounty) BLiptovskymikulas利普托夫斯基米库拉什() feud.Barony {
	return c.利普托夫斯基米库拉什Liptovskymikulas
}

func (c *兹沃伦OravaCounty) BNameszto纳迈斯托() feud.Barony {
	return c.纳迈斯托Nameszto
}

func (c *兹沃伦OravaCounty) BNemetlipcse内迈特利普切() feud.Barony {
	return c.内迈特利普切Nemetlipcse
}

func (c *兹沃伦OravaCounty) BRozsahegy罗饶海吉() feud.Barony {
	return c.罗饶海吉Rozsahegy
}

func (c *兹沃伦OravaCounty) BTrsztena特尔斯泰纳() feud.Barony {
	return c.特尔斯泰纳Trsztena
}

func (c *兹沃伦OravaCounty) BTurdossin图尔多欣() feud.Barony {
	return c.图尔多欣Turdossin
}

func (c *兹沃伦OravaCounty) BZolyom佐约姆() feud.Barony {
	return c.佐约姆Zolyom
}

var COrava兹沃伦 OravaCounty = &兹沃伦OravaCounty{}

func init() {
	f := COrava兹沃伦.(*兹沃伦OravaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "525",
		Title:     "orava",
		TitleName: "兹沃伦",
		TitleCode: "c_orava",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔沃堡Arvavara = BArvavara阿尔沃堡
	f.阿尔沃堡Arvavara.SetParent(f)

	f.利普托夫斯基米库拉什Liptovskymikulas = BLiptovskymikulas利普托夫斯基米库拉什
	f.利普托夫斯基米库拉什Liptovskymikulas.SetParent(f)

	f.纳迈斯托Nameszto = BNameszto纳迈斯托
	f.纳迈斯托Nameszto.SetParent(f)

	f.内迈特利普切Nemetlipcse = BNemetlipcse内迈特利普切
	f.内迈特利普切Nemetlipcse.SetParent(f)

	f.罗饶海吉Rozsahegy = BRozsahegy罗饶海吉
	f.罗饶海吉Rozsahegy.SetParent(f)

	f.特尔斯泰纳Trsztena = BTrsztena特尔斯泰纳
	f.特尔斯泰纳Trsztena.SetParent(f)

	f.图尔多欣Turdossin = BTurdossin图尔多欣
	f.图尔多欣Turdossin.SetParent(f)

	f.佐约姆Zolyom = BZolyom佐约姆
	f.佐约姆Zolyom.SetParent(f)

}
