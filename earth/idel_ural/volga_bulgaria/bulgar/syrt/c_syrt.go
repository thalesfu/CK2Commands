package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyrtCounty interface {
    feud.County
    BKinel基涅利() 	feud.Barony
    BOsinki奥辛基() 	feud.Barony
    BSamara萨马拉() 	feud.Barony
    BSarbay萨尔拜() 	feud.Barony
    BShungut顺古特() 	feud.Barony
    BSyzran塞兹兰() 	feud.Barony
    BTashia塔希亚() 	feud.Barony
}

type 瑟尔特SyrtCounty struct {
	feud.BaseCounty
	基涅利Kinel 	feud.Barony
	奥辛基Osinki 	feud.Barony
	萨马拉Samara 	feud.Barony
	萨尔拜Sarbay 	feud.Barony
	顺古特Shungut 	feud.Barony
	塞兹兰Syzran 	feud.Barony
	塔希亚Tashia 	feud.Barony
}

func (c *瑟尔特SyrtCounty) BKinel基涅利() feud.Barony {
	return c.基涅利Kinel
}
    
func (c *瑟尔特SyrtCounty) BOsinki奥辛基() feud.Barony {
	return c.奥辛基Osinki
}
    
func (c *瑟尔特SyrtCounty) BSamara萨马拉() feud.Barony {
	return c.萨马拉Samara
}
    
func (c *瑟尔特SyrtCounty) BSarbay萨尔拜() feud.Barony {
	return c.萨尔拜Sarbay
}
    
func (c *瑟尔特SyrtCounty) BShungut顺古特() feud.Barony {
	return c.顺古特Shungut
}
    
func (c *瑟尔特SyrtCounty) BSyzran塞兹兰() feud.Barony {
	return c.塞兹兰Syzran
}
    
func (c *瑟尔特SyrtCounty) BTashia塔希亚() feud.Barony {
	return c.塔希亚Tashia
}
    
var CSyrt瑟尔特 SyrtCounty = &瑟尔特SyrtCounty{}

func init() {
	f := CSyrt瑟尔特.(*瑟尔特SyrtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "609",
		Title:     "syrt",
		TitleName: "瑟尔特",
		TitleCode: "c_syrt",
		Baronies:  map[string]feud.Barony{},
	}

	f.基涅利Kinel = BKinel基涅利
	f.基涅利Kinel.SetParent(f)
	
	f.奥辛基Osinki = BOsinki奥辛基
	f.奥辛基Osinki.SetParent(f)
	
	f.萨马拉Samara = BSamara萨马拉
	f.萨马拉Samara.SetParent(f)
	
	f.萨尔拜Sarbay = BSarbay萨尔拜
	f.萨尔拜Sarbay.SetParent(f)
	
	f.顺古特Shungut = BShungut顺古特
	f.顺古特Shungut.SetParent(f)
	
	f.塞兹兰Syzran = BSyzran塞兹兰
	f.塞兹兰Syzran.SetParent(f)
	
	f.塔希亚Tashia = BTashia塔希亚
	f.塔希亚Tashia.SetParent(f)
	
}
