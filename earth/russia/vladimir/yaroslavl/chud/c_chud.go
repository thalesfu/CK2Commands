package chud

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChudCounty interface {
	feud.County
	BChunlovka春洛夫卡() feud.Barony
	BKamchuga卡姆丘加() feud.Barony
	BKrbor克里博尔() feud.Barony
	BKrutayaosyp克鲁塔亚奥瑟皮() feud.Barony
	BTotma托季马() feud.Barony
	BVeldvor韦利德沃尔() feud.Barony
	BZalese扎列西耶() feud.Barony
	BZaytsevo宰采沃() feud.Barony
}

type 楚德ChudCounty struct {
	feud.BaseCounty
	春洛夫卡Chunlovka      feud.Barony
	卡姆丘加Kamchuga       feud.Barony
	克里博尔Krbor          feud.Barony
	克鲁塔亚奥瑟皮Krutayaosyp feud.Barony
	托季马Totma           feud.Barony
	韦利德沃尔Veldvor       feud.Barony
	扎列西耶Zalese         feud.Barony
	宰采沃Zaytsevo        feud.Barony
}

func (c *楚德ChudCounty) BChunlovka春洛夫卡() feud.Barony {
	return c.春洛夫卡Chunlovka
}

func (c *楚德ChudCounty) BKamchuga卡姆丘加() feud.Barony {
	return c.卡姆丘加Kamchuga
}

func (c *楚德ChudCounty) BKrbor克里博尔() feud.Barony {
	return c.克里博尔Krbor
}

func (c *楚德ChudCounty) BKrutayaosyp克鲁塔亚奥瑟皮() feud.Barony {
	return c.克鲁塔亚奥瑟皮Krutayaosyp
}

func (c *楚德ChudCounty) BTotma托季马() feud.Barony {
	return c.托季马Totma
}

func (c *楚德ChudCounty) BVeldvor韦利德沃尔() feud.Barony {
	return c.韦利德沃尔Veldvor
}

func (c *楚德ChudCounty) BZalese扎列西耶() feud.Barony {
	return c.扎列西耶Zalese
}

func (c *楚德ChudCounty) BZaytsevo宰采沃() feud.Barony {
	return c.宰采沃Zaytsevo
}

var CChud楚德 ChudCounty = &楚德ChudCounty{}

func init() {
	f := CChud楚德.(*楚德ChudCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "405",
		Title:     "chud",
		TitleName: "楚德",
		TitleCode: "c_chud",
		Baronies:  map[string]feud.Barony{},
	}

	f.春洛夫卡Chunlovka = BChunlovka春洛夫卡
	f.春洛夫卡Chunlovka.SetParent(f)

	f.卡姆丘加Kamchuga = BKamchuga卡姆丘加
	f.卡姆丘加Kamchuga.SetParent(f)

	f.克里博尔Krbor = BKrbor克里博尔
	f.克里博尔Krbor.SetParent(f)

	f.克鲁塔亚奥瑟皮Krutayaosyp = BKrutayaosyp克鲁塔亚奥瑟皮
	f.克鲁塔亚奥瑟皮Krutayaosyp.SetParent(f)

	f.托季马Totma = BTotma托季马
	f.托季马Totma.SetParent(f)

	f.韦利德沃尔Veldvor = BVeldvor韦利德沃尔
	f.韦利德沃尔Veldvor.SetParent(f)

	f.扎列西耶Zalese = BZalese扎列西耶
	f.扎列西耶Zalese.SetParent(f)

	f.宰采沃Zaytsevo = BZaytsevo宰采沃
	f.宰采沃Zaytsevo.SetParent(f)

}
