package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GrazCounty interface {
    feud.County
    BFeldbach费尔德巴赫() 	feud.Barony
    BFriedberg_graz弗里德贝格() 	feud.Barony
    BGraz格拉茨() 	feud.Barony
    BLeibnitz莱布尼茨() 	feud.Barony
    BRadkersburg拉德克斯堡() 	feud.Barony
    BReun罗因() 	feud.Barony
    BStainz施泰因茨() 	feud.Barony
    BVoitsberg福伊茨贝格() 	feud.Barony
}

type 格拉茨GrazCounty struct {
	feud.BaseCounty
	费尔德巴赫Feldbach 	feud.Barony
	弗里德贝格Friedberg_graz 	feud.Barony
	格拉茨Graz 	feud.Barony
	莱布尼茨Leibnitz 	feud.Barony
	拉德克斯堡Radkersburg 	feud.Barony
	罗因Reun 	feud.Barony
	施泰因茨Stainz 	feud.Barony
	福伊茨贝格Voitsberg 	feud.Barony
}

func (c *格拉茨GrazCounty) BFeldbach费尔德巴赫() feud.Barony {
	return c.费尔德巴赫Feldbach
}
    
func (c *格拉茨GrazCounty) BFriedberg_graz弗里德贝格() feud.Barony {
	return c.弗里德贝格Friedberg_graz
}
    
func (c *格拉茨GrazCounty) BGraz格拉茨() feud.Barony {
	return c.格拉茨Graz
}
    
func (c *格拉茨GrazCounty) BLeibnitz莱布尼茨() feud.Barony {
	return c.莱布尼茨Leibnitz
}
    
func (c *格拉茨GrazCounty) BRadkersburg拉德克斯堡() feud.Barony {
	return c.拉德克斯堡Radkersburg
}
    
func (c *格拉茨GrazCounty) BReun罗因() feud.Barony {
	return c.罗因Reun
}
    
func (c *格拉茨GrazCounty) BStainz施泰因茨() feud.Barony {
	return c.施泰因茨Stainz
}
    
func (c *格拉茨GrazCounty) BVoitsberg福伊茨贝格() feud.Barony {
	return c.福伊茨贝格Voitsberg
}
    
var CGraz格拉茨 GrazCounty = &格拉茨GrazCounty{}

func init() {
	f := CGraz格拉茨.(*格拉茨GrazCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1693",
		Title:     "graz",
		TitleName: "格拉茨",
		TitleCode: "c_graz",
		Baronies:  map[string]feud.Barony{},
	}

	f.费尔德巴赫Feldbach = BFeldbach费尔德巴赫
	f.费尔德巴赫Feldbach.SetParent(f)
	
	f.弗里德贝格Friedberg_graz = BFriedberg_graz弗里德贝格
	f.弗里德贝格Friedberg_graz.SetParent(f)
	
	f.格拉茨Graz = BGraz格拉茨
	f.格拉茨Graz.SetParent(f)
	
	f.莱布尼茨Leibnitz = BLeibnitz莱布尼茨
	f.莱布尼茨Leibnitz.SetParent(f)
	
	f.拉德克斯堡Radkersburg = BRadkersburg拉德克斯堡
	f.拉德克斯堡Radkersburg.SetParent(f)
	
	f.罗因Reun = BReun罗因
	f.罗因Reun.SetParent(f)
	
	f.施泰因茨Stainz = BStainz施泰因茨
	f.施泰因茨Stainz.SetParent(f)
	
	f.福伊茨贝格Voitsberg = BVoitsberg福伊茨贝格
	f.福伊茨贝格Voitsberg.SetParent(f)
	
}
