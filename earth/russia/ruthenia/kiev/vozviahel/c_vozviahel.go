package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VozviahelCounty interface {
	feud.County
	BBaranivka巴拉尼夫卡() feud.Barony
	BBerdytchiv别尔基切夫() feud.Barony
	BHorodnytsia戈罗德尼齐亚() feud.Barony
	BHrud格鲁德() feud.Barony
	BKorets科列茨() feud.Barony
	BVozvyahel沃兹维亚赫利() feud.Barony
	BYemilchyne埃米利奇涅() feud.Barony
}

type 沃兹维亚格尔VozviahelCounty struct {
	feud.BaseCounty
	巴拉尼夫卡Baranivka    feud.Barony
	别尔基切夫Berdytchiv   feud.Barony
	戈罗德尼齐亚Horodnytsia feud.Barony
	格鲁德Hrud           feud.Barony
	科列茨Korets         feud.Barony
	沃兹维亚赫利Vozvyahel   feud.Barony
	埃米利奇涅Yemilchyne   feud.Barony
}

func (c *沃兹维亚格尔VozviahelCounty) BBaranivka巴拉尼夫卡() feud.Barony {
	return c.巴拉尼夫卡Baranivka
}

func (c *沃兹维亚格尔VozviahelCounty) BBerdytchiv别尔基切夫() feud.Barony {
	return c.别尔基切夫Berdytchiv
}

func (c *沃兹维亚格尔VozviahelCounty) BHorodnytsia戈罗德尼齐亚() feud.Barony {
	return c.戈罗德尼齐亚Horodnytsia
}

func (c *沃兹维亚格尔VozviahelCounty) BHrud格鲁德() feud.Barony {
	return c.格鲁德Hrud
}

func (c *沃兹维亚格尔VozviahelCounty) BKorets科列茨() feud.Barony {
	return c.科列茨Korets
}

func (c *沃兹维亚格尔VozviahelCounty) BVozvyahel沃兹维亚赫利() feud.Barony {
	return c.沃兹维亚赫利Vozvyahel
}

func (c *沃兹维亚格尔VozviahelCounty) BYemilchyne埃米利奇涅() feud.Barony {
	return c.埃米利奇涅Yemilchyne
}

var CVozviahel沃兹维亚格尔 VozviahelCounty = &沃兹维亚格尔VozviahelCounty{}

func init() {
	f := CVozviahel沃兹维亚格尔.(*沃兹维亚格尔VozviahelCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1651",
		Title:     "vozviahel",
		TitleName: "沃兹维亚格尔",
		TitleCode: "c_vozviahel",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴拉尼夫卡Baranivka = BBaranivka巴拉尼夫卡
	f.巴拉尼夫卡Baranivka.SetParent(f)

	f.别尔基切夫Berdytchiv = BBerdytchiv别尔基切夫
	f.别尔基切夫Berdytchiv.SetParent(f)

	f.戈罗德尼齐亚Horodnytsia = BHorodnytsia戈罗德尼齐亚
	f.戈罗德尼齐亚Horodnytsia.SetParent(f)

	f.格鲁德Hrud = BHrud格鲁德
	f.格鲁德Hrud.SetParent(f)

	f.科列茨Korets = BKorets科列茨
	f.科列茨Korets.SetParent(f)

	f.沃兹维亚赫利Vozvyahel = BVozvyahel沃兹维亚赫利
	f.沃兹维亚赫利Vozvyahel.SetParent(f)

	f.埃米利奇涅Yemilchyne = BYemilchyne埃米利奇涅
	f.埃米利奇涅Yemilchyne.SetParent(f)

}
