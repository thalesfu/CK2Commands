package adana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AdanaCounty interface {
    feud.County
    BAdana阿达纳() 	feud.Barony
    BAnazarba阿那扎巴() 	feud.Barony
    BLajazzo拉扎佐() 	feud.Barony
    BMamistra曼密西() 	feud.Barony
    BMopsuestia摩普绥提亚() 	feud.Barony
    BSis斯丝() 	feud.Barony
    BTrazak特拉扎克() 	feud.Barony
    BVahka费凯() 	feud.Barony
}

type 阿达纳AdanaCounty struct {
	feud.BaseCounty
	阿达纳Adana 	feud.Barony
	阿那扎巴Anazarba 	feud.Barony
	拉扎佐Lajazzo 	feud.Barony
	曼密西Mamistra 	feud.Barony
	摩普绥提亚Mopsuestia 	feud.Barony
	斯丝Sis 	feud.Barony
	特拉扎克Trazak 	feud.Barony
	费凯Vahka 	feud.Barony
}

func (c *阿达纳AdanaCounty) BAdana阿达纳() feud.Barony {
	return c.阿达纳Adana
}
    
func (c *阿达纳AdanaCounty) BAnazarba阿那扎巴() feud.Barony {
	return c.阿那扎巴Anazarba
}
    
func (c *阿达纳AdanaCounty) BLajazzo拉扎佐() feud.Barony {
	return c.拉扎佐Lajazzo
}
    
func (c *阿达纳AdanaCounty) BMamistra曼密西() feud.Barony {
	return c.曼密西Mamistra
}
    
func (c *阿达纳AdanaCounty) BMopsuestia摩普绥提亚() feud.Barony {
	return c.摩普绥提亚Mopsuestia
}
    
func (c *阿达纳AdanaCounty) BSis斯丝() feud.Barony {
	return c.斯丝Sis
}
    
func (c *阿达纳AdanaCounty) BTrazak特拉扎克() feud.Barony {
	return c.特拉扎克Trazak
}
    
func (c *阿达纳AdanaCounty) BVahka费凯() feud.Barony {
	return c.费凯Vahka
}
    
var CAdana阿达纳 AdanaCounty = &阿达纳AdanaCounty{}

func init() {
	f := CAdana阿达纳.(*阿达纳AdanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "762",
		Title:     "adana",
		TitleName: "阿达纳",
		TitleCode: "c_adana",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿达纳Adana = BAdana阿达纳
	f.阿达纳Adana.SetParent(f)
	
	f.阿那扎巴Anazarba = BAnazarba阿那扎巴
	f.阿那扎巴Anazarba.SetParent(f)
	
	f.拉扎佐Lajazzo = BLajazzo拉扎佐
	f.拉扎佐Lajazzo.SetParent(f)
	
	f.曼密西Mamistra = BMamistra曼密西
	f.曼密西Mamistra.SetParent(f)
	
	f.摩普绥提亚Mopsuestia = BMopsuestia摩普绥提亚
	f.摩普绥提亚Mopsuestia.SetParent(f)
	
	f.斯丝Sis = BSis斯丝
	f.斯丝Sis.SetParent(f)
	
	f.特拉扎克Trazak = BTrazak特拉扎克
	f.特拉扎克Trazak.SetParent(f)
	
	f.费凯Vahka = BVahka费凯
	f.费凯Vahka.SetParent(f)
	
}
