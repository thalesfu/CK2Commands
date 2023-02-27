package ejin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EjinCounty interface {
    feud.County
    BHwantsaopa黄草坝() 	feud.Barony
    BJijiagou纪家沟() 	feud.Barony
    BKhara_khoto黑水城() 	feud.Barony
    BMazongshan马鬃山() 	feud.Barony
    BTongchongcha桐冲岔() 	feud.Barony
    BYuhe玉河() 	feud.Barony
    BZhugentan竹根滩() 	feud.Barony
}

type 亦集乃EjinCounty struct {
	feud.BaseCounty
	黄草坝Hwantsaopa 	feud.Barony
	纪家沟Jijiagou 	feud.Barony
	黑水城Khara_khoto 	feud.Barony
	马鬃山Mazongshan 	feud.Barony
	桐冲岔Tongchongcha 	feud.Barony
	玉河Yuhe 	feud.Barony
	竹根滩Zhugentan 	feud.Barony
}

func (c *亦集乃EjinCounty) BHwantsaopa黄草坝() feud.Barony {
	return c.黄草坝Hwantsaopa
}
    
func (c *亦集乃EjinCounty) BJijiagou纪家沟() feud.Barony {
	return c.纪家沟Jijiagou
}
    
func (c *亦集乃EjinCounty) BKhara_khoto黑水城() feud.Barony {
	return c.黑水城Khara_khoto
}
    
func (c *亦集乃EjinCounty) BMazongshan马鬃山() feud.Barony {
	return c.马鬃山Mazongshan
}
    
func (c *亦集乃EjinCounty) BTongchongcha桐冲岔() feud.Barony {
	return c.桐冲岔Tongchongcha
}
    
func (c *亦集乃EjinCounty) BYuhe玉河() feud.Barony {
	return c.玉河Yuhe
}
    
func (c *亦集乃EjinCounty) BZhugentan竹根滩() feud.Barony {
	return c.竹根滩Zhugentan
}
    
var CEjin亦集乃 EjinCounty = &亦集乃EjinCounty{}

func init() {
	f := CEjin亦集乃.(*亦集乃EjinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1514",
		Title:     "ejin",
		TitleName: "亦集乃",
		TitleCode: "c_ejin",
		Baronies:  map[string]feud.Barony{},
	}

	f.黄草坝Hwantsaopa = BHwantsaopa黄草坝
	f.黄草坝Hwantsaopa.SetParent(f)
	
	f.纪家沟Jijiagou = BJijiagou纪家沟
	f.纪家沟Jijiagou.SetParent(f)
	
	f.黑水城Khara_khoto = BKhara_khoto黑水城
	f.黑水城Khara_khoto.SetParent(f)
	
	f.马鬃山Mazongshan = BMazongshan马鬃山
	f.马鬃山Mazongshan.SetParent(f)
	
	f.桐冲岔Tongchongcha = BTongchongcha桐冲岔
	f.桐冲岔Tongchongcha.SetParent(f)
	
	f.玉河Yuhe = BYuhe玉河
	f.玉河Yuhe.SetParent(f)
	
	f.竹根滩Zhugentan = BZhugentan竹根滩
	f.竹根滩Zhugentan.SetParent(f)
	
}
