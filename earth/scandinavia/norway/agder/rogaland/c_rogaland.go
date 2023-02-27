package rogaland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RogalandCounty interface {
    feud.County
    BBygdeborg布格德堡() 	feud.Barony
    BEikundarsund埃昆德松() 	feud.Barony
    BHesby赫斯比() 	feud.Barony
    BJonegarden约尼格达登() 	feud.Barony
    BKlepp克莱普() 	feud.Barony
    BNaerbo内尔伯() 	feud.Barony
    BRoldal勒尔达尔() 	feud.Barony
    BStavanger斯塔万格() 	feud.Barony
}

type 罗加兰RogalandCounty struct {
	feud.BaseCounty
	布格德堡Bygdeborg 	feud.Barony
	埃昆德松Eikundarsund 	feud.Barony
	赫斯比Hesby 	feud.Barony
	约尼格达登Jonegarden 	feud.Barony
	克莱普Klepp 	feud.Barony
	内尔伯Naerbo 	feud.Barony
	勒尔达尔Roldal 	feud.Barony
	斯塔万格Stavanger 	feud.Barony
}

func (c *罗加兰RogalandCounty) BBygdeborg布格德堡() feud.Barony {
	return c.布格德堡Bygdeborg
}
    
func (c *罗加兰RogalandCounty) BEikundarsund埃昆德松() feud.Barony {
	return c.埃昆德松Eikundarsund
}
    
func (c *罗加兰RogalandCounty) BHesby赫斯比() feud.Barony {
	return c.赫斯比Hesby
}
    
func (c *罗加兰RogalandCounty) BJonegarden约尼格达登() feud.Barony {
	return c.约尼格达登Jonegarden
}
    
func (c *罗加兰RogalandCounty) BKlepp克莱普() feud.Barony {
	return c.克莱普Klepp
}
    
func (c *罗加兰RogalandCounty) BNaerbo内尔伯() feud.Barony {
	return c.内尔伯Naerbo
}
    
func (c *罗加兰RogalandCounty) BRoldal勒尔达尔() feud.Barony {
	return c.勒尔达尔Roldal
}
    
func (c *罗加兰RogalandCounty) BStavanger斯塔万格() feud.Barony {
	return c.斯塔万格Stavanger
}
    
var CRogaland罗加兰 RogalandCounty = &罗加兰RogalandCounty{}

func init() {
	f := CRogaland罗加兰.(*罗加兰RogalandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "269",
		Title:     "rogaland",
		TitleName: "罗加兰",
		TitleCode: "c_rogaland",
		Baronies:  map[string]feud.Barony{},
	}

	f.布格德堡Bygdeborg = BBygdeborg布格德堡
	f.布格德堡Bygdeborg.SetParent(f)
	
	f.埃昆德松Eikundarsund = BEikundarsund埃昆德松
	f.埃昆德松Eikundarsund.SetParent(f)
	
	f.赫斯比Hesby = BHesby赫斯比
	f.赫斯比Hesby.SetParent(f)
	
	f.约尼格达登Jonegarden = BJonegarden约尼格达登
	f.约尼格达登Jonegarden.SetParent(f)
	
	f.克莱普Klepp = BKlepp克莱普
	f.克莱普Klepp.SetParent(f)
	
	f.内尔伯Naerbo = BNaerbo内尔伯
	f.内尔伯Naerbo.SetParent(f)
	
	f.勒尔达尔Roldal = BRoldal勒尔达尔
	f.勒尔达尔Roldal.SetParent(f)
	
	f.斯塔万格Stavanger = BStavanger斯塔万格
	f.斯塔万格Stavanger.SetParent(f)
	
}
