package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TagantCounty interface {
	feud.County
	BApowa阿波瓦() feud.Barony
	BBagble巴格布勒() feud.Barony
	BDawba道巴() feud.Barony
	BNyinuito尼努伊托() feud.Barony
	BTagant塔甘特() feud.Barony
	BTibogona蒂博戈纳() feud.Barony
	BTichitt提希特() feud.Barony
}

type 塔甘特TagantCounty struct {
	feud.BaseCounty
	阿波瓦Apowa     feud.Barony
	巴格布勒Bagble   feud.Barony
	道巴Dawba      feud.Barony
	尼努伊托Nyinuito feud.Barony
	塔甘特Tagant    feud.Barony
	蒂博戈纳Tibogona feud.Barony
	提希特Tichitt   feud.Barony
}

func (c *塔甘特TagantCounty) BApowa阿波瓦() feud.Barony {
	return c.阿波瓦Apowa
}

func (c *塔甘特TagantCounty) BBagble巴格布勒() feud.Barony {
	return c.巴格布勒Bagble
}

func (c *塔甘特TagantCounty) BDawba道巴() feud.Barony {
	return c.道巴Dawba
}

func (c *塔甘特TagantCounty) BNyinuito尼努伊托() feud.Barony {
	return c.尼努伊托Nyinuito
}

func (c *塔甘特TagantCounty) BTagant塔甘特() feud.Barony {
	return c.塔甘特Tagant
}

func (c *塔甘特TagantCounty) BTibogona蒂博戈纳() feud.Barony {
	return c.蒂博戈纳Tibogona
}

func (c *塔甘特TagantCounty) BTichitt提希特() feud.Barony {
	return c.提希特Tichitt
}

var CTagant塔甘特 TagantCounty = &塔甘特TagantCounty{}

func init() {
	f := CTagant塔甘特.(*塔甘特TagantCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1762",
		Title:     "tagant",
		TitleName: "塔甘特",
		TitleCode: "c_tagant",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿波瓦Apowa = BApowa阿波瓦
	f.阿波瓦Apowa.SetParent(f)

	f.巴格布勒Bagble = BBagble巴格布勒
	f.巴格布勒Bagble.SetParent(f)

	f.道巴Dawba = BDawba道巴
	f.道巴Dawba.SetParent(f)

	f.尼努伊托Nyinuito = BNyinuito尼努伊托
	f.尼努伊托Nyinuito.SetParent(f)

	f.塔甘特Tagant = BTagant塔甘特
	f.塔甘特Tagant.SetParent(f)

	f.蒂博戈纳Tibogona = BTibogona蒂博戈纳
	f.蒂博戈纳Tibogona.SetParent(f)

	f.提希特Tichitt = BTichitt提希特
	f.提希特Tichitt.SetParent(f)

}
