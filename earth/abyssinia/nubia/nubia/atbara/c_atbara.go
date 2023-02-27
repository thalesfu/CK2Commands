package atbara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AtbaraCounty interface {
    feud.County
    BAbudis阿布迪斯() 	feud.Barony
    BAtbara阿特巴拉() 	feud.Barony
    BBegarawiyah比格拉维亚() 	feud.Barony
    BHajaralasla哈贾雷阿萨() 	feud.Barony
    BMeruwah麦罗埃() 	feud.Barony
    BShendi尚迪() 	feud.Barony
    BUmmmardi乌姆迈迪() 	feud.Barony
}

type 阿特巴拉AtbaraCounty struct {
	feud.BaseCounty
	阿布迪斯Abudis 	feud.Barony
	阿特巴拉Atbara 	feud.Barony
	比格拉维亚Begarawiyah 	feud.Barony
	哈贾雷阿萨Hajaralasla 	feud.Barony
	麦罗埃Meruwah 	feud.Barony
	尚迪Shendi 	feud.Barony
	乌姆迈迪Ummmardi 	feud.Barony
}

func (c *阿特巴拉AtbaraCounty) BAbudis阿布迪斯() feud.Barony {
	return c.阿布迪斯Abudis
}
    
func (c *阿特巴拉AtbaraCounty) BAtbara阿特巴拉() feud.Barony {
	return c.阿特巴拉Atbara
}
    
func (c *阿特巴拉AtbaraCounty) BBegarawiyah比格拉维亚() feud.Barony {
	return c.比格拉维亚Begarawiyah
}
    
func (c *阿特巴拉AtbaraCounty) BHajaralasla哈贾雷阿萨() feud.Barony {
	return c.哈贾雷阿萨Hajaralasla
}
    
func (c *阿特巴拉AtbaraCounty) BMeruwah麦罗埃() feud.Barony {
	return c.麦罗埃Meruwah
}
    
func (c *阿特巴拉AtbaraCounty) BShendi尚迪() feud.Barony {
	return c.尚迪Shendi
}
    
func (c *阿特巴拉AtbaraCounty) BUmmmardi乌姆迈迪() feud.Barony {
	return c.乌姆迈迪Ummmardi
}
    
var CAtbara阿特巴拉 AtbaraCounty = &阿特巴拉AtbaraCounty{}

func init() {
	f := CAtbara阿特巴拉.(*阿特巴拉AtbaraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "879",
		Title:     "atbara",
		TitleName: "阿特巴拉",
		TitleCode: "c_atbara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布迪斯Abudis = BAbudis阿布迪斯
	f.阿布迪斯Abudis.SetParent(f)
	
	f.阿特巴拉Atbara = BAtbara阿特巴拉
	f.阿特巴拉Atbara.SetParent(f)
	
	f.比格拉维亚Begarawiyah = BBegarawiyah比格拉维亚
	f.比格拉维亚Begarawiyah.SetParent(f)
	
	f.哈贾雷阿萨Hajaralasla = BHajaralasla哈贾雷阿萨
	f.哈贾雷阿萨Hajaralasla.SetParent(f)
	
	f.麦罗埃Meruwah = BMeruwah麦罗埃
	f.麦罗埃Meruwah.SetParent(f)
	
	f.尚迪Shendi = BShendi尚迪
	f.尚迪Shendi.SetParent(f)
	
	f.乌姆迈迪Ummmardi = BUmmmardi乌姆迈迪
	f.乌姆迈迪Ummmardi.SetParent(f)
	
}
