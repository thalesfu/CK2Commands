package westfriesland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WestfrieslandCounty interface {
    feud.County
    BAlkmaar阿尔克马尔() 	feud.Barony
    BEgmond埃赫蒙德() 	feud.Barony
    BEnkhuizen恩克黑曾() 	feud.Barony
    BHeerhugowaard海尔许霍瓦德() 	feud.Barony
    BHoorn霍伦() 	feud.Barony
    BMedemblik梅登布利克() 	feud.Barony
    BSchagen斯哈亨() 	feud.Barony
    BTexel特塞尔() 	feud.Barony
}

type 西弗里斯兰WestfrieslandCounty struct {
	feud.BaseCounty
	阿尔克马尔Alkmaar 	feud.Barony
	埃赫蒙德Egmond 	feud.Barony
	恩克黑曾Enkhuizen 	feud.Barony
	海尔许霍瓦德Heerhugowaard 	feud.Barony
	霍伦Hoorn 	feud.Barony
	梅登布利克Medemblik 	feud.Barony
	斯哈亨Schagen 	feud.Barony
	特塞尔Texel 	feud.Barony
}

func (c *西弗里斯兰WestfrieslandCounty) BAlkmaar阿尔克马尔() feud.Barony {
	return c.阿尔克马尔Alkmaar
}
    
func (c *西弗里斯兰WestfrieslandCounty) BEgmond埃赫蒙德() feud.Barony {
	return c.埃赫蒙德Egmond
}
    
func (c *西弗里斯兰WestfrieslandCounty) BEnkhuizen恩克黑曾() feud.Barony {
	return c.恩克黑曾Enkhuizen
}
    
func (c *西弗里斯兰WestfrieslandCounty) BHeerhugowaard海尔许霍瓦德() feud.Barony {
	return c.海尔许霍瓦德Heerhugowaard
}
    
func (c *西弗里斯兰WestfrieslandCounty) BHoorn霍伦() feud.Barony {
	return c.霍伦Hoorn
}
    
func (c *西弗里斯兰WestfrieslandCounty) BMedemblik梅登布利克() feud.Barony {
	return c.梅登布利克Medemblik
}
    
func (c *西弗里斯兰WestfrieslandCounty) BSchagen斯哈亨() feud.Barony {
	return c.斯哈亨Schagen
}
    
func (c *西弗里斯兰WestfrieslandCounty) BTexel特塞尔() feud.Barony {
	return c.特塞尔Texel
}
    
var CWestfriesland西弗里斯兰 WestfrieslandCounty = &西弗里斯兰WestfrieslandCounty{}

func init() {
	f := CWestfriesland西弗里斯兰.(*西弗里斯兰WestfrieslandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "81",
		Title:     "westfriesland",
		TitleName: "西弗里斯兰",
		TitleCode: "c_westfriesland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔克马尔Alkmaar = BAlkmaar阿尔克马尔
	f.阿尔克马尔Alkmaar.SetParent(f)
	
	f.埃赫蒙德Egmond = BEgmond埃赫蒙德
	f.埃赫蒙德Egmond.SetParent(f)
	
	f.恩克黑曾Enkhuizen = BEnkhuizen恩克黑曾
	f.恩克黑曾Enkhuizen.SetParent(f)
	
	f.海尔许霍瓦德Heerhugowaard = BHeerhugowaard海尔许霍瓦德
	f.海尔许霍瓦德Heerhugowaard.SetParent(f)
	
	f.霍伦Hoorn = BHoorn霍伦
	f.霍伦Hoorn.SetParent(f)
	
	f.梅登布利克Medemblik = BMedemblik梅登布利克
	f.梅登布利克Medemblik.SetParent(f)
	
	f.斯哈亨Schagen = BSchagen斯哈亨
	f.斯哈亨Schagen.SetParent(f)
	
	f.特塞尔Texel = BTexel特塞尔
	f.特塞尔Texel.SetParent(f)
	
}
