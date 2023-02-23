package dyamare

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DyamareCounty interface {
	feud.County
	BAletu阿勒图() feud.Barony
	BAondaka阿温达卡() feud.Barony
	BDosso多索() feud.Barony
	BDyamare迪亚马尔() feud.Barony
	BSabarumowa萨巴鲁莫瓦() feud.Barony
}

type 迪亚马尔DyamareCounty struct {
	feud.BaseCounty
	阿勒图Aletu        feud.Barony
	阿温达卡Aondaka     feud.Barony
	多索Dosso         feud.Barony
	迪亚马尔Dyamare     feud.Barony
	萨巴鲁莫瓦Sabarumowa feud.Barony
}

func (c *迪亚马尔DyamareCounty) BAletu阿勒图() feud.Barony {
	return c.阿勒图Aletu
}

func (c *迪亚马尔DyamareCounty) BAondaka阿温达卡() feud.Barony {
	return c.阿温达卡Aondaka
}

func (c *迪亚马尔DyamareCounty) BDosso多索() feud.Barony {
	return c.多索Dosso
}

func (c *迪亚马尔DyamareCounty) BDyamare迪亚马尔() feud.Barony {
	return c.迪亚马尔Dyamare
}

func (c *迪亚马尔DyamareCounty) BSabarumowa萨巴鲁莫瓦() feud.Barony {
	return c.萨巴鲁莫瓦Sabarumowa
}

var CDyamare迪亚马尔 DyamareCounty = &迪亚马尔DyamareCounty{}

func init() {
	f := CDyamare迪亚马尔.(*迪亚马尔DyamareCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1753",
		Title:     "dyamare",
		TitleName: "迪亚马尔",
		TitleCode: "c_dyamare",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿勒图Aletu = BAletu阿勒图
	f.阿勒图Aletu.SetParent(f)

	f.阿温达卡Aondaka = BAondaka阿温达卡
	f.阿温达卡Aondaka.SetParent(f)

	f.多索Dosso = BDosso多索
	f.多索Dosso.SetParent(f)

	f.迪亚马尔Dyamare = BDyamare迪亚马尔
	f.迪亚马尔Dyamare.SetParent(f)

	f.萨巴鲁莫瓦Sabarumowa = BSabarumowa萨巴鲁莫瓦
	f.萨巴鲁莫瓦Sabarumowa.SetParent(f)

}
