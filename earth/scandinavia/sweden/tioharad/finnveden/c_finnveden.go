package finnveden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinnvedenCounty interface {
    feud.County
    BBolmso博尔姆瑟() 	feud.Barony
    BGislaved伊斯拉韦德() 	feud.Barony
    BLjungby永比() 	feud.Barony
    BMarkaryd马尔卡吕德() 	feud.Barony
    BOlmestad厄尔默斯塔德() 	feud.Barony
    BPiksborg皮克斯堡() 	feud.Barony
    BVarnamo韦纳穆() 	feud.Barony
}

type 芬韦登FinnvedenCounty struct {
	feud.BaseCounty
	博尔姆瑟Bolmso 	feud.Barony
	伊斯拉韦德Gislaved 	feud.Barony
	永比Ljungby 	feud.Barony
	马尔卡吕德Markaryd 	feud.Barony
	厄尔默斯塔德Olmestad 	feud.Barony
	皮克斯堡Piksborg 	feud.Barony
	韦纳穆Varnamo 	feud.Barony
}

func (c *芬韦登FinnvedenCounty) BBolmso博尔姆瑟() feud.Barony {
	return c.博尔姆瑟Bolmso
}
    
func (c *芬韦登FinnvedenCounty) BGislaved伊斯拉韦德() feud.Barony {
	return c.伊斯拉韦德Gislaved
}
    
func (c *芬韦登FinnvedenCounty) BLjungby永比() feud.Barony {
	return c.永比Ljungby
}
    
func (c *芬韦登FinnvedenCounty) BMarkaryd马尔卡吕德() feud.Barony {
	return c.马尔卡吕德Markaryd
}
    
func (c *芬韦登FinnvedenCounty) BOlmestad厄尔默斯塔德() feud.Barony {
	return c.厄尔默斯塔德Olmestad
}
    
func (c *芬韦登FinnvedenCounty) BPiksborg皮克斯堡() feud.Barony {
	return c.皮克斯堡Piksborg
}
    
func (c *芬韦登FinnvedenCounty) BVarnamo韦纳穆() feud.Barony {
	return c.韦纳穆Varnamo
}
    
var CFinnveden芬韦登 FinnvedenCounty = &芬韦登FinnvedenCounty{}

func init() {
	f := CFinnveden芬韦登.(*芬韦登FinnvedenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1702",
		Title:     "finnveden",
		TitleName: "芬韦登",
		TitleCode: "c_finnveden",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔姆瑟Bolmso = BBolmso博尔姆瑟
	f.博尔姆瑟Bolmso.SetParent(f)
	
	f.伊斯拉韦德Gislaved = BGislaved伊斯拉韦德
	f.伊斯拉韦德Gislaved.SetParent(f)
	
	f.永比Ljungby = BLjungby永比
	f.永比Ljungby.SetParent(f)
	
	f.马尔卡吕德Markaryd = BMarkaryd马尔卡吕德
	f.马尔卡吕德Markaryd.SetParent(f)
	
	f.厄尔默斯塔德Olmestad = BOlmestad厄尔默斯塔德
	f.厄尔默斯塔德Olmestad.SetParent(f)
	
	f.皮克斯堡Piksborg = BPiksborg皮克斯堡
	f.皮克斯堡Piksborg.SetParent(f)
	
	f.韦纳穆Varnamo = BVarnamo韦纳穆
	f.韦纳穆Varnamo.SetParent(f)
	
}
