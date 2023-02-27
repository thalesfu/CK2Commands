package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OpoleCounty interface {
    feud.County
    BKozle科伊莱() 	feud.Barony
    BLubliniec卢布利涅茨() 	feud.Barony
    BOpole奥波莱() 	feud.Barony
    BRaciborz拉齐布日() 	feud.Barony
    BRybnik里布尼克() 	feud.Barony
    BSiewerz谢维日() 	feud.Barony
    BStrzelce斯切尔采() 	feud.Barony
}

type 上西里西亚OpoleCounty struct {
	feud.BaseCounty
	科伊莱Kozle 	feud.Barony
	卢布利涅茨Lubliniec 	feud.Barony
	奥波莱Opole 	feud.Barony
	拉齐布日Raciborz 	feud.Barony
	里布尼克Rybnik 	feud.Barony
	谢维日Siewerz 	feud.Barony
	斯切尔采Strzelce 	feud.Barony
}

func (c *上西里西亚OpoleCounty) BKozle科伊莱() feud.Barony {
	return c.科伊莱Kozle
}
    
func (c *上西里西亚OpoleCounty) BLubliniec卢布利涅茨() feud.Barony {
	return c.卢布利涅茨Lubliniec
}
    
func (c *上西里西亚OpoleCounty) BOpole奥波莱() feud.Barony {
	return c.奥波莱Opole
}
    
func (c *上西里西亚OpoleCounty) BRaciborz拉齐布日() feud.Barony {
	return c.拉齐布日Raciborz
}
    
func (c *上西里西亚OpoleCounty) BRybnik里布尼克() feud.Barony {
	return c.里布尼克Rybnik
}
    
func (c *上西里西亚OpoleCounty) BSiewerz谢维日() feud.Barony {
	return c.谢维日Siewerz
}
    
func (c *上西里西亚OpoleCounty) BStrzelce斯切尔采() feud.Barony {
	return c.斯切尔采Strzelce
}
    
var COpole上西里西亚 OpoleCounty = &上西里西亚OpoleCounty{}

func init() {
	f := COpole上西里西亚.(*上西里西亚OpoleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "433",
		Title:     "opole",
		TitleName: "上西里西亚",
		TitleCode: "c_opole",
		Baronies:  map[string]feud.Barony{},
	}

	f.科伊莱Kozle = BKozle科伊莱
	f.科伊莱Kozle.SetParent(f)
	
	f.卢布利涅茨Lubliniec = BLubliniec卢布利涅茨
	f.卢布利涅茨Lubliniec.SetParent(f)
	
	f.奥波莱Opole = BOpole奥波莱
	f.奥波莱Opole.SetParent(f)
	
	f.拉齐布日Raciborz = BRaciborz拉齐布日
	f.拉齐布日Raciborz.SetParent(f)
	
	f.里布尼克Rybnik = BRybnik里布尼克
	f.里布尼克Rybnik.SetParent(f)
	
	f.谢维日Siewerz = BSiewerz谢维日
	f.谢维日Siewerz.SetParent(f)
	
	f.斯切尔采Strzelce = BStrzelce斯切尔采
	f.斯切尔采Strzelce.SetParent(f)
	
}
