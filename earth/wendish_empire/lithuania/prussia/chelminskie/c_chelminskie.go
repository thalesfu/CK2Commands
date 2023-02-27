package chelminskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChelminskieCounty interface {
    feud.County
    BBriesen布里森() 	feud.Barony
    BChelmno海乌姆诺() 	feud.Barony
    BEylau艾劳() 	feud.Barony
    BKulm沃扎() 	feud.Barony
    BLobau勒包() 	feud.Barony
    BNiedenburg尼登贝格() 	feud.Barony
    BRehden雷登() 	feud.Barony
    BThorn托恩() 	feud.Barony
}

type 海乌姆诺ChelminskieCounty struct {
	feud.BaseCounty
	布里森Briesen 	feud.Barony
	海乌姆诺Chelmno 	feud.Barony
	艾劳Eylau 	feud.Barony
	沃扎Kulm 	feud.Barony
	勒包Lobau 	feud.Barony
	尼登贝格Niedenburg 	feud.Barony
	雷登Rehden 	feud.Barony
	托恩Thorn 	feud.Barony
}

func (c *海乌姆诺ChelminskieCounty) BBriesen布里森() feud.Barony {
	return c.布里森Briesen
}
    
func (c *海乌姆诺ChelminskieCounty) BChelmno海乌姆诺() feud.Barony {
	return c.海乌姆诺Chelmno
}
    
func (c *海乌姆诺ChelminskieCounty) BEylau艾劳() feud.Barony {
	return c.艾劳Eylau
}
    
func (c *海乌姆诺ChelminskieCounty) BKulm沃扎() feud.Barony {
	return c.沃扎Kulm
}
    
func (c *海乌姆诺ChelminskieCounty) BLobau勒包() feud.Barony {
	return c.勒包Lobau
}
    
func (c *海乌姆诺ChelminskieCounty) BNiedenburg尼登贝格() feud.Barony {
	return c.尼登贝格Niedenburg
}
    
func (c *海乌姆诺ChelminskieCounty) BRehden雷登() feud.Barony {
	return c.雷登Rehden
}
    
func (c *海乌姆诺ChelminskieCounty) BThorn托恩() feud.Barony {
	return c.托恩Thorn
}
    
var CChelminskie海乌姆诺 ChelminskieCounty = &海乌姆诺ChelminskieCounty{}

func init() {
	f := CChelminskie海乌姆诺.(*海乌姆诺ChelminskieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "369",
		Title:     "chelminskie",
		TitleName: "海乌姆诺",
		TitleCode: "c_chelminskie",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里森Briesen = BBriesen布里森
	f.布里森Briesen.SetParent(f)
	
	f.海乌姆诺Chelmno = BChelmno海乌姆诺
	f.海乌姆诺Chelmno.SetParent(f)
	
	f.艾劳Eylau = BEylau艾劳
	f.艾劳Eylau.SetParent(f)
	
	f.沃扎Kulm = BKulm沃扎
	f.沃扎Kulm.SetParent(f)
	
	f.勒包Lobau = BLobau勒包
	f.勒包Lobau.SetParent(f)
	
	f.尼登贝格Niedenburg = BNiedenburg尼登贝格
	f.尼登贝格Niedenburg.SetParent(f)
	
	f.雷登Rehden = BRehden雷登
	f.雷登Rehden.SetParent(f)
	
	f.托恩Thorn = BThorn托恩
	f.托恩Thorn.SetParent(f)
	
}
