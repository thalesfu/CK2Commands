package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PerfeddwladCounty interface {
    feud.County
    BBasingwerk巴辛韦克() 	feud.Barony
    BDenbigh登比() 	feud.Barony
    BEwloe尤洛() 	feud.Barony
    BFlint弗林特() 	feud.Barony
    BHawarden哈登() 	feud.Barony
    BLlanelwy兰埃尔威() 	feud.Barony
    BRhuddlan里兹兰() 	feud.Barony
    BRuthin里辛() 	feud.Barony
}

type 佩菲杜拉德PerfeddwladCounty struct {
	feud.BaseCounty
	巴辛韦克Basingwerk 	feud.Barony
	登比Denbigh 	feud.Barony
	尤洛Ewloe 	feud.Barony
	弗林特Flint 	feud.Barony
	哈登Hawarden 	feud.Barony
	兰埃尔威Llanelwy 	feud.Barony
	里兹兰Rhuddlan 	feud.Barony
	里辛Ruthin 	feud.Barony
}

func (c *佩菲杜拉德PerfeddwladCounty) BBasingwerk巴辛韦克() feud.Barony {
	return c.巴辛韦克Basingwerk
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BDenbigh登比() feud.Barony {
	return c.登比Denbigh
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BEwloe尤洛() feud.Barony {
	return c.尤洛Ewloe
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BFlint弗林特() feud.Barony {
	return c.弗林特Flint
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BHawarden哈登() feud.Barony {
	return c.哈登Hawarden
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BLlanelwy兰埃尔威() feud.Barony {
	return c.兰埃尔威Llanelwy
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BRhuddlan里兹兰() feud.Barony {
	return c.里兹兰Rhuddlan
}
    
func (c *佩菲杜拉德PerfeddwladCounty) BRuthin里辛() feud.Barony {
	return c.里辛Ruthin
}
    
var CPerfeddwlad佩菲杜拉德 PerfeddwladCounty = &佩菲杜拉德PerfeddwladCounty{}

func init() {
	f := CPerfeddwlad佩菲杜拉德.(*佩菲杜拉德PerfeddwladCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "60",
		Title:     "perfeddwlad",
		TitleName: "佩菲杜拉德",
		TitleCode: "c_perfeddwlad",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴辛韦克Basingwerk = BBasingwerk巴辛韦克
	f.巴辛韦克Basingwerk.SetParent(f)
	
	f.登比Denbigh = BDenbigh登比
	f.登比Denbigh.SetParent(f)
	
	f.尤洛Ewloe = BEwloe尤洛
	f.尤洛Ewloe.SetParent(f)
	
	f.弗林特Flint = BFlint弗林特
	f.弗林特Flint.SetParent(f)
	
	f.哈登Hawarden = BHawarden哈登
	f.哈登Hawarden.SetParent(f)
	
	f.兰埃尔威Llanelwy = BLlanelwy兰埃尔威
	f.兰埃尔威Llanelwy.SetParent(f)
	
	f.里兹兰Rhuddlan = BRhuddlan里兹兰
	f.里兹兰Rhuddlan.SetParent(f)
	
	f.里辛Ruthin = BRuthin里辛
	f.里辛Ruthin.SetParent(f)
	
}
