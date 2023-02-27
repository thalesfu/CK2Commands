package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type East_kareliaCounty interface {
    feud.County
    BMasl玛斯勒() 	feud.Barony
    BOndo翁多() 	feud.Barony
    BSegozero谢戈泽罗() 	feud.Barony
    BSekee塞凯() 	feud.Barony
    BSoroka索罗卡() 	feud.Barony
    BVyg维格() 	feud.Barony
    BYelm耶尔姆() 	feud.Barony
}

type 波莫里耶East_kareliaCounty struct {
	feud.BaseCounty
	玛斯勒Masl 	feud.Barony
	翁多Ondo 	feud.Barony
	谢戈泽罗Segozero 	feud.Barony
	塞凯Sekee 	feud.Barony
	索罗卡Soroka 	feud.Barony
	维格Vyg 	feud.Barony
	耶尔姆Yelm 	feud.Barony
}

func (c *波莫里耶East_kareliaCounty) BMasl玛斯勒() feud.Barony {
	return c.玛斯勒Masl
}
    
func (c *波莫里耶East_kareliaCounty) BOndo翁多() feud.Barony {
	return c.翁多Ondo
}
    
func (c *波莫里耶East_kareliaCounty) BSegozero谢戈泽罗() feud.Barony {
	return c.谢戈泽罗Segozero
}
    
func (c *波莫里耶East_kareliaCounty) BSekee塞凯() feud.Barony {
	return c.塞凯Sekee
}
    
func (c *波莫里耶East_kareliaCounty) BSoroka索罗卡() feud.Barony {
	return c.索罗卡Soroka
}
    
func (c *波莫里耶East_kareliaCounty) BVyg维格() feud.Barony {
	return c.维格Vyg
}
    
func (c *波莫里耶East_kareliaCounty) BYelm耶尔姆() feud.Barony {
	return c.耶尔姆Yelm
}
    
var CEast_karelia波莫里耶 East_kareliaCounty = &波莫里耶East_kareliaCounty{}

func init() {
	f := CEast_karelia波莫里耶.(*波莫里耶East_kareliaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1603",
		Title:     "east_karelia",
		TitleName: "波莫里耶",
		TitleCode: "c_east_karelia",
		Baronies:  map[string]feud.Barony{},
	}

	f.玛斯勒Masl = BMasl玛斯勒
	f.玛斯勒Masl.SetParent(f)
	
	f.翁多Ondo = BOndo翁多
	f.翁多Ondo.SetParent(f)
	
	f.谢戈泽罗Segozero = BSegozero谢戈泽罗
	f.谢戈泽罗Segozero.SetParent(f)
	
	f.塞凯Sekee = BSekee塞凯
	f.塞凯Sekee.SetParent(f)
	
	f.索罗卡Soroka = BSoroka索罗卡
	f.索罗卡Soroka.SetParent(f)
	
	f.维格Vyg = BVyg维格
	f.维格Vyg.SetParent(f)
	
	f.耶尔姆Yelm = BYelm耶尔姆
	f.耶尔姆Yelm.SetParent(f)
	
}
