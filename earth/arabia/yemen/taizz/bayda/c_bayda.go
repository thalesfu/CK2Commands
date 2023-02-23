package bayda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaydaCounty interface {
	feud.County
	BAlkhabr海卜尔() feud.Barony
	BAshshubaykah舒拜卡() feud.Barony
	BAtaq阿塔格() feud.Barony
	BBayda贝达() feud.Barony
	BHabban哈班() feud.Barony
	BTimna提姆纳() feud.Barony
	BYashbum耶什布姆() feud.Barony
	BZinjibar津吉巴尔() feud.Barony
}

type 贝达BaydaCounty struct {
	feud.BaseCounty
	海卜尔Alkhabr      feud.Barony
	舒拜卡Ashshubaykah feud.Barony
	阿塔格Ataq         feud.Barony
	贝达Bayda         feud.Barony
	哈班Habban        feud.Barony
	提姆纳Timna        feud.Barony
	耶什布姆Yashbum     feud.Barony
	津吉巴尔Zinjibar    feud.Barony
}

func (c *贝达BaydaCounty) BAlkhabr海卜尔() feud.Barony {
	return c.海卜尔Alkhabr
}

func (c *贝达BaydaCounty) BAshshubaykah舒拜卡() feud.Barony {
	return c.舒拜卡Ashshubaykah
}

func (c *贝达BaydaCounty) BAtaq阿塔格() feud.Barony {
	return c.阿塔格Ataq
}

func (c *贝达BaydaCounty) BBayda贝达() feud.Barony {
	return c.贝达Bayda
}

func (c *贝达BaydaCounty) BHabban哈班() feud.Barony {
	return c.哈班Habban
}

func (c *贝达BaydaCounty) BTimna提姆纳() feud.Barony {
	return c.提姆纳Timna
}

func (c *贝达BaydaCounty) BYashbum耶什布姆() feud.Barony {
	return c.耶什布姆Yashbum
}

func (c *贝达BaydaCounty) BZinjibar津吉巴尔() feud.Barony {
	return c.津吉巴尔Zinjibar
}

var CBayda贝达 BaydaCounty = &贝达BaydaCounty{}

func init() {
	f := CBayda贝达.(*贝达BaydaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "857",
		Title:     "bayda",
		TitleName: "贝达",
		TitleCode: "c_bayda",
		Baronies:  map[string]feud.Barony{},
	}

	f.海卜尔Alkhabr = BAlkhabr海卜尔
	f.海卜尔Alkhabr.SetParent(f)

	f.舒拜卡Ashshubaykah = BAshshubaykah舒拜卡
	f.舒拜卡Ashshubaykah.SetParent(f)

	f.阿塔格Ataq = BAtaq阿塔格
	f.阿塔格Ataq.SetParent(f)

	f.贝达Bayda = BBayda贝达
	f.贝达Bayda.SetParent(f)

	f.哈班Habban = BHabban哈班
	f.哈班Habban.SetParent(f)

	f.提姆纳Timna = BTimna提姆纳
	f.提姆纳Timna.SetParent(f)

	f.耶什布姆Yashbum = BYashbum耶什布姆
	f.耶什布姆Yashbum.SetParent(f)

	f.津吉巴尔Zinjibar = BZinjibar津吉巴尔
	f.津吉巴尔Zinjibar.SetParent(f)

}
