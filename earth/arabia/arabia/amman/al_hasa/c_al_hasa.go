package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_hasaCounty interface {
    feud.County
    BAbqaiq布盖格() 	feud.Barony
    BAlhasa哈萨() 	feud.Barony
    BAlmubarraz穆拜赖兹() 	feud.Barony
    BAlomran奥姆兰() 	feud.Barony
    BAloyoon奥约恩() 	feud.Barony
    BFoda佛达() 	feud.Barony
    BGhunan吉胡纳() 	feud.Barony
    BHoluf胡富夫() 	feud.Barony
    BKhobar胡拜尔() 	feud.Barony
}

type 哈萨Al_hasaCounty struct {
	feud.BaseCounty
	布盖格Abqaiq 	feud.Barony
	哈萨Alhasa 	feud.Barony
	穆拜赖兹Almubarraz 	feud.Barony
	奥姆兰Alomran 	feud.Barony
	奥约恩Aloyoon 	feud.Barony
	佛达Foda 	feud.Barony
	吉胡纳Ghunan 	feud.Barony
	胡富夫Holuf 	feud.Barony
	胡拜尔Khobar 	feud.Barony
}

func (c *哈萨Al_hasaCounty) BAbqaiq布盖格() feud.Barony {
	return c.布盖格Abqaiq
}
    
func (c *哈萨Al_hasaCounty) BAlhasa哈萨() feud.Barony {
	return c.哈萨Alhasa
}
    
func (c *哈萨Al_hasaCounty) BAlmubarraz穆拜赖兹() feud.Barony {
	return c.穆拜赖兹Almubarraz
}
    
func (c *哈萨Al_hasaCounty) BAlomran奥姆兰() feud.Barony {
	return c.奥姆兰Alomran
}
    
func (c *哈萨Al_hasaCounty) BAloyoon奥约恩() feud.Barony {
	return c.奥约恩Aloyoon
}
    
func (c *哈萨Al_hasaCounty) BFoda佛达() feud.Barony {
	return c.佛达Foda
}
    
func (c *哈萨Al_hasaCounty) BGhunan吉胡纳() feud.Barony {
	return c.吉胡纳Ghunan
}
    
func (c *哈萨Al_hasaCounty) BHoluf胡富夫() feud.Barony {
	return c.胡富夫Holuf
}
    
func (c *哈萨Al_hasaCounty) BKhobar胡拜尔() feud.Barony {
	return c.胡拜尔Khobar
}
    
var CAl_hasa哈萨 Al_hasaCounty = &哈萨Al_hasaCounty{}

func init() {
	f := CAl_hasa哈萨.(*哈萨Al_hasaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "652",
		Title:     "al_hasa",
		TitleName: "哈萨",
		TitleCode: "c_al_hasa",
		Baronies:  map[string]feud.Barony{},
	}

	f.布盖格Abqaiq = BAbqaiq布盖格
	f.布盖格Abqaiq.SetParent(f)
	
	f.哈萨Alhasa = BAlhasa哈萨
	f.哈萨Alhasa.SetParent(f)
	
	f.穆拜赖兹Almubarraz = BAlmubarraz穆拜赖兹
	f.穆拜赖兹Almubarraz.SetParent(f)
	
	f.奥姆兰Alomran = BAlomran奥姆兰
	f.奥姆兰Alomran.SetParent(f)
	
	f.奥约恩Aloyoon = BAloyoon奥约恩
	f.奥约恩Aloyoon.SetParent(f)
	
	f.佛达Foda = BFoda佛达
	f.佛达Foda.SetParent(f)
	
	f.吉胡纳Ghunan = BGhunan吉胡纳
	f.吉胡纳Ghunan.SetParent(f)
	
	f.胡富夫Holuf = BHoluf胡富夫
	f.胡富夫Holuf.SetParent(f)
	
	f.胡拜尔Khobar = BKhobar胡拜尔
	f.胡拜尔Khobar.SetParent(f)
	
}
