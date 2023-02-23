package noli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NoliCounty interface {
	feud.County
	BAlaxio阿拉克西奥() feud.Barony
	BAlbenga阿尔本加() feud.Barony
	BNoli诺利() feud.Barony
	BOneglia奥内利亚() feud.Barony
	BSanremo圣雷莫() feud.Barony
	BSavona萨沃纳() feud.Barony
	BVentimiglia文蒂米利亚() feud.Barony
}

type 诺利NoliCounty struct {
	feud.BaseCounty
	阿拉克西奥Alaxio      feud.Barony
	阿尔本加Albenga      feud.Barony
	诺利Noli           feud.Barony
	奥内利亚Oneglia      feud.Barony
	圣雷莫Sanremo       feud.Barony
	萨沃纳Savona        feud.Barony
	文蒂米利亚Ventimiglia feud.Barony
}

func (c *诺利NoliCounty) BAlaxio阿拉克西奥() feud.Barony {
	return c.阿拉克西奥Alaxio
}

func (c *诺利NoliCounty) BAlbenga阿尔本加() feud.Barony {
	return c.阿尔本加Albenga
}

func (c *诺利NoliCounty) BNoli诺利() feud.Barony {
	return c.诺利Noli
}

func (c *诺利NoliCounty) BOneglia奥内利亚() feud.Barony {
	return c.奥内利亚Oneglia
}

func (c *诺利NoliCounty) BSanremo圣雷莫() feud.Barony {
	return c.圣雷莫Sanremo
}

func (c *诺利NoliCounty) BSavona萨沃纳() feud.Barony {
	return c.萨沃纳Savona
}

func (c *诺利NoliCounty) BVentimiglia文蒂米利亚() feud.Barony {
	return c.文蒂米利亚Ventimiglia
}

var CNoli诺利 NoliCounty = &诺利NoliCounty{}

func init() {
	f := CNoli诺利.(*诺利NoliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1714",
		Title:     "noli",
		TitleName: "诺利",
		TitleCode: "c_noli",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉克西奥Alaxio = BAlaxio阿拉克西奥
	f.阿拉克西奥Alaxio.SetParent(f)

	f.阿尔本加Albenga = BAlbenga阿尔本加
	f.阿尔本加Albenga.SetParent(f)

	f.诺利Noli = BNoli诺利
	f.诺利Noli.SetParent(f)

	f.奥内利亚Oneglia = BOneglia奥内利亚
	f.奥内利亚Oneglia.SetParent(f)

	f.圣雷莫Sanremo = BSanremo圣雷莫
	f.圣雷莫Sanremo.SetParent(f)

	f.萨沃纳Savona = BSavona萨沃纳
	f.萨沃纳Savona.SetParent(f)

	f.文蒂米利亚Ventimiglia = BVentimiglia文蒂米利亚
	f.文蒂米利亚Ventimiglia.SetParent(f)

}
