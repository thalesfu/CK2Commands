package azov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AzovCounty interface {
    feud.County
    BAkhtarsk阿赫塔尔斯克() 	feud.Barony
    BAzaq阿扎克() 	feud.Barony
    BAzov亚速() 	feud.Barony
    BEysk叶伊斯克() 	feud.Barony
    BKagalnik卡加利尼克() 	feud.Barony
    BKaton卡通() 	feud.Barony
    BKugey库格伊() 	feud.Barony
    BSadki萨德基() 	feud.Barony
}

type 亚速AzovCounty struct {
	feud.BaseCounty
	阿赫塔尔斯克Akhtarsk 	feud.Barony
	阿扎克Azaq 	feud.Barony
	亚速Azov 	feud.Barony
	叶伊斯克Eysk 	feud.Barony
	卡加利尼克Kagalnik 	feud.Barony
	卡通Katon 	feud.Barony
	库格伊Kugey 	feud.Barony
	萨德基Sadki 	feud.Barony
}

func (c *亚速AzovCounty) BAkhtarsk阿赫塔尔斯克() feud.Barony {
	return c.阿赫塔尔斯克Akhtarsk
}
    
func (c *亚速AzovCounty) BAzaq阿扎克() feud.Barony {
	return c.阿扎克Azaq
}
    
func (c *亚速AzovCounty) BAzov亚速() feud.Barony {
	return c.亚速Azov
}
    
func (c *亚速AzovCounty) BEysk叶伊斯克() feud.Barony {
	return c.叶伊斯克Eysk
}
    
func (c *亚速AzovCounty) BKagalnik卡加利尼克() feud.Barony {
	return c.卡加利尼克Kagalnik
}
    
func (c *亚速AzovCounty) BKaton卡通() feud.Barony {
	return c.卡通Katon
}
    
func (c *亚速AzovCounty) BKugey库格伊() feud.Barony {
	return c.库格伊Kugey
}
    
func (c *亚速AzovCounty) BSadki萨德基() feud.Barony {
	return c.萨德基Sadki
}
    
var CAzov亚速 AzovCounty = &亚速AzovCounty{}

func init() {
	f := CAzov亚速.(*亚速AzovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "597",
		Title:     "azov",
		TitleName: "亚速",
		TitleCode: "c_azov",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫塔尔斯克Akhtarsk = BAkhtarsk阿赫塔尔斯克
	f.阿赫塔尔斯克Akhtarsk.SetParent(f)
	
	f.阿扎克Azaq = BAzaq阿扎克
	f.阿扎克Azaq.SetParent(f)
	
	f.亚速Azov = BAzov亚速
	f.亚速Azov.SetParent(f)
	
	f.叶伊斯克Eysk = BEysk叶伊斯克
	f.叶伊斯克Eysk.SetParent(f)
	
	f.卡加利尼克Kagalnik = BKagalnik卡加利尼克
	f.卡加利尼克Kagalnik.SetParent(f)
	
	f.卡通Katon = BKaton卡通
	f.卡通Katon.SetParent(f)
	
	f.库格伊Kugey = BKugey库格伊
	f.库格伊Kugey.SetParent(f)
	
	f.萨德基Sadki = BSadki萨德基
	f.萨德基Sadki.SetParent(f)
	
}
