package tsagaannuur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsagaannuurCounty interface {
    feud.County
    BDarvi达尔维() 	feud.Barony
    BKhaliun哈里温() 	feud.Barony
    BKhar哈尔() 	feud.Barony
    BKhovd科布多() 	feud.Barony
    BMankhan曼汗() 	feud.Barony
    BNaiman乃蛮() 	feud.Barony
    BTsagaannuur查干诺尔() 	feud.Barony
}

type 扎不罕TsagaannuurCounty struct {
	feud.BaseCounty
	达尔维Darvi 	feud.Barony
	哈里温Khaliun 	feud.Barony
	哈尔Khar 	feud.Barony
	科布多Khovd 	feud.Barony
	曼汗Mankhan 	feud.Barony
	乃蛮Naiman 	feud.Barony
	查干诺尔Tsagaannuur 	feud.Barony
}

func (c *扎不罕TsagaannuurCounty) BDarvi达尔维() feud.Barony {
	return c.达尔维Darvi
}
    
func (c *扎不罕TsagaannuurCounty) BKhaliun哈里温() feud.Barony {
	return c.哈里温Khaliun
}
    
func (c *扎不罕TsagaannuurCounty) BKhar哈尔() feud.Barony {
	return c.哈尔Khar
}
    
func (c *扎不罕TsagaannuurCounty) BKhovd科布多() feud.Barony {
	return c.科布多Khovd
}
    
func (c *扎不罕TsagaannuurCounty) BMankhan曼汗() feud.Barony {
	return c.曼汗Mankhan
}
    
func (c *扎不罕TsagaannuurCounty) BNaiman乃蛮() feud.Barony {
	return c.乃蛮Naiman
}
    
func (c *扎不罕TsagaannuurCounty) BTsagaannuur查干诺尔() feud.Barony {
	return c.查干诺尔Tsagaannuur
}
    
var CTsagaannuur扎不罕 TsagaannuurCounty = &扎不罕TsagaannuurCounty{}

func init() {
	f := CTsagaannuur扎不罕.(*扎不罕TsagaannuurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1455",
		Title:     "tsagaannuur",
		TitleName: "扎不罕",
		TitleCode: "c_tsagaannuur",
		Baronies:  map[string]feud.Barony{},
	}

	f.达尔维Darvi = BDarvi达尔维
	f.达尔维Darvi.SetParent(f)
	
	f.哈里温Khaliun = BKhaliun哈里温
	f.哈里温Khaliun.SetParent(f)
	
	f.哈尔Khar = BKhar哈尔
	f.哈尔Khar.SetParent(f)
	
	f.科布多Khovd = BKhovd科布多
	f.科布多Khovd.SetParent(f)
	
	f.曼汗Mankhan = BMankhan曼汗
	f.曼汗Mankhan.SetParent(f)
	
	f.乃蛮Naiman = BNaiman乃蛮
	f.乃蛮Naiman.SetParent(f)
	
	f.查干诺尔Tsagaannuur = BTsagaannuur查干诺尔
	f.查干诺尔Tsagaannuur.SetParent(f)
	
}
