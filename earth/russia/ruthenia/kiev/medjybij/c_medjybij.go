package medjybij

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedjybijCounty interface {
    feud.County
    BBerdychiv别尔基切夫() 	feud.Barony
    BChudniv丘德尼夫() 	feud.Barony
    BHubyn古宾() 	feud.Barony
    BKhmilnyk赫梅利尼克() 	feud.Barony
    BLetychiv列季奇夫() 	feud.Barony
    BLiubar柳巴尔() 	feud.Barony
    BMedjybij梅德日比日() 	feud.Barony
}

type 梅德日比日MedjybijCounty struct {
	feud.BaseCounty
	别尔基切夫Berdychiv 	feud.Barony
	丘德尼夫Chudniv 	feud.Barony
	古宾Hubyn 	feud.Barony
	赫梅利尼克Khmilnyk 	feud.Barony
	列季奇夫Letychiv 	feud.Barony
	柳巴尔Liubar 	feud.Barony
	梅德日比日Medjybij 	feud.Barony
}

func (c *梅德日比日MedjybijCounty) BBerdychiv别尔基切夫() feud.Barony {
	return c.别尔基切夫Berdychiv
}
    
func (c *梅德日比日MedjybijCounty) BChudniv丘德尼夫() feud.Barony {
	return c.丘德尼夫Chudniv
}
    
func (c *梅德日比日MedjybijCounty) BHubyn古宾() feud.Barony {
	return c.古宾Hubyn
}
    
func (c *梅德日比日MedjybijCounty) BKhmilnyk赫梅利尼克() feud.Barony {
	return c.赫梅利尼克Khmilnyk
}
    
func (c *梅德日比日MedjybijCounty) BLetychiv列季奇夫() feud.Barony {
	return c.列季奇夫Letychiv
}
    
func (c *梅德日比日MedjybijCounty) BLiubar柳巴尔() feud.Barony {
	return c.柳巴尔Liubar
}
    
func (c *梅德日比日MedjybijCounty) BMedjybij梅德日比日() feud.Barony {
	return c.梅德日比日Medjybij
}
    
var CMedjybij梅德日比日 MedjybijCounty = &梅德日比日MedjybijCounty{}

func init() {
	f := CMedjybij梅德日比日.(*梅德日比日MedjybijCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1648",
		Title:     "medjybij",
		TitleName: "梅德日比日",
		TitleCode: "c_medjybij",
		Baronies:  map[string]feud.Barony{},
	}

	f.别尔基切夫Berdychiv = BBerdychiv别尔基切夫
	f.别尔基切夫Berdychiv.SetParent(f)
	
	f.丘德尼夫Chudniv = BChudniv丘德尼夫
	f.丘德尼夫Chudniv.SetParent(f)
	
	f.古宾Hubyn = BHubyn古宾
	f.古宾Hubyn.SetParent(f)
	
	f.赫梅利尼克Khmilnyk = BKhmilnyk赫梅利尼克
	f.赫梅利尼克Khmilnyk.SetParent(f)
	
	f.列季奇夫Letychiv = BLetychiv列季奇夫
	f.列季奇夫Letychiv.SetParent(f)
	
	f.柳巴尔Liubar = BLiubar柳巴尔
	f.柳巴尔Liubar.SetParent(f)
	
	f.梅德日比日Medjybij = BMedjybij梅德日比日
	f.梅德日比日Medjybij.SetParent(f)
	
}
