package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrentCounty interface {
    feud.County
    BBrixen布里申() 	feud.Barony
    BBruneck布鲁内克() 	feud.Barony
    BMolven莫尔韦诺() 	feud.Barony
    BRiva里瓦() 	feud.Barony
    BRovereto罗韦雷托() 	feud.Barony
    BSchlanders施兰德斯() 	feud.Barony
    BTrento特伦托() 	feud.Barony
}

type 特伦托TrentCounty struct {
	feud.BaseCounty
	布里申Brixen 	feud.Barony
	布鲁内克Bruneck 	feud.Barony
	莫尔韦诺Molven 	feud.Barony
	里瓦Riva 	feud.Barony
	罗韦雷托Rovereto 	feud.Barony
	施兰德斯Schlanders 	feud.Barony
	特伦托Trento 	feud.Barony
}

func (c *特伦托TrentCounty) BBrixen布里申() feud.Barony {
	return c.布里申Brixen
}
    
func (c *特伦托TrentCounty) BBruneck布鲁内克() feud.Barony {
	return c.布鲁内克Bruneck
}
    
func (c *特伦托TrentCounty) BMolven莫尔韦诺() feud.Barony {
	return c.莫尔韦诺Molven
}
    
func (c *特伦托TrentCounty) BRiva里瓦() feud.Barony {
	return c.里瓦Riva
}
    
func (c *特伦托TrentCounty) BRovereto罗韦雷托() feud.Barony {
	return c.罗韦雷托Rovereto
}
    
func (c *特伦托TrentCounty) BSchlanders施兰德斯() feud.Barony {
	return c.施兰德斯Schlanders
}
    
func (c *特伦托TrentCounty) BTrento特伦托() feud.Barony {
	return c.特伦托Trento
}
    
var CTrent特伦托 TrentCounty = &特伦托TrentCounty{}

func init() {
	f := CTrent特伦托.(*特伦托TrentCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "317",
		Title:     "trent",
		TitleName: "特伦托",
		TitleCode: "c_trent",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里申Brixen = BBrixen布里申
	f.布里申Brixen.SetParent(f)
	
	f.布鲁内克Bruneck = BBruneck布鲁内克
	f.布鲁内克Bruneck.SetParent(f)
	
	f.莫尔韦诺Molven = BMolven莫尔韦诺
	f.莫尔韦诺Molven.SetParent(f)
	
	f.里瓦Riva = BRiva里瓦
	f.里瓦Riva.SetParent(f)
	
	f.罗韦雷托Rovereto = BRovereto罗韦雷托
	f.罗韦雷托Rovereto.SetParent(f)
	
	f.施兰德斯Schlanders = BSchlanders施兰德斯
	f.施兰德斯Schlanders.SetParent(f)
	
	f.特伦托Trento = BTrento特伦托
	f.特伦托Trento.SetParent(f)
	
}
