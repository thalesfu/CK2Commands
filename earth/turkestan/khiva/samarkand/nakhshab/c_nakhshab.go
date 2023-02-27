package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NakhshabCounty interface {
    feud.County
    BBozariq博扎里克() 	feud.Barony
    BFirabr菲拉布() 	feud.Barony
    BMirzomomin米尔佐莫明() 	feud.Barony
    BNakhsab那色波() 	feud.Barony
    BPaykand拜坎德() 	feud.Barony
    BQizilcha齐兹尔查() 	feud.Barony
    BUkrach乌拉迟() 	feud.Barony
}

type 那色波NakhshabCounty struct {
	feud.BaseCounty
	博扎里克Bozariq 	feud.Barony
	菲拉布Firabr 	feud.Barony
	米尔佐莫明Mirzomomin 	feud.Barony
	那色波Nakhsab 	feud.Barony
	拜坎德Paykand 	feud.Barony
	齐兹尔查Qizilcha 	feud.Barony
	乌拉迟Ukrach 	feud.Barony
}

func (c *那色波NakhshabCounty) BBozariq博扎里克() feud.Barony {
	return c.博扎里克Bozariq
}
    
func (c *那色波NakhshabCounty) BFirabr菲拉布() feud.Barony {
	return c.菲拉布Firabr
}
    
func (c *那色波NakhshabCounty) BMirzomomin米尔佐莫明() feud.Barony {
	return c.米尔佐莫明Mirzomomin
}
    
func (c *那色波NakhshabCounty) BNakhsab那色波() feud.Barony {
	return c.那色波Nakhsab
}
    
func (c *那色波NakhshabCounty) BPaykand拜坎德() feud.Barony {
	return c.拜坎德Paykand
}
    
func (c *那色波NakhshabCounty) BQizilcha齐兹尔查() feud.Barony {
	return c.齐兹尔查Qizilcha
}
    
func (c *那色波NakhshabCounty) BUkrach乌拉迟() feud.Barony {
	return c.乌拉迟Ukrach
}
    
var CNakhshab那色波 NakhshabCounty = &那色波NakhshabCounty{}

func init() {
	f := CNakhshab那色波.(*那色波NakhshabCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1545",
		Title:     "nakhshab",
		TitleName: "那色波",
		TitleCode: "c_nakhshab",
		Baronies:  map[string]feud.Barony{},
	}

	f.博扎里克Bozariq = BBozariq博扎里克
	f.博扎里克Bozariq.SetParent(f)
	
	f.菲拉布Firabr = BFirabr菲拉布
	f.菲拉布Firabr.SetParent(f)
	
	f.米尔佐莫明Mirzomomin = BMirzomomin米尔佐莫明
	f.米尔佐莫明Mirzomomin.SetParent(f)
	
	f.那色波Nakhsab = BNakhsab那色波
	f.那色波Nakhsab.SetParent(f)
	
	f.拜坎德Paykand = BPaykand拜坎德
	f.拜坎德Paykand.SetParent(f)
	
	f.齐兹尔查Qizilcha = BQizilcha齐兹尔查
	f.齐兹尔查Qizilcha.SetParent(f)
	
	f.乌拉迟Ukrach = BUkrach乌拉迟
	f.乌拉迟Ukrach.SetParent(f)
	
}
