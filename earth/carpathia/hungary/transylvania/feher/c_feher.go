package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FeherCounty interface {
    feud.County
    BAbrudbanya奥布鲁德巴尼奥() 	feud.Barony
    BArad阿拉德() 	feud.Barony
    BElek埃莱克() 	feud.Barony
    BFeher费黑尔() 	feud.Barony
    BGyulafehervar久洛白堡() 	feud.Barony
    BNagyenyed瑙杰涅德() 	feud.Barony
    BTovis特维什() 	feud.Barony
    BVizakna维佐克瑙() 	feud.Barony
}

type 费黑尔FeherCounty struct {
	feud.BaseCounty
	奥布鲁德巴尼奥Abrudbanya 	feud.Barony
	阿拉德Arad 	feud.Barony
	埃莱克Elek 	feud.Barony
	费黑尔Feher 	feud.Barony
	久洛白堡Gyulafehervar 	feud.Barony
	瑙杰涅德Nagyenyed 	feud.Barony
	特维什Tovis 	feud.Barony
	维佐克瑙Vizakna 	feud.Barony
}

func (c *费黑尔FeherCounty) BAbrudbanya奥布鲁德巴尼奥() feud.Barony {
	return c.奥布鲁德巴尼奥Abrudbanya
}
    
func (c *费黑尔FeherCounty) BArad阿拉德() feud.Barony {
	return c.阿拉德Arad
}
    
func (c *费黑尔FeherCounty) BElek埃莱克() feud.Barony {
	return c.埃莱克Elek
}
    
func (c *费黑尔FeherCounty) BFeher费黑尔() feud.Barony {
	return c.费黑尔Feher
}
    
func (c *费黑尔FeherCounty) BGyulafehervar久洛白堡() feud.Barony {
	return c.久洛白堡Gyulafehervar
}
    
func (c *费黑尔FeherCounty) BNagyenyed瑙杰涅德() feud.Barony {
	return c.瑙杰涅德Nagyenyed
}
    
func (c *费黑尔FeherCounty) BTovis特维什() feud.Barony {
	return c.特维什Tovis
}
    
func (c *费黑尔FeherCounty) BVizakna维佐克瑙() feud.Barony {
	return c.维佐克瑙Vizakna
}
    
var CFeher费黑尔 FeherCounty = &费黑尔FeherCounty{}

func init() {
	f := CFeher费黑尔.(*费黑尔FeherCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "519",
		Title:     "feher",
		TitleName: "费黑尔",
		TitleCode: "c_feher",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥布鲁德巴尼奥Abrudbanya = BAbrudbanya奥布鲁德巴尼奥
	f.奥布鲁德巴尼奥Abrudbanya.SetParent(f)
	
	f.阿拉德Arad = BArad阿拉德
	f.阿拉德Arad.SetParent(f)
	
	f.埃莱克Elek = BElek埃莱克
	f.埃莱克Elek.SetParent(f)
	
	f.费黑尔Feher = BFeher费黑尔
	f.费黑尔Feher.SetParent(f)
	
	f.久洛白堡Gyulafehervar = BGyulafehervar久洛白堡
	f.久洛白堡Gyulafehervar.SetParent(f)
	
	f.瑙杰涅德Nagyenyed = BNagyenyed瑙杰涅德
	f.瑙杰涅德Nagyenyed.SetParent(f)
	
	f.特维什Tovis = BTovis特维什
	f.特维什Tovis.SetParent(f)
	
	f.维佐克瑙Vizakna = BVizakna维佐克瑙
	f.维佐克瑙Vizakna.SetParent(f)
	
}
