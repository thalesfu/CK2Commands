package rouergue

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RouergueCounty interface {
    feud.County
    BCaylus凯吕斯() 	feud.Barony
    BEstaing埃斯坦() 	feud.Barony
    BMillau米约() 	feud.Barony
    BNajac纳雅克() 	feud.Barony
    BRodez罗德兹() 	feud.Barony
    BStaffrique圣阿夫里克() 	feud.Barony
    BVabres瓦布尔() 	feud.Barony
    BVillefranche自由城() 	feud.Barony
}

type 鲁埃格RouergueCounty struct {
	feud.BaseCounty
	凯吕斯Caylus 	feud.Barony
	埃斯坦Estaing 	feud.Barony
	米约Millau 	feud.Barony
	纳雅克Najac 	feud.Barony
	罗德兹Rodez 	feud.Barony
	圣阿夫里克Staffrique 	feud.Barony
	瓦布尔Vabres 	feud.Barony
	自由城Villefranche 	feud.Barony
}

func (c *鲁埃格RouergueCounty) BCaylus凯吕斯() feud.Barony {
	return c.凯吕斯Caylus
}
    
func (c *鲁埃格RouergueCounty) BEstaing埃斯坦() feud.Barony {
	return c.埃斯坦Estaing
}
    
func (c *鲁埃格RouergueCounty) BMillau米约() feud.Barony {
	return c.米约Millau
}
    
func (c *鲁埃格RouergueCounty) BNajac纳雅克() feud.Barony {
	return c.纳雅克Najac
}
    
func (c *鲁埃格RouergueCounty) BRodez罗德兹() feud.Barony {
	return c.罗德兹Rodez
}
    
func (c *鲁埃格RouergueCounty) BStaffrique圣阿夫里克() feud.Barony {
	return c.圣阿夫里克Staffrique
}
    
func (c *鲁埃格RouergueCounty) BVabres瓦布尔() feud.Barony {
	return c.瓦布尔Vabres
}
    
func (c *鲁埃格RouergueCounty) BVillefranche自由城() feud.Barony {
	return c.自由城Villefranche
}
    
var CRouergue鲁埃格 RouergueCounty = &鲁埃格RouergueCounty{}

func init() {
	f := CRouergue鲁埃格.(*鲁埃格RouergueCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "218",
		Title:     "rouergue",
		TitleName: "鲁埃格",
		TitleCode: "c_rouergue",
		Baronies:  map[string]feud.Barony{},
	}

	f.凯吕斯Caylus = BCaylus凯吕斯
	f.凯吕斯Caylus.SetParent(f)
	
	f.埃斯坦Estaing = BEstaing埃斯坦
	f.埃斯坦Estaing.SetParent(f)
	
	f.米约Millau = BMillau米约
	f.米约Millau.SetParent(f)
	
	f.纳雅克Najac = BNajac纳雅克
	f.纳雅克Najac.SetParent(f)
	
	f.罗德兹Rodez = BRodez罗德兹
	f.罗德兹Rodez.SetParent(f)
	
	f.圣阿夫里克Staffrique = BStaffrique圣阿夫里克
	f.圣阿夫里克Staffrique.SetParent(f)
	
	f.瓦布尔Vabres = BVabres瓦布尔
	f.瓦布尔Vabres.SetParent(f)
	
	f.自由城Villefranche = BVillefranche自由城
	f.自由城Villefranche.SetParent(f)
	
}
