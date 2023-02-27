package lubeck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LubeckCounty interface {
    feud.County
    BBucu布库() 	feud.Barony
    BLubeck吕贝克() 	feud.Barony
    BRatzeburg拉策堡() 	feud.Barony
    BSchlutup施卢图普() 	feud.Barony
    BStarigard斯塔里加德() 	feud.Barony
    BTravemunde特拉弗明德() 	feud.Barony
    BWeslo韦斯勒() 	feud.Barony
    BWulfsdorf伍尔夫斯多夫() 	feud.Barony
}

type 吕贝克LubeckCounty struct {
	feud.BaseCounty
	布库Bucu 	feud.Barony
	吕贝克Lubeck 	feud.Barony
	拉策堡Ratzeburg 	feud.Barony
	施卢图普Schlutup 	feud.Barony
	斯塔里加德Starigard 	feud.Barony
	特拉弗明德Travemunde 	feud.Barony
	韦斯勒Weslo 	feud.Barony
	伍尔夫斯多夫Wulfsdorf 	feud.Barony
}

func (c *吕贝克LubeckCounty) BBucu布库() feud.Barony {
	return c.布库Bucu
}
    
func (c *吕贝克LubeckCounty) BLubeck吕贝克() feud.Barony {
	return c.吕贝克Lubeck
}
    
func (c *吕贝克LubeckCounty) BRatzeburg拉策堡() feud.Barony {
	return c.拉策堡Ratzeburg
}
    
func (c *吕贝克LubeckCounty) BSchlutup施卢图普() feud.Barony {
	return c.施卢图普Schlutup
}
    
func (c *吕贝克LubeckCounty) BStarigard斯塔里加德() feud.Barony {
	return c.斯塔里加德Starigard
}
    
func (c *吕贝克LubeckCounty) BTravemunde特拉弗明德() feud.Barony {
	return c.特拉弗明德Travemunde
}
    
func (c *吕贝克LubeckCounty) BWeslo韦斯勒() feud.Barony {
	return c.韦斯勒Weslo
}
    
func (c *吕贝克LubeckCounty) BWulfsdorf伍尔夫斯多夫() feud.Barony {
	return c.伍尔夫斯多夫Wulfsdorf
}
    
var CLubeck吕贝克 LubeckCounty = &吕贝克LubeckCounty{}

func init() {
	f := CLubeck吕贝克.(*吕贝克LubeckCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "262",
		Title:     "lubeck",
		TitleName: "吕贝克",
		TitleCode: "c_lubeck",
		Baronies:  map[string]feud.Barony{},
	}

	f.布库Bucu = BBucu布库
	f.布库Bucu.SetParent(f)
	
	f.吕贝克Lubeck = BLubeck吕贝克
	f.吕贝克Lubeck.SetParent(f)
	
	f.拉策堡Ratzeburg = BRatzeburg拉策堡
	f.拉策堡Ratzeburg.SetParent(f)
	
	f.施卢图普Schlutup = BSchlutup施卢图普
	f.施卢图普Schlutup.SetParent(f)
	
	f.斯塔里加德Starigard = BStarigard斯塔里加德
	f.斯塔里加德Starigard.SetParent(f)
	
	f.特拉弗明德Travemunde = BTravemunde特拉弗明德
	f.特拉弗明德Travemunde.SetParent(f)
	
	f.韦斯勒Weslo = BWeslo韦斯勒
	f.韦斯勒Weslo.SetParent(f)
	
	f.伍尔夫斯多夫Wulfsdorf = BWulfsdorf伍尔夫斯多夫
	f.伍尔夫斯多夫Wulfsdorf.SetParent(f)
	
}
