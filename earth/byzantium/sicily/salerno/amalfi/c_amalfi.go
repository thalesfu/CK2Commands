package amalfi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmalfiCounty interface {
	feud.County
	BAmalfi阿马尔菲() feud.Barony
	BCapri卡普里() feud.Barony
	BCastellamare海堡() feud.Barony
	BMinori米诺丽() feud.Barony
	BPositano波西塔诺() feud.Barony
	BRavello拉韦洛() feud.Barony
	BSorrento索伦托() feud.Barony
	BTramonti特拉蒙蒂() feud.Barony
}

type 阿马尔菲AmalfiCounty struct {
	feud.BaseCounty
	阿马尔菲Amalfi     feud.Barony
	卡普里Capri       feud.Barony
	海堡Castellamare feud.Barony
	米诺丽Minori      feud.Barony
	波西塔诺Positano   feud.Barony
	拉韦洛Ravello     feud.Barony
	索伦托Sorrento    feud.Barony
	特拉蒙蒂Tramonti   feud.Barony
}

func (c *阿马尔菲AmalfiCounty) BAmalfi阿马尔菲() feud.Barony {
	return c.阿马尔菲Amalfi
}

func (c *阿马尔菲AmalfiCounty) BCapri卡普里() feud.Barony {
	return c.卡普里Capri
}

func (c *阿马尔菲AmalfiCounty) BCastellamare海堡() feud.Barony {
	return c.海堡Castellamare
}

func (c *阿马尔菲AmalfiCounty) BMinori米诺丽() feud.Barony {
	return c.米诺丽Minori
}

func (c *阿马尔菲AmalfiCounty) BPositano波西塔诺() feud.Barony {
	return c.波西塔诺Positano
}

func (c *阿马尔菲AmalfiCounty) BRavello拉韦洛() feud.Barony {
	return c.拉韦洛Ravello
}

func (c *阿马尔菲AmalfiCounty) BSorrento索伦托() feud.Barony {
	return c.索伦托Sorrento
}

func (c *阿马尔菲AmalfiCounty) BTramonti特拉蒙蒂() feud.Barony {
	return c.特拉蒙蒂Tramonti
}

var CAmalfi阿马尔菲 AmalfiCounty = &阿马尔菲AmalfiCounty{}

func init() {
	f := CAmalfi阿马尔菲.(*阿马尔菲AmalfiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "935",
		Title:     "amalfi",
		TitleName: "阿马尔菲",
		TitleCode: "c_amalfi",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马尔菲Amalfi = BAmalfi阿马尔菲
	f.阿马尔菲Amalfi.SetParent(f)

	f.卡普里Capri = BCapri卡普里
	f.卡普里Capri.SetParent(f)

	f.海堡Castellamare = BCastellamare海堡
	f.海堡Castellamare.SetParent(f)

	f.米诺丽Minori = BMinori米诺丽
	f.米诺丽Minori.SetParent(f)

	f.波西塔诺Positano = BPositano波西塔诺
	f.波西塔诺Positano.SetParent(f)

	f.拉韦洛Ravello = BRavello拉韦洛
	f.拉韦洛Ravello.SetParent(f)

	f.索伦托Sorrento = BSorrento索伦托
	f.索伦托Sorrento.SetParent(f)

	f.特拉蒙蒂Tramonti = BTramonti特拉蒙蒂
	f.特拉蒙蒂Tramonti.SetParent(f)

}
