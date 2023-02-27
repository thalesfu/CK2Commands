package nisibin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NisibinCounty interface {
    feud.County
    BCizre茨雷() 	feud.Barony
    BDairodmorhannanyo达罗摩汗那尤() 	feud.Barony
    BDayrodmorgabriel达罗摩尔格波利() 	feud.Barony
    BKerburan克尔布兰() 	feud.Barony
    BNusaybin努赛宾() 	feud.Barony
    BSavur萨武尔() 	feud.Barony
    BSittiradviyyemadrasa斯提拉德维耶马德雷萨() 	feud.Barony
    BYaqobnsibnaya亚喀尼斯纳亚() 	feud.Barony
}

type 尼西宾NisibinCounty struct {
	feud.BaseCounty
	茨雷Cizre 	feud.Barony
	达罗摩汗那尤Dairodmorhannanyo 	feud.Barony
	达罗摩尔格波利Dayrodmorgabriel 	feud.Barony
	克尔布兰Kerburan 	feud.Barony
	努赛宾Nusaybin 	feud.Barony
	萨武尔Savur 	feud.Barony
	斯提拉德维耶马德雷萨Sittiradviyyemadrasa 	feud.Barony
	亚喀尼斯纳亚Yaqobnsibnaya 	feud.Barony
}

func (c *尼西宾NisibinCounty) BCizre茨雷() feud.Barony {
	return c.茨雷Cizre
}
    
func (c *尼西宾NisibinCounty) BDairodmorhannanyo达罗摩汗那尤() feud.Barony {
	return c.达罗摩汗那尤Dairodmorhannanyo
}
    
func (c *尼西宾NisibinCounty) BDayrodmorgabriel达罗摩尔格波利() feud.Barony {
	return c.达罗摩尔格波利Dayrodmorgabriel
}
    
func (c *尼西宾NisibinCounty) BKerburan克尔布兰() feud.Barony {
	return c.克尔布兰Kerburan
}
    
func (c *尼西宾NisibinCounty) BNusaybin努赛宾() feud.Barony {
	return c.努赛宾Nusaybin
}
    
func (c *尼西宾NisibinCounty) BSavur萨武尔() feud.Barony {
	return c.萨武尔Savur
}
    
func (c *尼西宾NisibinCounty) BSittiradviyyemadrasa斯提拉德维耶马德雷萨() feud.Barony {
	return c.斯提拉德维耶马德雷萨Sittiradviyyemadrasa
}
    
func (c *尼西宾NisibinCounty) BYaqobnsibnaya亚喀尼斯纳亚() feud.Barony {
	return c.亚喀尼斯纳亚Yaqobnsibnaya
}
    
var CNisibin尼西宾 NisibinCounty = &尼西宾NisibinCounty{}

func init() {
	f := CNisibin尼西宾.(*尼西宾NisibinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "684",
		Title:     "nisibin",
		TitleName: "尼西宾",
		TitleCode: "c_nisibin",
		Baronies:  map[string]feud.Barony{},
	}

	f.茨雷Cizre = BCizre茨雷
	f.茨雷Cizre.SetParent(f)
	
	f.达罗摩汗那尤Dairodmorhannanyo = BDairodmorhannanyo达罗摩汗那尤
	f.达罗摩汗那尤Dairodmorhannanyo.SetParent(f)
	
	f.达罗摩尔格波利Dayrodmorgabriel = BDayrodmorgabriel达罗摩尔格波利
	f.达罗摩尔格波利Dayrodmorgabriel.SetParent(f)
	
	f.克尔布兰Kerburan = BKerburan克尔布兰
	f.克尔布兰Kerburan.SetParent(f)
	
	f.努赛宾Nusaybin = BNusaybin努赛宾
	f.努赛宾Nusaybin.SetParent(f)
	
	f.萨武尔Savur = BSavur萨武尔
	f.萨武尔Savur.SetParent(f)
	
	f.斯提拉德维耶马德雷萨Sittiradviyyemadrasa = BSittiradviyyemadrasa斯提拉德维耶马德雷萨
	f.斯提拉德维耶马德雷萨Sittiradviyyemadrasa.SetParent(f)
	
	f.亚喀尼斯纳亚Yaqobnsibnaya = BYaqobnsibnaya亚喀尼斯纳亚
	f.亚喀尼斯纳亚Yaqobnsibnaya.SetParent(f)
	
}
