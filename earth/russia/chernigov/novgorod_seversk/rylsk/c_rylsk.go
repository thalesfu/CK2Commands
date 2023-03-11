package rylsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RylskCounty interface {
    feud.County
    BGlushkvo格卢什科沃() 	feud.Barony
    BIvanovskoye伊万诺夫斯科耶() 	feud.Barony
    BKapystichi卡佩斯季奇() 	feud.Barony
    BKorenovo科列诺沃() 	feud.Barony
    BKrupets克鲁佩茨() 	feud.Barony
    BRylsk雷利斯克() 	feud.Barony
    BVyr维尔() 	feud.Barony
}

type 雷利斯克RylskCounty struct {
	feud.BaseCounty
	格卢什科沃Glushkvo 	feud.Barony
	伊万诺夫斯科耶Ivanovskoye 	feud.Barony
	卡佩斯季奇Kapystichi 	feud.Barony
	科列诺沃Korenovo 	feud.Barony
	克鲁佩茨Krupets 	feud.Barony
	雷利斯克Rylsk 	feud.Barony
	维尔Vyr 	feud.Barony
}

func (c *雷利斯克RylskCounty) BGlushkvo格卢什科沃() feud.Barony {
	return c.格卢什科沃Glushkvo
}
    
func (c *雷利斯克RylskCounty) BIvanovskoye伊万诺夫斯科耶() feud.Barony {
	return c.伊万诺夫斯科耶Ivanovskoye
}
    
func (c *雷利斯克RylskCounty) BKapystichi卡佩斯季奇() feud.Barony {
	return c.卡佩斯季奇Kapystichi
}
    
func (c *雷利斯克RylskCounty) BKorenovo科列诺沃() feud.Barony {
	return c.科列诺沃Korenovo
}
    
func (c *雷利斯克RylskCounty) BKrupets克鲁佩茨() feud.Barony {
	return c.克鲁佩茨Krupets
}
    
func (c *雷利斯克RylskCounty) BRylsk雷利斯克() feud.Barony {
	return c.雷利斯克Rylsk
}
    
func (c *雷利斯克RylskCounty) BVyr维尔() feud.Barony {
	return c.维尔Vyr
}
    
var CRylsk雷利斯克 RylskCounty = &雷利斯克RylskCounty{}

func init() {
	f := CRylsk雷利斯克.(*雷利斯克RylskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1671",
		Title:     "rylsk",
		TitleName: "雷利斯克",
		TitleCode: "c_rylsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.格卢什科沃Glushkvo = BGlushkvo格卢什科沃
	f.格卢什科沃Glushkvo.SetParent(f)
	
	f.伊万诺夫斯科耶Ivanovskoye = BIvanovskoye伊万诺夫斯科耶
	f.伊万诺夫斯科耶Ivanovskoye.SetParent(f)
	
	f.卡佩斯季奇Kapystichi = BKapystichi卡佩斯季奇
	f.卡佩斯季奇Kapystichi.SetParent(f)
	
	f.科列诺沃Korenovo = BKorenovo科列诺沃
	f.科列诺沃Korenovo.SetParent(f)
	
	f.克鲁佩茨Krupets = BKrupets克鲁佩茨
	f.克鲁佩茨Krupets.SetParent(f)
	
	f.雷利斯克Rylsk = BRylsk雷利斯克
	f.雷利斯克Rylsk.SetParent(f)
	
	f.维尔Vyr = BVyr维尔
	f.维尔Vyr.SetParent(f)
	
}
