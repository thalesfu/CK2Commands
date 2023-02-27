package bolshoy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BolshoyCounty interface {
    feud.County
    BBanyy巴内伊() 	feud.Barony
    BBolshoy大尤甘() 	feud.Barony
    BLempino连皮诺() 	feud.Barony
    BSingapay辛加派() 	feud.Barony
    BSytomino瑟托米诺() 	feud.Barony
    BTundrino通德里诺() 	feud.Barony
    BYugan尤甘() 	feud.Barony
}

type 大尤甘BolshoyCounty struct {
	feud.BaseCounty
	巴内伊Banyy 	feud.Barony
	大尤甘Bolshoy 	feud.Barony
	连皮诺Lempino 	feud.Barony
	辛加派Singapay 	feud.Barony
	瑟托米诺Sytomino 	feud.Barony
	通德里诺Tundrino 	feud.Barony
	尤甘Yugan 	feud.Barony
}

func (c *大尤甘BolshoyCounty) BBanyy巴内伊() feud.Barony {
	return c.巴内伊Banyy
}
    
func (c *大尤甘BolshoyCounty) BBolshoy大尤甘() feud.Barony {
	return c.大尤甘Bolshoy
}
    
func (c *大尤甘BolshoyCounty) BLempino连皮诺() feud.Barony {
	return c.连皮诺Lempino
}
    
func (c *大尤甘BolshoyCounty) BSingapay辛加派() feud.Barony {
	return c.辛加派Singapay
}
    
func (c *大尤甘BolshoyCounty) BSytomino瑟托米诺() feud.Barony {
	return c.瑟托米诺Sytomino
}
    
func (c *大尤甘BolshoyCounty) BTundrino通德里诺() feud.Barony {
	return c.通德里诺Tundrino
}
    
func (c *大尤甘BolshoyCounty) BYugan尤甘() feud.Barony {
	return c.尤甘Yugan
}
    
var CBolshoy大尤甘 BolshoyCounty = &大尤甘BolshoyCounty{}

func init() {
	f := CBolshoy大尤甘.(*大尤甘BolshoyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1859",
		Title:     "bolshoy",
		TitleName: "大尤甘",
		TitleCode: "c_bolshoy",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴内伊Banyy = BBanyy巴内伊
	f.巴内伊Banyy.SetParent(f)
	
	f.大尤甘Bolshoy = BBolshoy大尤甘
	f.大尤甘Bolshoy.SetParent(f)
	
	f.连皮诺Lempino = BLempino连皮诺
	f.连皮诺Lempino.SetParent(f)
	
	f.辛加派Singapay = BSingapay辛加派
	f.辛加派Singapay.SetParent(f)
	
	f.瑟托米诺Sytomino = BSytomino瑟托米诺
	f.瑟托米诺Sytomino.SetParent(f)
	
	f.通德里诺Tundrino = BTundrino通德里诺
	f.通德里诺Tundrino.SetParent(f)
	
	f.尤甘Yugan = BYugan尤甘
	f.尤甘Yugan.SetParent(f)
	
}
