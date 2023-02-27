package kusbulak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KusbulakCounty interface {
    feud.County
    BAidabul艾达布尔() 	feud.Barony
    BGusken古斯肯() 	feud.Barony
    BJasliq扎斯雷克() 	feud.Barony
    BKusbulak库斯布拉克() 	feud.Barony
    BKyushe久舍() 	feud.Barony
    BKyzyl_bulak克孜勒布拉克() 	feud.Barony
    BSellizure塞利祖尔() 	feud.Barony
}

type 库斯布拉克KusbulakCounty struct {
	feud.BaseCounty
	艾达布尔Aidabul 	feud.Barony
	古斯肯Gusken 	feud.Barony
	扎斯雷克Jasliq 	feud.Barony
	库斯布拉克Kusbulak 	feud.Barony
	久舍Kyushe 	feud.Barony
	克孜勒布拉克Kyzyl_bulak 	feud.Barony
	塞利祖尔Sellizure 	feud.Barony
}

func (c *库斯布拉克KusbulakCounty) BAidabul艾达布尔() feud.Barony {
	return c.艾达布尔Aidabul
}
    
func (c *库斯布拉克KusbulakCounty) BGusken古斯肯() feud.Barony {
	return c.古斯肯Gusken
}
    
func (c *库斯布拉克KusbulakCounty) BJasliq扎斯雷克() feud.Barony {
	return c.扎斯雷克Jasliq
}
    
func (c *库斯布拉克KusbulakCounty) BKusbulak库斯布拉克() feud.Barony {
	return c.库斯布拉克Kusbulak
}
    
func (c *库斯布拉克KusbulakCounty) BKyushe久舍() feud.Barony {
	return c.久舍Kyushe
}
    
func (c *库斯布拉克KusbulakCounty) BKyzyl_bulak克孜勒布拉克() feud.Barony {
	return c.克孜勒布拉克Kyzyl_bulak
}
    
func (c *库斯布拉克KusbulakCounty) BSellizure塞利祖尔() feud.Barony {
	return c.塞利祖尔Sellizure
}
    
var CKusbulak库斯布拉克 KusbulakCounty = &库斯布拉克KusbulakCounty{}

func init() {
	f := CKusbulak库斯布拉克.(*库斯布拉克KusbulakCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1792",
		Title:     "kusbulak",
		TitleName: "库斯布拉克",
		TitleCode: "c_kusbulak",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾达布尔Aidabul = BAidabul艾达布尔
	f.艾达布尔Aidabul.SetParent(f)
	
	f.古斯肯Gusken = BGusken古斯肯
	f.古斯肯Gusken.SetParent(f)
	
	f.扎斯雷克Jasliq = BJasliq扎斯雷克
	f.扎斯雷克Jasliq.SetParent(f)
	
	f.库斯布拉克Kusbulak = BKusbulak库斯布拉克
	f.库斯布拉克Kusbulak.SetParent(f)
	
	f.久舍Kyushe = BKyushe久舍
	f.久舍Kyushe.SetParent(f)
	
	f.克孜勒布拉克Kyzyl_bulak = BKyzyl_bulak克孜勒布拉克
	f.克孜勒布拉克Kyzyl_bulak.SetParent(f)
	
	f.塞利祖尔Sellizure = BSellizure塞利祖尔
	f.塞利祖尔Sellizure.SetParent(f)
	
}
