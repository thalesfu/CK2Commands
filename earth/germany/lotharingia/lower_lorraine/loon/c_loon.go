package loon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LoonCounty interface {
    feud.County
    BHasselt哈瑟尔特() 	feud.Barony
    BHeinsberg海因斯贝格() 	feud.Barony
    BHesbaie埃斯拜() 	feud.Barony
    BLoon洛翁() 	feud.Barony
    BMaastricht马斯特里赫特() 	feud.Barony
    BTongeren通厄伦() 	feud.Barony
    BValkenburg法尔肯堡() 	feud.Barony
    BWassenberg瓦森贝格() 	feud.Barony
}

type 洛翁LoonCounty struct {
	feud.BaseCounty
	哈瑟尔特Hasselt 	feud.Barony
	海因斯贝格Heinsberg 	feud.Barony
	埃斯拜Hesbaie 	feud.Barony
	洛翁Loon 	feud.Barony
	马斯特里赫特Maastricht 	feud.Barony
	通厄伦Tongeren 	feud.Barony
	法尔肯堡Valkenburg 	feud.Barony
	瓦森贝格Wassenberg 	feud.Barony
}

func (c *洛翁LoonCounty) BHasselt哈瑟尔特() feud.Barony {
	return c.哈瑟尔特Hasselt
}
    
func (c *洛翁LoonCounty) BHeinsberg海因斯贝格() feud.Barony {
	return c.海因斯贝格Heinsberg
}
    
func (c *洛翁LoonCounty) BHesbaie埃斯拜() feud.Barony {
	return c.埃斯拜Hesbaie
}
    
func (c *洛翁LoonCounty) BLoon洛翁() feud.Barony {
	return c.洛翁Loon
}
    
func (c *洛翁LoonCounty) BMaastricht马斯特里赫特() feud.Barony {
	return c.马斯特里赫特Maastricht
}
    
func (c *洛翁LoonCounty) BTongeren通厄伦() feud.Barony {
	return c.通厄伦Tongeren
}
    
func (c *洛翁LoonCounty) BValkenburg法尔肯堡() feud.Barony {
	return c.法尔肯堡Valkenburg
}
    
func (c *洛翁LoonCounty) BWassenberg瓦森贝格() feud.Barony {
	return c.瓦森贝格Wassenberg
}
    
var CLoon洛翁 LoonCounty = &洛翁LoonCounty{}

func init() {
	f := CLoon洛翁.(*洛翁LoonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "91",
		Title:     "loon",
		TitleName: "洛翁",
		TitleCode: "c_loon",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈瑟尔特Hasselt = BHasselt哈瑟尔特
	f.哈瑟尔特Hasselt.SetParent(f)
	
	f.海因斯贝格Heinsberg = BHeinsberg海因斯贝格
	f.海因斯贝格Heinsberg.SetParent(f)
	
	f.埃斯拜Hesbaie = BHesbaie埃斯拜
	f.埃斯拜Hesbaie.SetParent(f)
	
	f.洛翁Loon = BLoon洛翁
	f.洛翁Loon.SetParent(f)
	
	f.马斯特里赫特Maastricht = BMaastricht马斯特里赫特
	f.马斯特里赫特Maastricht.SetParent(f)
	
	f.通厄伦Tongeren = BTongeren通厄伦
	f.通厄伦Tongeren.SetParent(f)
	
	f.法尔肯堡Valkenburg = BValkenburg法尔肯堡
	f.法尔肯堡Valkenburg.SetParent(f)
	
	f.瓦森贝格Wassenberg = BWassenberg瓦森贝格
	f.瓦森贝格Wassenberg.SetParent(f)
	
}
