package banbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BanbarCounty interface {
    feud.County
    BBanbar边坝() 	feud.Barony
    BCoka草卡() 	feud.Barony
    BDengqen丁青() 	feud.Barony
    BDomartang东马塘() 	feud.Barony
    BEgage俄嘎格() 	feud.Barony
    BSadeng沙丁() 	feud.Barony
    BShagong沙贡() 	feud.Barony
}

type 边坝BanbarCounty struct {
	feud.BaseCounty
	边坝Banbar 	feud.Barony
	草卡Coka 	feud.Barony
	丁青Dengqen 	feud.Barony
	东马塘Domartang 	feud.Barony
	俄嘎格Egage 	feud.Barony
	沙丁Sadeng 	feud.Barony
	沙贡Shagong 	feud.Barony
}

func (c *边坝BanbarCounty) BBanbar边坝() feud.Barony {
	return c.边坝Banbar
}
    
func (c *边坝BanbarCounty) BCoka草卡() feud.Barony {
	return c.草卡Coka
}
    
func (c *边坝BanbarCounty) BDengqen丁青() feud.Barony {
	return c.丁青Dengqen
}
    
func (c *边坝BanbarCounty) BDomartang东马塘() feud.Barony {
	return c.东马塘Domartang
}
    
func (c *边坝BanbarCounty) BEgage俄嘎格() feud.Barony {
	return c.俄嘎格Egage
}
    
func (c *边坝BanbarCounty) BSadeng沙丁() feud.Barony {
	return c.沙丁Sadeng
}
    
func (c *边坝BanbarCounty) BShagong沙贡() feud.Barony {
	return c.沙贡Shagong
}
    
var CBanbar边坝 BanbarCounty = &边坝BanbarCounty{}

func init() {
	f := CBanbar边坝.(*边坝BanbarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1554",
		Title:     "banbar",
		TitleName: "边坝",
		TitleCode: "c_banbar",
		Baronies:  map[string]feud.Barony{},
	}

	f.边坝Banbar = BBanbar边坝
	f.边坝Banbar.SetParent(f)
	
	f.草卡Coka = BCoka草卡
	f.草卡Coka.SetParent(f)
	
	f.丁青Dengqen = BDengqen丁青
	f.丁青Dengqen.SetParent(f)
	
	f.东马塘Domartang = BDomartang东马塘
	f.东马塘Domartang.SetParent(f)
	
	f.俄嘎格Egage = BEgage俄嘎格
	f.俄嘎格Egage.SetParent(f)
	
	f.沙丁Sadeng = BSadeng沙丁
	f.沙丁Sadeng.SetParent(f)
	
	f.沙贡Shagong = BShagong沙贡
	f.沙贡Shagong.SetParent(f)
	
}
