package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarisCounty interface {
	feud.County
	BBartfa巴尔特褔() feud.Barony
	BEperjes艾派尔耶什() feud.Barony
	BGiralth吉拉尔特() feud.Barony
	BHethars海塔尔什() feud.Barony
	BKisszeben基什塞本() feud.Barony
	BLemesany莱梅沙尼() feud.Barony
	BSaris斯皮什() feud.Barony
	BScyuidnyk斯奎德尼克() feud.Barony
}

type 斯皮什SarisCounty struct {
	feud.BaseCounty
	巴尔特褔Bartfa     feud.Barony
	艾派尔耶什Eperjes   feud.Barony
	吉拉尔特Giralth    feud.Barony
	海塔尔什Hethars    feud.Barony
	基什塞本Kisszeben  feud.Barony
	莱梅沙尼Lemesany   feud.Barony
	斯皮什Saris       feud.Barony
	斯奎德尼克Scyuidnyk feud.Barony
}

func (c *斯皮什SarisCounty) BBartfa巴尔特褔() feud.Barony {
	return c.巴尔特褔Bartfa
}

func (c *斯皮什SarisCounty) BEperjes艾派尔耶什() feud.Barony {
	return c.艾派尔耶什Eperjes
}

func (c *斯皮什SarisCounty) BGiralth吉拉尔特() feud.Barony {
	return c.吉拉尔特Giralth
}

func (c *斯皮什SarisCounty) BHethars海塔尔什() feud.Barony {
	return c.海塔尔什Hethars
}

func (c *斯皮什SarisCounty) BKisszeben基什塞本() feud.Barony {
	return c.基什塞本Kisszeben
}

func (c *斯皮什SarisCounty) BLemesany莱梅沙尼() feud.Barony {
	return c.莱梅沙尼Lemesany
}

func (c *斯皮什SarisCounty) BSaris斯皮什() feud.Barony {
	return c.斯皮什Saris
}

func (c *斯皮什SarisCounty) BScyuidnyk斯奎德尼克() feud.Barony {
	return c.斯奎德尼克Scyuidnyk
}

var CSaris斯皮什 SarisCounty = &斯皮什SarisCounty{}

func init() {
	f := CSaris斯皮什.(*斯皮什SarisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "533",
		Title:     "saris",
		TitleName: "斯皮什",
		TitleCode: "c_saris",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔特褔Bartfa = BBartfa巴尔特褔
	f.巴尔特褔Bartfa.SetParent(f)

	f.艾派尔耶什Eperjes = BEperjes艾派尔耶什
	f.艾派尔耶什Eperjes.SetParent(f)

	f.吉拉尔特Giralth = BGiralth吉拉尔特
	f.吉拉尔特Giralth.SetParent(f)

	f.海塔尔什Hethars = BHethars海塔尔什
	f.海塔尔什Hethars.SetParent(f)

	f.基什塞本Kisszeben = BKisszeben基什塞本
	f.基什塞本Kisszeben.SetParent(f)

	f.莱梅沙尼Lemesany = BLemesany莱梅沙尼
	f.莱梅沙尼Lemesany.SetParent(f)

	f.斯皮什Saris = BSaris斯皮什
	f.斯皮什Saris.SetParent(f)

	f.斯奎德尼克Scyuidnyk = BScyuidnyk斯奎德尼克
	f.斯奎德尼克Scyuidnyk.SetParent(f)

}
