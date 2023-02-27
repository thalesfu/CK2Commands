package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NedongCounty interface {
    feud.County
    BNedong乃东() 	feud.Barony
    BPhagmodru伯木古鲁() 	feud.Barony
    BSangri桑日() 	feud.Barony
    BTradruk昌珠() 	feud.Barony
    BYungbulakang雍布拉康() 	feud.Barony
    BZengqi增期() 	feud.Barony
    BZetang孜塘() 	feud.Barony
}

type 乃东NedongCounty struct {
	feud.BaseCounty
	乃东Nedong 	feud.Barony
	伯木古鲁Phagmodru 	feud.Barony
	桑日Sangri 	feud.Barony
	昌珠Tradruk 	feud.Barony
	雍布拉康Yungbulakang 	feud.Barony
	增期Zengqi 	feud.Barony
	孜塘Zetang 	feud.Barony
}

func (c *乃东NedongCounty) BNedong乃东() feud.Barony {
	return c.乃东Nedong
}
    
func (c *乃东NedongCounty) BPhagmodru伯木古鲁() feud.Barony {
	return c.伯木古鲁Phagmodru
}
    
func (c *乃东NedongCounty) BSangri桑日() feud.Barony {
	return c.桑日Sangri
}
    
func (c *乃东NedongCounty) BTradruk昌珠() feud.Barony {
	return c.昌珠Tradruk
}
    
func (c *乃东NedongCounty) BYungbulakang雍布拉康() feud.Barony {
	return c.雍布拉康Yungbulakang
}
    
func (c *乃东NedongCounty) BZengqi增期() feud.Barony {
	return c.增期Zengqi
}
    
func (c *乃东NedongCounty) BZetang孜塘() feud.Barony {
	return c.孜塘Zetang
}
    
var CNedong乃东 NedongCounty = &乃东NedongCounty{}

func init() {
	f := CNedong乃东.(*乃东NedongCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1579",
		Title:     "nedong",
		TitleName: "乃东",
		TitleCode: "c_nedong",
		Baronies:  map[string]feud.Barony{},
	}

	f.乃东Nedong = BNedong乃东
	f.乃东Nedong.SetParent(f)
	
	f.伯木古鲁Phagmodru = BPhagmodru伯木古鲁
	f.伯木古鲁Phagmodru.SetParent(f)
	
	f.桑日Sangri = BSangri桑日
	f.桑日Sangri.SetParent(f)
	
	f.昌珠Tradruk = BTradruk昌珠
	f.昌珠Tradruk.SetParent(f)
	
	f.雍布拉康Yungbulakang = BYungbulakang雍布拉康
	f.雍布拉康Yungbulakang.SetParent(f)
	
	f.增期Zengqi = BZengqi增期
	f.增期Zengqi.SetParent(f)
	
	f.孜塘Zetang = BZetang孜塘
	f.孜塘Zetang.SetParent(f)
	
}
