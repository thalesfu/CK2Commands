package multan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MultanCounty interface {
	feud.County
	BAlipur阿里补罗() feud.Barony
	BKabirwala卡比瓦拉() feud.Barony
	BMulasthana茂罗三部卢() feud.Barony
	BMultan木尔坦() feud.Barony
	BPrahladpuri钵罗曷罗陀补利() feud.Barony
	BTulamba图兰巴() feud.Barony
}

type 木尔坦MultanCounty struct {
	feud.BaseCounty
	阿里补罗Alipur         feud.Barony
	卡比瓦拉Kabirwala      feud.Barony
	茂罗三部卢Mulasthana    feud.Barony
	木尔坦Multan          feud.Barony
	钵罗曷罗陀补利Prahladpuri feud.Barony
	图兰巴Tulamba         feud.Barony
}

func (c *木尔坦MultanCounty) BAlipur阿里补罗() feud.Barony {
	return c.阿里补罗Alipur
}

func (c *木尔坦MultanCounty) BKabirwala卡比瓦拉() feud.Barony {
	return c.卡比瓦拉Kabirwala
}

func (c *木尔坦MultanCounty) BMulasthana茂罗三部卢() feud.Barony {
	return c.茂罗三部卢Mulasthana
}

func (c *木尔坦MultanCounty) BMultan木尔坦() feud.Barony {
	return c.木尔坦Multan
}

func (c *木尔坦MultanCounty) BPrahladpuri钵罗曷罗陀补利() feud.Barony {
	return c.钵罗曷罗陀补利Prahladpuri
}

func (c *木尔坦MultanCounty) BTulamba图兰巴() feud.Barony {
	return c.图兰巴Tulamba
}

var CMultan木尔坦 MultanCounty = &木尔坦MultanCounty{}

func init() {
	f := CMultan木尔坦.(*木尔坦MultanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1338",
		Title:     "multan",
		TitleName: "木尔坦",
		TitleCode: "c_multan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿里补罗Alipur = BAlipur阿里补罗
	f.阿里补罗Alipur.SetParent(f)

	f.卡比瓦拉Kabirwala = BKabirwala卡比瓦拉
	f.卡比瓦拉Kabirwala.SetParent(f)

	f.茂罗三部卢Mulasthana = BMulasthana茂罗三部卢
	f.茂罗三部卢Mulasthana.SetParent(f)

	f.木尔坦Multan = BMultan木尔坦
	f.木尔坦Multan.SetParent(f)

	f.钵罗曷罗陀补利Prahladpuri = BPrahladpuri钵罗曷罗陀补利
	f.钵罗曷罗陀补利Prahladpuri.SetParent(f)

	f.图兰巴Tulamba = BTulamba图兰巴
	f.图兰巴Tulamba.SetParent(f)

}
