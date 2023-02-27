package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CagliariCounty interface {
    feud.County
    BAssemini阿塞米尼() 	feud.Barony
    BCagliari卡利亚里() 	feud.Barony
    BCapoterra卡波泰拉() 	feud.Barony
    BCarbonia卡尔博尼亚() 	feud.Barony
    BDolianova多利亚诺瓦() 	feud.Barony
    BIglesias伊格莱西亚斯() 	feud.Barony
    BSantaigia圣伊吉亚() 	feud.Barony
}

type 卡利亚里CagliariCounty struct {
	feud.BaseCounty
	阿塞米尼Assemini 	feud.Barony
	卡利亚里Cagliari 	feud.Barony
	卡波泰拉Capoterra 	feud.Barony
	卡尔博尼亚Carbonia 	feud.Barony
	多利亚诺瓦Dolianova 	feud.Barony
	伊格莱西亚斯Iglesias 	feud.Barony
	圣伊吉亚Santaigia 	feud.Barony
}

func (c *卡利亚里CagliariCounty) BAssemini阿塞米尼() feud.Barony {
	return c.阿塞米尼Assemini
}
    
func (c *卡利亚里CagliariCounty) BCagliari卡利亚里() feud.Barony {
	return c.卡利亚里Cagliari
}
    
func (c *卡利亚里CagliariCounty) BCapoterra卡波泰拉() feud.Barony {
	return c.卡波泰拉Capoterra
}
    
func (c *卡利亚里CagliariCounty) BCarbonia卡尔博尼亚() feud.Barony {
	return c.卡尔博尼亚Carbonia
}
    
func (c *卡利亚里CagliariCounty) BDolianova多利亚诺瓦() feud.Barony {
	return c.多利亚诺瓦Dolianova
}
    
func (c *卡利亚里CagliariCounty) BIglesias伊格莱西亚斯() feud.Barony {
	return c.伊格莱西亚斯Iglesias
}
    
func (c *卡利亚里CagliariCounty) BSantaigia圣伊吉亚() feud.Barony {
	return c.圣伊吉亚Santaigia
}
    
var CCagliari卡利亚里 CagliariCounty = &卡利亚里CagliariCounty{}

func init() {
	f := CCagliari卡利亚里.(*卡利亚里CagliariCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "326",
		Title:     "cagliari",
		TitleName: "卡利亚里",
		TitleCode: "c_cagliari",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿塞米尼Assemini = BAssemini阿塞米尼
	f.阿塞米尼Assemini.SetParent(f)
	
	f.卡利亚里Cagliari = BCagliari卡利亚里
	f.卡利亚里Cagliari.SetParent(f)
	
	f.卡波泰拉Capoterra = BCapoterra卡波泰拉
	f.卡波泰拉Capoterra.SetParent(f)
	
	f.卡尔博尼亚Carbonia = BCarbonia卡尔博尼亚
	f.卡尔博尼亚Carbonia.SetParent(f)
	
	f.多利亚诺瓦Dolianova = BDolianova多利亚诺瓦
	f.多利亚诺瓦Dolianova.SetParent(f)
	
	f.伊格莱西亚斯Iglesias = BIglesias伊格莱西亚斯
	f.伊格莱西亚斯Iglesias.SetParent(f)
	
	f.圣伊吉亚Santaigia = BSantaigia圣伊吉亚
	f.圣伊吉亚Santaigia.SetParent(f)
	
}
