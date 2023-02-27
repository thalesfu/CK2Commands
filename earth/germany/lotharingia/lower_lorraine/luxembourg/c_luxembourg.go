package luxembourg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LuxembourgCounty interface {
    feud.County
    BArlon阿尔隆() 	feud.Barony
    BBouillon布永() 	feud.Barony
    BDifferdange迪弗当日() 	feud.Barony
    BEttelbruck埃特尔布吕克() 	feud.Barony
    BLongwy隆维() 	feud.Barony
    BLuxembourg卢森堡() 	feud.Barony
    BNeufchateau讷沙托() 	feud.Barony
    BSaintvith圣维特() 	feud.Barony
}

type 卢森堡LuxembourgCounty struct {
	feud.BaseCounty
	阿尔隆Arlon 	feud.Barony
	布永Bouillon 	feud.Barony
	迪弗当日Differdange 	feud.Barony
	埃特尔布吕克Ettelbruck 	feud.Barony
	隆维Longwy 	feud.Barony
	卢森堡Luxembourg 	feud.Barony
	讷沙托Neufchateau 	feud.Barony
	圣维特Saintvith 	feud.Barony
}

func (c *卢森堡LuxembourgCounty) BArlon阿尔隆() feud.Barony {
	return c.阿尔隆Arlon
}
    
func (c *卢森堡LuxembourgCounty) BBouillon布永() feud.Barony {
	return c.布永Bouillon
}
    
func (c *卢森堡LuxembourgCounty) BDifferdange迪弗当日() feud.Barony {
	return c.迪弗当日Differdange
}
    
func (c *卢森堡LuxembourgCounty) BEttelbruck埃特尔布吕克() feud.Barony {
	return c.埃特尔布吕克Ettelbruck
}
    
func (c *卢森堡LuxembourgCounty) BLongwy隆维() feud.Barony {
	return c.隆维Longwy
}
    
func (c *卢森堡LuxembourgCounty) BLuxembourg卢森堡() feud.Barony {
	return c.卢森堡Luxembourg
}
    
func (c *卢森堡LuxembourgCounty) BNeufchateau讷沙托() feud.Barony {
	return c.讷沙托Neufchateau
}
    
func (c *卢森堡LuxembourgCounty) BSaintvith圣维特() feud.Barony {
	return c.圣维特Saintvith
}
    
var CLuxembourg卢森堡 LuxembourgCounty = &卢森堡LuxembourgCounty{}

func init() {
	f := CLuxembourg卢森堡.(*卢森堡LuxembourgCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "115",
		Title:     "luxembourg",
		TitleName: "卢森堡",
		TitleCode: "c_luxembourg",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔隆Arlon = BArlon阿尔隆
	f.阿尔隆Arlon.SetParent(f)
	
	f.布永Bouillon = BBouillon布永
	f.布永Bouillon.SetParent(f)
	
	f.迪弗当日Differdange = BDifferdange迪弗当日
	f.迪弗当日Differdange.SetParent(f)
	
	f.埃特尔布吕克Ettelbruck = BEttelbruck埃特尔布吕克
	f.埃特尔布吕克Ettelbruck.SetParent(f)
	
	f.隆维Longwy = BLongwy隆维
	f.隆维Longwy.SetParent(f)
	
	f.卢森堡Luxembourg = BLuxembourg卢森堡
	f.卢森堡Luxembourg.SetParent(f)
	
	f.讷沙托Neufchateau = BNeufchateau讷沙托
	f.讷沙托Neufchateau.SetParent(f)
	
	f.圣维特Saintvith = BSaintvith圣维特
	f.圣维特Saintvith.SetParent(f)
	
}
