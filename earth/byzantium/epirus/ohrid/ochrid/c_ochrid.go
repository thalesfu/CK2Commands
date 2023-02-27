package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OchridCounty interface {
    feud.County
    BDebar德巴尔() 	feud.Barony
    BKicevo基切沃() 	feud.Barony
    BKrusevo克鲁舍沃() 	feud.Barony
    BOhrid奥赫里德() 	feud.Barony
    BStruga斯特鲁加() 	feud.Barony
    BSvetigrad斯韦蒂格拉德() 	feud.Barony
    BTomot土默特() 	feud.Barony
}

type 奥赫里德OchridCounty struct {
	feud.BaseCounty
	德巴尔Debar 	feud.Barony
	基切沃Kicevo 	feud.Barony
	克鲁舍沃Krusevo 	feud.Barony
	奥赫里德Ohrid 	feud.Barony
	斯特鲁加Struga 	feud.Barony
	斯韦蒂格拉德Svetigrad 	feud.Barony
	土默特Tomot 	feud.Barony
}

func (c *奥赫里德OchridCounty) BDebar德巴尔() feud.Barony {
	return c.德巴尔Debar
}
    
func (c *奥赫里德OchridCounty) BKicevo基切沃() feud.Barony {
	return c.基切沃Kicevo
}
    
func (c *奥赫里德OchridCounty) BKrusevo克鲁舍沃() feud.Barony {
	return c.克鲁舍沃Krusevo
}
    
func (c *奥赫里德OchridCounty) BOhrid奥赫里德() feud.Barony {
	return c.奥赫里德Ohrid
}
    
func (c *奥赫里德OchridCounty) BStruga斯特鲁加() feud.Barony {
	return c.斯特鲁加Struga
}
    
func (c *奥赫里德OchridCounty) BSvetigrad斯韦蒂格拉德() feud.Barony {
	return c.斯韦蒂格拉德Svetigrad
}
    
func (c *奥赫里德OchridCounty) BTomot土默特() feud.Barony {
	return c.土默特Tomot
}
    
var COchrid奥赫里德 OchridCounty = &奥赫里德OchridCounty{}

func init() {
	f := COchrid奥赫里德.(*奥赫里德OchridCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "471",
		Title:     "ochrid",
		TitleName: "奥赫里德",
		TitleCode: "c_ochrid",
		Baronies:  map[string]feud.Barony{},
	}

	f.德巴尔Debar = BDebar德巴尔
	f.德巴尔Debar.SetParent(f)
	
	f.基切沃Kicevo = BKicevo基切沃
	f.基切沃Kicevo.SetParent(f)
	
	f.克鲁舍沃Krusevo = BKrusevo克鲁舍沃
	f.克鲁舍沃Krusevo.SetParent(f)
	
	f.奥赫里德Ohrid = BOhrid奥赫里德
	f.奥赫里德Ohrid.SetParent(f)
	
	f.斯特鲁加Struga = BStruga斯特鲁加
	f.斯特鲁加Struga.SetParent(f)
	
	f.斯韦蒂格拉德Svetigrad = BSvetigrad斯韦蒂格拉德
	f.斯韦蒂格拉德Svetigrad.SetParent(f)
	
	f.土默特Tomot = BTomot土默特
	f.土默特Tomot.SetParent(f)
	
}
