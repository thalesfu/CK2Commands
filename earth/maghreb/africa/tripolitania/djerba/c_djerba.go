package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DjerbaCounty interface {
    feud.County
    BAbarda阿巴拉达() 	feud.Barony
    BAghir艾吉尔() 	feud.Barony
    BAjim艾吉姆() 	feud.Barony
    BAljanain杰纳因() 	feud.Barony
    BCedriane瑟德里安尼() 	feud.Barony
    BElmey马伊() 	feud.Barony
    BHoumtsouk豪迈特苏克() 	feud.Barony
    BMidoun米都恩() 	feud.Barony
    BTaguermess塔古尔梅斯() 	feud.Barony
    BTataouine泰塔温() 	feud.Barony
}

type 杰尔巴DjerbaCounty struct {
	feud.BaseCounty
	阿巴拉达Abarda 	feud.Barony
	艾吉尔Aghir 	feud.Barony
	艾吉姆Ajim 	feud.Barony
	杰纳因Aljanain 	feud.Barony
	瑟德里安尼Cedriane 	feud.Barony
	马伊Elmey 	feud.Barony
	豪迈特苏克Houmtsouk 	feud.Barony
	米都恩Midoun 	feud.Barony
	塔古尔梅斯Taguermess 	feud.Barony
	泰塔温Tataouine 	feud.Barony
}

func (c *杰尔巴DjerbaCounty) BAbarda阿巴拉达() feud.Barony {
	return c.阿巴拉达Abarda
}
    
func (c *杰尔巴DjerbaCounty) BAghir艾吉尔() feud.Barony {
	return c.艾吉尔Aghir
}
    
func (c *杰尔巴DjerbaCounty) BAjim艾吉姆() feud.Barony {
	return c.艾吉姆Ajim
}
    
func (c *杰尔巴DjerbaCounty) BAljanain杰纳因() feud.Barony {
	return c.杰纳因Aljanain
}
    
func (c *杰尔巴DjerbaCounty) BCedriane瑟德里安尼() feud.Barony {
	return c.瑟德里安尼Cedriane
}
    
func (c *杰尔巴DjerbaCounty) BElmey马伊() feud.Barony {
	return c.马伊Elmey
}
    
func (c *杰尔巴DjerbaCounty) BHoumtsouk豪迈特苏克() feud.Barony {
	return c.豪迈特苏克Houmtsouk
}
    
func (c *杰尔巴DjerbaCounty) BMidoun米都恩() feud.Barony {
	return c.米都恩Midoun
}
    
func (c *杰尔巴DjerbaCounty) BTaguermess塔古尔梅斯() feud.Barony {
	return c.塔古尔梅斯Taguermess
}
    
func (c *杰尔巴DjerbaCounty) BTataouine泰塔温() feud.Barony {
	return c.泰塔温Tataouine
}
    
var CDjerba杰尔巴 DjerbaCounty = &杰尔巴DjerbaCounty{}

func init() {
	f := CDjerba杰尔巴.(*杰尔巴DjerbaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "813",
		Title:     "djerba",
		TitleName: "杰尔巴",
		TitleCode: "c_djerba",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿巴拉达Abarda = BAbarda阿巴拉达
	f.阿巴拉达Abarda.SetParent(f)
	
	f.艾吉尔Aghir = BAghir艾吉尔
	f.艾吉尔Aghir.SetParent(f)
	
	f.艾吉姆Ajim = BAjim艾吉姆
	f.艾吉姆Ajim.SetParent(f)
	
	f.杰纳因Aljanain = BAljanain杰纳因
	f.杰纳因Aljanain.SetParent(f)
	
	f.瑟德里安尼Cedriane = BCedriane瑟德里安尼
	f.瑟德里安尼Cedriane.SetParent(f)
	
	f.马伊Elmey = BElmey马伊
	f.马伊Elmey.SetParent(f)
	
	f.豪迈特苏克Houmtsouk = BHoumtsouk豪迈特苏克
	f.豪迈特苏克Houmtsouk.SetParent(f)
	
	f.米都恩Midoun = BMidoun米都恩
	f.米都恩Midoun.SetParent(f)
	
	f.塔古尔梅斯Taguermess = BTaguermess塔古尔梅斯
	f.塔古尔梅斯Taguermess.SetParent(f)
	
	f.泰塔温Tataouine = BTataouine泰塔温
	f.泰塔温Tataouine.SetParent(f)
	
}
