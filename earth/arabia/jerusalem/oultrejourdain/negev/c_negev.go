package negev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NegevCounty interface {
    feud.County
    BAlbaqar加扎勒() 	feud.Barony
    BAvdat阿伏达特() 	feud.Barony
    BDimona迪莫纳() 	feud.Barony
    BEzuz埃祖兹() 	feud.Barony
    BHaluza哈鲁扎() 	feud.Barony
    BKmehin梅因() 	feud.Barony
    BNegev内盖夫() 	feud.Barony
    BYeruham耶鲁汉姆() 	feud.Barony
}

type 内盖夫NegevCounty struct {
	feud.BaseCounty
	加扎勒Albaqar 	feud.Barony
	阿伏达特Avdat 	feud.Barony
	迪莫纳Dimona 	feud.Barony
	埃祖兹Ezuz 	feud.Barony
	哈鲁扎Haluza 	feud.Barony
	梅因Kmehin 	feud.Barony
	内盖夫Negev 	feud.Barony
	耶鲁汉姆Yeruham 	feud.Barony
}

func (c *内盖夫NegevCounty) BAlbaqar加扎勒() feud.Barony {
	return c.加扎勒Albaqar
}
    
func (c *内盖夫NegevCounty) BAvdat阿伏达特() feud.Barony {
	return c.阿伏达特Avdat
}
    
func (c *内盖夫NegevCounty) BDimona迪莫纳() feud.Barony {
	return c.迪莫纳Dimona
}
    
func (c *内盖夫NegevCounty) BEzuz埃祖兹() feud.Barony {
	return c.埃祖兹Ezuz
}
    
func (c *内盖夫NegevCounty) BHaluza哈鲁扎() feud.Barony {
	return c.哈鲁扎Haluza
}
    
func (c *内盖夫NegevCounty) BKmehin梅因() feud.Barony {
	return c.梅因Kmehin
}
    
func (c *内盖夫NegevCounty) BNegev内盖夫() feud.Barony {
	return c.内盖夫Negev
}
    
func (c *内盖夫NegevCounty) BYeruham耶鲁汉姆() feud.Barony {
	return c.耶鲁汉姆Yeruham
}
    
var CNegev内盖夫 NegevCounty = &内盖夫NegevCounty{}

func init() {
	f := CNegev内盖夫.(*内盖夫NegevCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "782",
		Title:     "negev",
		TitleName: "内盖夫",
		TitleCode: "c_negev",
		Baronies:  map[string]feud.Barony{},
	}

	f.加扎勒Albaqar = BAlbaqar加扎勒
	f.加扎勒Albaqar.SetParent(f)
	
	f.阿伏达特Avdat = BAvdat阿伏达特
	f.阿伏达特Avdat.SetParent(f)
	
	f.迪莫纳Dimona = BDimona迪莫纳
	f.迪莫纳Dimona.SetParent(f)
	
	f.埃祖兹Ezuz = BEzuz埃祖兹
	f.埃祖兹Ezuz.SetParent(f)
	
	f.哈鲁扎Haluza = BHaluza哈鲁扎
	f.哈鲁扎Haluza.SetParent(f)
	
	f.梅因Kmehin = BKmehin梅因
	f.梅因Kmehin.SetParent(f)
	
	f.内盖夫Negev = BNegev内盖夫
	f.内盖夫Negev.SetParent(f)
	
	f.耶鲁汉姆Yeruham = BYeruham耶鲁汉姆
	f.耶鲁汉姆Yeruham.SetParent(f)
	
}
