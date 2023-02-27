package tyrone

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyroneCounty interface {
    feud.County
    BAileach艾丽赫() 	feud.Barony
    BColeraine科尔雷恩() 	feud.Barony
    BDerry德里() 	feud.Barony
    BDungannon邓甘嫩() 	feud.Barony
    BDungiven邓吉文() 	feud.Barony
    BMaghera马赫拉() 	feud.Barony
    BOmagh奥马() 	feud.Barony
    BTullyhogue图里霍格() 	feud.Barony
}

type 蒂龙TyroneCounty struct {
	feud.BaseCounty
	艾丽赫Aileach 	feud.Barony
	科尔雷恩Coleraine 	feud.Barony
	德里Derry 	feud.Barony
	邓甘嫩Dungannon 	feud.Barony
	邓吉文Dungiven 	feud.Barony
	马赫拉Maghera 	feud.Barony
	奥马Omagh 	feud.Barony
	图里霍格Tullyhogue 	feud.Barony
}

func (c *蒂龙TyroneCounty) BAileach艾丽赫() feud.Barony {
	return c.艾丽赫Aileach
}
    
func (c *蒂龙TyroneCounty) BColeraine科尔雷恩() feud.Barony {
	return c.科尔雷恩Coleraine
}
    
func (c *蒂龙TyroneCounty) BDerry德里() feud.Barony {
	return c.德里Derry
}
    
func (c *蒂龙TyroneCounty) BDungannon邓甘嫩() feud.Barony {
	return c.邓甘嫩Dungannon
}
    
func (c *蒂龙TyroneCounty) BDungiven邓吉文() feud.Barony {
	return c.邓吉文Dungiven
}
    
func (c *蒂龙TyroneCounty) BMaghera马赫拉() feud.Barony {
	return c.马赫拉Maghera
}
    
func (c *蒂龙TyroneCounty) BOmagh奥马() feud.Barony {
	return c.奥马Omagh
}
    
func (c *蒂龙TyroneCounty) BTullyhogue图里霍格() feud.Barony {
	return c.图里霍格Tullyhogue
}
    
var CTyrone蒂龙 TyroneCounty = &蒂龙TyroneCounty{}

func init() {
	f := CTyrone蒂龙.(*蒂龙TyroneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "4",
		Title:     "tyrone",
		TitleName: "蒂龙",
		TitleCode: "c_tyrone",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾丽赫Aileach = BAileach艾丽赫
	f.艾丽赫Aileach.SetParent(f)
	
	f.科尔雷恩Coleraine = BColeraine科尔雷恩
	f.科尔雷恩Coleraine.SetParent(f)
	
	f.德里Derry = BDerry德里
	f.德里Derry.SetParent(f)
	
	f.邓甘嫩Dungannon = BDungannon邓甘嫩
	f.邓甘嫩Dungannon.SetParent(f)
	
	f.邓吉文Dungiven = BDungiven邓吉文
	f.邓吉文Dungiven.SetParent(f)
	
	f.马赫拉Maghera = BMaghera马赫拉
	f.马赫拉Maghera.SetParent(f)
	
	f.奥马Omagh = BOmagh奥马
	f.奥马Omagh.SetParent(f)
	
	f.图里霍格Tullyhogue = BTullyhogue图里霍格
	f.图里霍格Tullyhogue.SetParent(f)
	
}
