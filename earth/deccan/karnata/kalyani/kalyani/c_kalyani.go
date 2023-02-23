package kalyani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KalyaniCounty interface {
	feud.County
	BAusa阿萨() feud.Barony
	BHanakuni诃那俱尼() feud.Barony
	BKalyani迦梨耶尼() feud.Barony
	BNarayanapur那罗衍那补罗() feud.Barony
	BSarsala萨娑罗() feud.Barony
	BSarseir萨西尔() feud.Barony
	BUmapur尤摩补罗() feud.Barony
}

type 迦梨耶尼KalyaniCounty struct {
	feud.BaseCounty
	阿萨Ausa            feud.Barony
	诃那俱尼Hanakuni      feud.Barony
	迦梨耶尼Kalyani       feud.Barony
	那罗衍那补罗Narayanapur feud.Barony
	萨娑罗Sarsala        feud.Barony
	萨西尔Sarseir        feud.Barony
	尤摩补罗Umapur        feud.Barony
}

func (c *迦梨耶尼KalyaniCounty) BAusa阿萨() feud.Barony {
	return c.阿萨Ausa
}

func (c *迦梨耶尼KalyaniCounty) BHanakuni诃那俱尼() feud.Barony {
	return c.诃那俱尼Hanakuni
}

func (c *迦梨耶尼KalyaniCounty) BKalyani迦梨耶尼() feud.Barony {
	return c.迦梨耶尼Kalyani
}

func (c *迦梨耶尼KalyaniCounty) BNarayanapur那罗衍那补罗() feud.Barony {
	return c.那罗衍那补罗Narayanapur
}

func (c *迦梨耶尼KalyaniCounty) BSarsala萨娑罗() feud.Barony {
	return c.萨娑罗Sarsala
}

func (c *迦梨耶尼KalyaniCounty) BSarseir萨西尔() feud.Barony {
	return c.萨西尔Sarseir
}

func (c *迦梨耶尼KalyaniCounty) BUmapur尤摩补罗() feud.Barony {
	return c.尤摩补罗Umapur
}

var CKalyani迦梨耶尼 KalyaniCounty = &迦梨耶尼KalyaniCounty{}

func init() {
	f := CKalyani迦梨耶尼.(*迦梨耶尼KalyaniCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1143",
		Title:     "kalyani",
		TitleName: "迦梨耶尼",
		TitleCode: "c_kalyani",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿萨Ausa = BAusa阿萨
	f.阿萨Ausa.SetParent(f)

	f.诃那俱尼Hanakuni = BHanakuni诃那俱尼
	f.诃那俱尼Hanakuni.SetParent(f)

	f.迦梨耶尼Kalyani = BKalyani迦梨耶尼
	f.迦梨耶尼Kalyani.SetParent(f)

	f.那罗衍那补罗Narayanapur = BNarayanapur那罗衍那补罗
	f.那罗衍那补罗Narayanapur.SetParent(f)

	f.萨娑罗Sarsala = BSarsala萨娑罗
	f.萨娑罗Sarsala.SetParent(f)

	f.萨西尔Sarseir = BSarseir萨西尔
	f.萨西尔Sarseir.SetParent(f)

	f.尤摩补罗Umapur = BUmapur尤摩补罗
	f.尤摩补罗Umapur.SetParent(f)

}
