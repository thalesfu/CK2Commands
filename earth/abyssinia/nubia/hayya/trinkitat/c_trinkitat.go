package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrinkitatCounty interface {
    feud.County
    BDerri德日() 	feud.Barony
    BFarata法拉塔() 	feud.Barony
    BKwababa卡瓦巴巴() 	feud.Barony
    BMukden穆克登() 	feud.Barony
    BTeb特卜() 	feud.Barony
    BTokar陶卡尔() 	feud.Barony
    BTrinkitat特林基塔特() 	feud.Barony
}

type 特林基塔特TrinkitatCounty struct {
	feud.BaseCounty
	德日Derri 	feud.Barony
	法拉塔Farata 	feud.Barony
	卡瓦巴巴Kwababa 	feud.Barony
	穆克登Mukden 	feud.Barony
	特卜Teb 	feud.Barony
	陶卡尔Tokar 	feud.Barony
	特林基塔特Trinkitat 	feud.Barony
}

func (c *特林基塔特TrinkitatCounty) BDerri德日() feud.Barony {
	return c.德日Derri
}
    
func (c *特林基塔特TrinkitatCounty) BFarata法拉塔() feud.Barony {
	return c.法拉塔Farata
}
    
func (c *特林基塔特TrinkitatCounty) BKwababa卡瓦巴巴() feud.Barony {
	return c.卡瓦巴巴Kwababa
}
    
func (c *特林基塔特TrinkitatCounty) BMukden穆克登() feud.Barony {
	return c.穆克登Mukden
}
    
func (c *特林基塔特TrinkitatCounty) BTeb特卜() feud.Barony {
	return c.特卜Teb
}
    
func (c *特林基塔特TrinkitatCounty) BTokar陶卡尔() feud.Barony {
	return c.陶卡尔Tokar
}
    
func (c *特林基塔特TrinkitatCounty) BTrinkitat特林基塔特() feud.Barony {
	return c.特林基塔特Trinkitat
}
    
var CTrinkitat特林基塔特 TrinkitatCounty = &特林基塔特TrinkitatCounty{}

func init() {
	f := CTrinkitat特林基塔特.(*特林基塔特TrinkitatCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1332",
		Title:     "trinkitat",
		TitleName: "特林基塔特",
		TitleCode: "c_trinkitat",
		Baronies:  map[string]feud.Barony{},
	}

	f.德日Derri = BDerri德日
	f.德日Derri.SetParent(f)
	
	f.法拉塔Farata = BFarata法拉塔
	f.法拉塔Farata.SetParent(f)
	
	f.卡瓦巴巴Kwababa = BKwababa卡瓦巴巴
	f.卡瓦巴巴Kwababa.SetParent(f)
	
	f.穆克登Mukden = BMukden穆克登
	f.穆克登Mukden.SetParent(f)
	
	f.特卜Teb = BTeb特卜
	f.特卜Teb.SetParent(f)
	
	f.陶卡尔Tokar = BTokar陶卡尔
	f.陶卡尔Tokar.SetParent(f)
	
	f.特林基塔特Trinkitat = BTrinkitat特林基塔特
	f.特林基塔特Trinkitat.SetParent(f)
	
}
