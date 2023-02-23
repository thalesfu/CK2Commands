package belukha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BelukhaCounty interface {
	feud.County
	BAktach阿克塔奇() feud.Barony
	BBelukha别卢哈() feud.Barony
	BInya伊尼亚() feud.Barony
	BKuray库赖() feud.Barony
	BMulta穆利塔() feud.Barony
	BTyungur琼古尔() feud.Barony
	BYustik尤斯季克() feud.Barony
}

type 别卢哈BelukhaCounty struct {
	feud.BaseCounty
	阿克塔奇Aktach feud.Barony
	别卢哈Belukha feud.Barony
	伊尼亚Inya    feud.Barony
	库赖Kuray    feud.Barony
	穆利塔Multa   feud.Barony
	琼古尔Tyungur feud.Barony
	尤斯季克Yustik feud.Barony
}

func (c *别卢哈BelukhaCounty) BAktach阿克塔奇() feud.Barony {
	return c.阿克塔奇Aktach
}

func (c *别卢哈BelukhaCounty) BBelukha别卢哈() feud.Barony {
	return c.别卢哈Belukha
}

func (c *别卢哈BelukhaCounty) BInya伊尼亚() feud.Barony {
	return c.伊尼亚Inya
}

func (c *别卢哈BelukhaCounty) BKuray库赖() feud.Barony {
	return c.库赖Kuray
}

func (c *别卢哈BelukhaCounty) BMulta穆利塔() feud.Barony {
	return c.穆利塔Multa
}

func (c *别卢哈BelukhaCounty) BTyungur琼古尔() feud.Barony {
	return c.琼古尔Tyungur
}

func (c *别卢哈BelukhaCounty) BYustik尤斯季克() feud.Barony {
	return c.尤斯季克Yustik
}

var CBelukha别卢哈 BelukhaCounty = &别卢哈BelukhaCounty{}

func init() {
	f := CBelukha别卢哈.(*别卢哈BelukhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1895",
		Title:     "belukha",
		TitleName: "别卢哈",
		TitleCode: "c_belukha",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克塔奇Aktach = BAktach阿克塔奇
	f.阿克塔奇Aktach.SetParent(f)

	f.别卢哈Belukha = BBelukha别卢哈
	f.别卢哈Belukha.SetParent(f)

	f.伊尼亚Inya = BInya伊尼亚
	f.伊尼亚Inya.SetParent(f)

	f.库赖Kuray = BKuray库赖
	f.库赖Kuray.SetParent(f)

	f.穆利塔Multa = BMulta穆利塔
	f.穆利塔Multa.SetParent(f)

	f.琼古尔Tyungur = BTyungur琼古尔
	f.琼古尔Tyungur.SetParent(f)

	f.尤斯季克Yustik = BYustik尤斯季克
	f.尤斯季克Yustik.SetParent(f)

}
