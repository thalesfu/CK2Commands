package al_aqabah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_aqabahCounty interface {
    feud.County
    BAqabah亚喀巴() 	feud.Barony
    BElifaz艾因法兹() 	feud.Barony
    BFeifa费法() 	feud.Barony
    BMazraa迈兹拉阿() 	feud.Barony
    BQuwairah古韦拉() 	feud.Barony
    BReeshah雷瑟哈() 	feud.Barony
    BSamar萨马尔() 	feud.Barony
    BYotvata约特瓦塔() 	feud.Barony
}

type 亚喀巴Al_aqabahCounty struct {
	feud.BaseCounty
	亚喀巴Aqabah 	feud.Barony
	艾因法兹Elifaz 	feud.Barony
	费法Feifa 	feud.Barony
	迈兹拉阿Mazraa 	feud.Barony
	古韦拉Quwairah 	feud.Barony
	雷瑟哈Reeshah 	feud.Barony
	萨马尔Samar 	feud.Barony
	约特瓦塔Yotvata 	feud.Barony
}

func (c *亚喀巴Al_aqabahCounty) BAqabah亚喀巴() feud.Barony {
	return c.亚喀巴Aqabah
}
    
func (c *亚喀巴Al_aqabahCounty) BElifaz艾因法兹() feud.Barony {
	return c.艾因法兹Elifaz
}
    
func (c *亚喀巴Al_aqabahCounty) BFeifa费法() feud.Barony {
	return c.费法Feifa
}
    
func (c *亚喀巴Al_aqabahCounty) BMazraa迈兹拉阿() feud.Barony {
	return c.迈兹拉阿Mazraa
}
    
func (c *亚喀巴Al_aqabahCounty) BQuwairah古韦拉() feud.Barony {
	return c.古韦拉Quwairah
}
    
func (c *亚喀巴Al_aqabahCounty) BReeshah雷瑟哈() feud.Barony {
	return c.雷瑟哈Reeshah
}
    
func (c *亚喀巴Al_aqabahCounty) BSamar萨马尔() feud.Barony {
	return c.萨马尔Samar
}
    
func (c *亚喀巴Al_aqabahCounty) BYotvata约特瓦塔() feud.Barony {
	return c.约特瓦塔Yotvata
}
    
var CAl_aqabah亚喀巴 Al_aqabahCounty = &亚喀巴Al_aqabahCounty{}

func init() {
	f := CAl_aqabah亚喀巴.(*亚喀巴Al_aqabahCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "784",
		Title:     "al_aqabah",
		TitleName: "亚喀巴",
		TitleCode: "c_al_aqabah",
		Baronies:  map[string]feud.Barony{},
	}

	f.亚喀巴Aqabah = BAqabah亚喀巴
	f.亚喀巴Aqabah.SetParent(f)
	
	f.艾因法兹Elifaz = BElifaz艾因法兹
	f.艾因法兹Elifaz.SetParent(f)
	
	f.费法Feifa = BFeifa费法
	f.费法Feifa.SetParent(f)
	
	f.迈兹拉阿Mazraa = BMazraa迈兹拉阿
	f.迈兹拉阿Mazraa.SetParent(f)
	
	f.古韦拉Quwairah = BQuwairah古韦拉
	f.古韦拉Quwairah.SetParent(f)
	
	f.雷瑟哈Reeshah = BReeshah雷瑟哈
	f.雷瑟哈Reeshah.SetParent(f)
	
	f.萨马尔Samar = BSamar萨马尔
	f.萨马尔Samar.SetParent(f)
	
	f.约特瓦塔Yotvata = BYotvata约特瓦塔
	f.约特瓦塔Yotvata.SetParent(f)
	
}
