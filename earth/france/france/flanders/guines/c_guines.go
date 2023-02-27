package guines

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GuinesCounty interface {
    feud.County
    BBourbourg布尔堡() 	feud.Barony
    BCalais加来() 	feud.Barony
    BCoulogne库洛涅() 	feud.Barony
    BDunkerque敦刻尔克() 	feud.Barony
    BGravelines格拉沃利讷() 	feud.Barony
    BGuines吉讷() 	feud.Barony
    BMarck马克() 	feud.Barony
    BSaintomer圣奥梅尔() 	feud.Barony
}

type 吉讷GuinesCounty struct {
	feud.BaseCounty
	布尔堡Bourbourg 	feud.Barony
	加来Calais 	feud.Barony
	库洛涅Coulogne 	feud.Barony
	敦刻尔克Dunkerque 	feud.Barony
	格拉沃利讷Gravelines 	feud.Barony
	吉讷Guines 	feud.Barony
	马克Marck 	feud.Barony
	圣奥梅尔Saintomer 	feud.Barony
}

func (c *吉讷GuinesCounty) BBourbourg布尔堡() feud.Barony {
	return c.布尔堡Bourbourg
}
    
func (c *吉讷GuinesCounty) BCalais加来() feud.Barony {
	return c.加来Calais
}
    
func (c *吉讷GuinesCounty) BCoulogne库洛涅() feud.Barony {
	return c.库洛涅Coulogne
}
    
func (c *吉讷GuinesCounty) BDunkerque敦刻尔克() feud.Barony {
	return c.敦刻尔克Dunkerque
}
    
func (c *吉讷GuinesCounty) BGravelines格拉沃利讷() feud.Barony {
	return c.格拉沃利讷Gravelines
}
    
func (c *吉讷GuinesCounty) BGuines吉讷() feud.Barony {
	return c.吉讷Guines
}
    
func (c *吉讷GuinesCounty) BMarck马克() feud.Barony {
	return c.马克Marck
}
    
func (c *吉讷GuinesCounty) BSaintomer圣奥梅尔() feud.Barony {
	return c.圣奥梅尔Saintomer
}
    
var CGuines吉讷 GuinesCounty = &吉讷GuinesCounty{}

func init() {
	f := CGuines吉讷.(*吉讷GuinesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "74",
		Title:     "guines",
		TitleName: "吉讷",
		TitleCode: "c_guines",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔堡Bourbourg = BBourbourg布尔堡
	f.布尔堡Bourbourg.SetParent(f)
	
	f.加来Calais = BCalais加来
	f.加来Calais.SetParent(f)
	
	f.库洛涅Coulogne = BCoulogne库洛涅
	f.库洛涅Coulogne.SetParent(f)
	
	f.敦刻尔克Dunkerque = BDunkerque敦刻尔克
	f.敦刻尔克Dunkerque.SetParent(f)
	
	f.格拉沃利讷Gravelines = BGravelines格拉沃利讷
	f.格拉沃利讷Gravelines.SetParent(f)
	
	f.吉讷Guines = BGuines吉讷
	f.吉讷Guines.SetParent(f)
	
	f.马克Marck = BMarck马克
	f.马克Marck.SetParent(f)
	
	f.圣奥梅尔Saintomer = BSaintomer圣奥梅尔
	f.圣奥梅尔Saintomer.SetParent(f)
	
}
