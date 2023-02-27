package tyrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyrusCounty interface {
    feud.County
    BCasalimbert伊姆伯特城堡() 	feud.Barony
    BHunin胡灵() 	feud.Barony
    BMegedel米格德尔() 	feud.Barony
    BSarafand萨拉凡德() 	feud.Barony
    BScandalon塞达洛斯() 	feud.Barony
    BSyrbelfort贝尔福() 	feud.Barony
    BSyrmontfort蒙福尔() 	feud.Barony
    BTyrus推罗() 	feud.Barony
}

type 推罗TyrusCounty struct {
	feud.BaseCounty
	伊姆伯特城堡Casalimbert 	feud.Barony
	胡灵Hunin 	feud.Barony
	米格德尔Megedel 	feud.Barony
	萨拉凡德Sarafand 	feud.Barony
	塞达洛斯Scandalon 	feud.Barony
	贝尔福Syrbelfort 	feud.Barony
	蒙福尔Syrmontfort 	feud.Barony
	推罗Tyrus 	feud.Barony
}

func (c *推罗TyrusCounty) BCasalimbert伊姆伯特城堡() feud.Barony {
	return c.伊姆伯特城堡Casalimbert
}
    
func (c *推罗TyrusCounty) BHunin胡灵() feud.Barony {
	return c.胡灵Hunin
}
    
func (c *推罗TyrusCounty) BMegedel米格德尔() feud.Barony {
	return c.米格德尔Megedel
}
    
func (c *推罗TyrusCounty) BSarafand萨拉凡德() feud.Barony {
	return c.萨拉凡德Sarafand
}
    
func (c *推罗TyrusCounty) BScandalon塞达洛斯() feud.Barony {
	return c.塞达洛斯Scandalon
}
    
func (c *推罗TyrusCounty) BSyrbelfort贝尔福() feud.Barony {
	return c.贝尔福Syrbelfort
}
    
func (c *推罗TyrusCounty) BSyrmontfort蒙福尔() feud.Barony {
	return c.蒙福尔Syrmontfort
}
    
func (c *推罗TyrusCounty) BTyrus推罗() feud.Barony {
	return c.推罗Tyrus
}
    
var CTyrus推罗 TyrusCounty = &推罗TyrusCounty{}

func init() {
	f := CTyrus推罗.(*推罗TyrusCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "771",
		Title:     "tyrus",
		TitleName: "推罗",
		TitleCode: "c_tyrus",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊姆伯特城堡Casalimbert = BCasalimbert伊姆伯特城堡
	f.伊姆伯特城堡Casalimbert.SetParent(f)
	
	f.胡灵Hunin = BHunin胡灵
	f.胡灵Hunin.SetParent(f)
	
	f.米格德尔Megedel = BMegedel米格德尔
	f.米格德尔Megedel.SetParent(f)
	
	f.萨拉凡德Sarafand = BSarafand萨拉凡德
	f.萨拉凡德Sarafand.SetParent(f)
	
	f.塞达洛斯Scandalon = BScandalon塞达洛斯
	f.塞达洛斯Scandalon.SetParent(f)
	
	f.贝尔福Syrbelfort = BSyrbelfort贝尔福
	f.贝尔福Syrbelfort.SetParent(f)
	
	f.蒙福尔Syrmontfort = BSyrmontfort蒙福尔
	f.蒙福尔Syrmontfort.SetParent(f)
	
	f.推罗Tyrus = BTyrus推罗
	f.推罗Tyrus.SetParent(f)
	
}
