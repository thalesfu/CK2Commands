package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DunbarCounty interface {
    feud.County
    BBerwick伯威克() 	feud.Barony
    BColdingham科尔丁厄姆() 	feud.Barony
    BCrichton克赖顿() 	feud.Barony
    BDunbar邓巴() 	feud.Barony
    BGordon戈登() 	feud.Barony
    BHuntly亨特利() 	feud.Barony
    BThirlestane瑟尔斯坦() 	feud.Barony
    BTyninghame泰宁汉姆() 	feud.Barony
}

type 邓巴DunbarCounty struct {
	feud.BaseCounty
	伯威克Berwick 	feud.Barony
	科尔丁厄姆Coldingham 	feud.Barony
	克赖顿Crichton 	feud.Barony
	邓巴Dunbar 	feud.Barony
	戈登Gordon 	feud.Barony
	亨特利Huntly 	feud.Barony
	瑟尔斯坦Thirlestane 	feud.Barony
	泰宁汉姆Tyninghame 	feud.Barony
}

func (c *邓巴DunbarCounty) BBerwick伯威克() feud.Barony {
	return c.伯威克Berwick
}
    
func (c *邓巴DunbarCounty) BColdingham科尔丁厄姆() feud.Barony {
	return c.科尔丁厄姆Coldingham
}
    
func (c *邓巴DunbarCounty) BCrichton克赖顿() feud.Barony {
	return c.克赖顿Crichton
}
    
func (c *邓巴DunbarCounty) BDunbar邓巴() feud.Barony {
	return c.邓巴Dunbar
}
    
func (c *邓巴DunbarCounty) BGordon戈登() feud.Barony {
	return c.戈登Gordon
}
    
func (c *邓巴DunbarCounty) BHuntly亨特利() feud.Barony {
	return c.亨特利Huntly
}
    
func (c *邓巴DunbarCounty) BThirlestane瑟尔斯坦() feud.Barony {
	return c.瑟尔斯坦Thirlestane
}
    
func (c *邓巴DunbarCounty) BTyninghame泰宁汉姆() feud.Barony {
	return c.泰宁汉姆Tyninghame
}
    
var CDunbar邓巴 DunbarCounty = &邓巴DunbarCounty{}

func init() {
	f := CDunbar邓巴.(*邓巴DunbarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "51",
		Title:     "dunbar",
		TitleName: "邓巴",
		TitleCode: "c_dunbar",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯威克Berwick = BBerwick伯威克
	f.伯威克Berwick.SetParent(f)
	
	f.科尔丁厄姆Coldingham = BColdingham科尔丁厄姆
	f.科尔丁厄姆Coldingham.SetParent(f)
	
	f.克赖顿Crichton = BCrichton克赖顿
	f.克赖顿Crichton.SetParent(f)
	
	f.邓巴Dunbar = BDunbar邓巴
	f.邓巴Dunbar.SetParent(f)
	
	f.戈登Gordon = BGordon戈登
	f.戈登Gordon.SetParent(f)
	
	f.亨特利Huntly = BHuntly亨特利
	f.亨特利Huntly.SetParent(f)
	
	f.瑟尔斯坦Thirlestane = BThirlestane瑟尔斯坦
	f.瑟尔斯坦Thirlestane.SetParent(f)
	
	f.泰宁汉姆Tyninghame = BTyninghame泰宁汉姆
	f.泰宁汉姆Tyninghame.SetParent(f)
	
}
