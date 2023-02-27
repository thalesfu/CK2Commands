package craiova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CraiovaCounty interface {
    feud.County
    BCalafat卡拉法特() 	feud.Barony
    BCaracal卡拉卡尔() 	feud.Barony
    BCorabia科拉比亚() 	feud.Barony
    BCraiova克拉约瓦() 	feud.Barony
    BHotarani霍特拉尼() 	feud.Barony
    BResita雷希察() 	feud.Barony
    BSlatina斯拉蒂纳() 	feud.Barony
}

type 克拉约瓦CraiovaCounty struct {
	feud.BaseCounty
	卡拉法特Calafat 	feud.Barony
	卡拉卡尔Caracal 	feud.Barony
	科拉比亚Corabia 	feud.Barony
	克拉约瓦Craiova 	feud.Barony
	霍特拉尼Hotarani 	feud.Barony
	雷希察Resita 	feud.Barony
	斯拉蒂纳Slatina 	feud.Barony
}

func (c *克拉约瓦CraiovaCounty) BCalafat卡拉法特() feud.Barony {
	return c.卡拉法特Calafat
}
    
func (c *克拉约瓦CraiovaCounty) BCaracal卡拉卡尔() feud.Barony {
	return c.卡拉卡尔Caracal
}
    
func (c *克拉约瓦CraiovaCounty) BCorabia科拉比亚() feud.Barony {
	return c.科拉比亚Corabia
}
    
func (c *克拉约瓦CraiovaCounty) BCraiova克拉约瓦() feud.Barony {
	return c.克拉约瓦Craiova
}
    
func (c *克拉约瓦CraiovaCounty) BHotarani霍特拉尼() feud.Barony {
	return c.霍特拉尼Hotarani
}
    
func (c *克拉约瓦CraiovaCounty) BResita雷希察() feud.Barony {
	return c.雷希察Resita
}
    
func (c *克拉约瓦CraiovaCounty) BSlatina斯拉蒂纳() feud.Barony {
	return c.斯拉蒂纳Slatina
}
    
var CCraiova克拉约瓦 CraiovaCounty = &克拉约瓦CraiovaCounty{}

func init() {
	f := CCraiova克拉约瓦.(*克拉约瓦CraiovaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1643",
		Title:     "craiova",
		TitleName: "克拉约瓦",
		TitleCode: "c_craiova",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉法特Calafat = BCalafat卡拉法特
	f.卡拉法特Calafat.SetParent(f)
	
	f.卡拉卡尔Caracal = BCaracal卡拉卡尔
	f.卡拉卡尔Caracal.SetParent(f)
	
	f.科拉比亚Corabia = BCorabia科拉比亚
	f.科拉比亚Corabia.SetParent(f)
	
	f.克拉约瓦Craiova = BCraiova克拉约瓦
	f.克拉约瓦Craiova.SetParent(f)
	
	f.霍特拉尼Hotarani = BHotarani霍特拉尼
	f.霍特拉尼Hotarani.SetParent(f)
	
	f.雷希察Resita = BResita雷希察
	f.雷希察Resita.SetParent(f)
	
	f.斯拉蒂纳Slatina = BSlatina斯拉蒂纳
	f.斯拉蒂纳Slatina.SetParent(f)
	
}
