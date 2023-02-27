package soyma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SoymaCounty interface {
    feud.County
    BChosha切沙() 	feud.Barony
    BIndiga因迪加() 	feud.Barony
    BKorgovoye科尔戈沃耶() 	feud.Barony
    BSoyma索伊马() 	feud.Barony
    BTiman_ridge季曼岭() 	feud.Barony
    BVolonga沃隆加() 	feud.Barony
    BVyucheyskiy维尤切斯基() 	feud.Barony
}

type 因迪加SoymaCounty struct {
	feud.BaseCounty
	切沙Chosha 	feud.Barony
	因迪加Indiga 	feud.Barony
	科尔戈沃耶Korgovoye 	feud.Barony
	索伊马Soyma 	feud.Barony
	季曼岭Timan_ridge 	feud.Barony
	沃隆加Volonga 	feud.Barony
	维尤切斯基Vyucheyskiy 	feud.Barony
}

func (c *因迪加SoymaCounty) BChosha切沙() feud.Barony {
	return c.切沙Chosha
}
    
func (c *因迪加SoymaCounty) BIndiga因迪加() feud.Barony {
	return c.因迪加Indiga
}
    
func (c *因迪加SoymaCounty) BKorgovoye科尔戈沃耶() feud.Barony {
	return c.科尔戈沃耶Korgovoye
}
    
func (c *因迪加SoymaCounty) BSoyma索伊马() feud.Barony {
	return c.索伊马Soyma
}
    
func (c *因迪加SoymaCounty) BTiman_ridge季曼岭() feud.Barony {
	return c.季曼岭Timan_ridge
}
    
func (c *因迪加SoymaCounty) BVolonga沃隆加() feud.Barony {
	return c.沃隆加Volonga
}
    
func (c *因迪加SoymaCounty) BVyucheyskiy维尤切斯基() feud.Barony {
	return c.维尤切斯基Vyucheyskiy
}
    
var CSoyma因迪加 SoymaCounty = &因迪加SoymaCounty{}

func init() {
	f := CSoyma因迪加.(*因迪加SoymaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1830",
		Title:     "soyma",
		TitleName: "因迪加",
		TitleCode: "c_soyma",
		Baronies:  map[string]feud.Barony{},
	}

	f.切沙Chosha = BChosha切沙
	f.切沙Chosha.SetParent(f)
	
	f.因迪加Indiga = BIndiga因迪加
	f.因迪加Indiga.SetParent(f)
	
	f.科尔戈沃耶Korgovoye = BKorgovoye科尔戈沃耶
	f.科尔戈沃耶Korgovoye.SetParent(f)
	
	f.索伊马Soyma = BSoyma索伊马
	f.索伊马Soyma.SetParent(f)
	
	f.季曼岭Timan_ridge = BTiman_ridge季曼岭
	f.季曼岭Timan_ridge.SetParent(f)
	
	f.沃隆加Volonga = BVolonga沃隆加
	f.沃隆加Volonga.SetParent(f)
	
	f.维尤切斯基Vyucheyskiy = BVyucheyskiy维尤切斯基
	f.维尤切斯基Vyucheyskiy.SetParent(f)
	
}
