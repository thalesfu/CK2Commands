package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BremenCounty interface {
    feud.County
    BAchim阿希姆() 	feud.Barony
    BBeverstedt贝沃施泰特() 	feud.Barony
    BBlumenthal布卢门塔尔() 	feud.Barony
    BHoya霍亚() 	feud.Barony
    BOsterbruch奥斯特布鲁赫() 	feud.Barony
    BRitzebuttel里策比特尔() 	feud.Barony
    BStade施塔德() 	feud.Barony
}

type 哈德尔恩BremenCounty struct {
	feud.BaseCounty
	阿希姆Achim 	feud.Barony
	贝沃施泰特Beverstedt 	feud.Barony
	布卢门塔尔Blumenthal 	feud.Barony
	霍亚Hoya 	feud.Barony
	奥斯特布鲁赫Osterbruch 	feud.Barony
	里策比特尔Ritzebuttel 	feud.Barony
	施塔德Stade 	feud.Barony
}

func (c *哈德尔恩BremenCounty) BAchim阿希姆() feud.Barony {
	return c.阿希姆Achim
}
    
func (c *哈德尔恩BremenCounty) BBeverstedt贝沃施泰特() feud.Barony {
	return c.贝沃施泰特Beverstedt
}
    
func (c *哈德尔恩BremenCounty) BBlumenthal布卢门塔尔() feud.Barony {
	return c.布卢门塔尔Blumenthal
}
    
func (c *哈德尔恩BremenCounty) BHoya霍亚() feud.Barony {
	return c.霍亚Hoya
}
    
func (c *哈德尔恩BremenCounty) BOsterbruch奥斯特布鲁赫() feud.Barony {
	return c.奥斯特布鲁赫Osterbruch
}
    
func (c *哈德尔恩BremenCounty) BRitzebuttel里策比特尔() feud.Barony {
	return c.里策比特尔Ritzebuttel
}
    
func (c *哈德尔恩BremenCounty) BStade施塔德() feud.Barony {
	return c.施塔德Stade
}
    
var CBremen哈德尔恩 BremenCounty = &哈德尔恩BremenCounty{}

func init() {
	f := CBremen哈德尔恩.(*哈德尔恩BremenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "848",
		Title:     "bremen",
		TitleName: "哈德尔恩",
		TitleCode: "c_bremen",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿希姆Achim = BAchim阿希姆
	f.阿希姆Achim.SetParent(f)
	
	f.贝沃施泰特Beverstedt = BBeverstedt贝沃施泰特
	f.贝沃施泰特Beverstedt.SetParent(f)
	
	f.布卢门塔尔Blumenthal = BBlumenthal布卢门塔尔
	f.布卢门塔尔Blumenthal.SetParent(f)
	
	f.霍亚Hoya = BHoya霍亚
	f.霍亚Hoya.SetParent(f)
	
	f.奥斯特布鲁赫Osterbruch = BOsterbruch奥斯特布鲁赫
	f.奥斯特布鲁赫Osterbruch.SetParent(f)
	
	f.里策比特尔Ritzebuttel = BRitzebuttel里策比特尔
	f.里策比特尔Ritzebuttel.SetParent(f)
	
	f.施塔德Stade = BStade施塔德
	f.施塔德Stade.SetParent(f)
	
}
