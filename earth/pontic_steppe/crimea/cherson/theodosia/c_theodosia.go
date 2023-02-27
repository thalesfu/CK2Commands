package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TheodosiaCounty interface {
    feud.County
    BCaffa卡法() 	feud.Barony
    BCaulita考里塔() 	feud.Barony
    BFunan弗南() 	feud.Barony
    BKimmerikon基梅里孔() 	feud.Barony
    BLusta路斯塔() 	feud.Barony
    BOlyva奥里瓦() 	feud.Barony
    BSoldaia索尔得亚() 	feud.Barony
    BTheodosia狄奥多西亚() 	feud.Barony
}

type 狄奥多西亚TheodosiaCounty struct {
	feud.BaseCounty
	卡法Caffa 	feud.Barony
	考里塔Caulita 	feud.Barony
	弗南Funan 	feud.Barony
	基梅里孔Kimmerikon 	feud.Barony
	路斯塔Lusta 	feud.Barony
	奥里瓦Olyva 	feud.Barony
	索尔得亚Soldaia 	feud.Barony
	狄奥多西亚Theodosia 	feud.Barony
}

func (c *狄奥多西亚TheodosiaCounty) BCaffa卡法() feud.Barony {
	return c.卡法Caffa
}
    
func (c *狄奥多西亚TheodosiaCounty) BCaulita考里塔() feud.Barony {
	return c.考里塔Caulita
}
    
func (c *狄奥多西亚TheodosiaCounty) BFunan弗南() feud.Barony {
	return c.弗南Funan
}
    
func (c *狄奥多西亚TheodosiaCounty) BKimmerikon基梅里孔() feud.Barony {
	return c.基梅里孔Kimmerikon
}
    
func (c *狄奥多西亚TheodosiaCounty) BLusta路斯塔() feud.Barony {
	return c.路斯塔Lusta
}
    
func (c *狄奥多西亚TheodosiaCounty) BOlyva奥里瓦() feud.Barony {
	return c.奥里瓦Olyva
}
    
func (c *狄奥多西亚TheodosiaCounty) BSoldaia索尔得亚() feud.Barony {
	return c.索尔得亚Soldaia
}
    
func (c *狄奥多西亚TheodosiaCounty) BTheodosia狄奥多西亚() feud.Barony {
	return c.狄奥多西亚Theodosia
}
    
var CTheodosia狄奥多西亚 TheodosiaCounty = &狄奥多西亚TheodosiaCounty{}

func init() {
	f := CTheodosia狄奥多西亚.(*狄奥多西亚TheodosiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "561",
		Title:     "theodosia",
		TitleName: "狄奥多西亚",
		TitleCode: "c_theodosia",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡法Caffa = BCaffa卡法
	f.卡法Caffa.SetParent(f)
	
	f.考里塔Caulita = BCaulita考里塔
	f.考里塔Caulita.SetParent(f)
	
	f.弗南Funan = BFunan弗南
	f.弗南Funan.SetParent(f)
	
	f.基梅里孔Kimmerikon = BKimmerikon基梅里孔
	f.基梅里孔Kimmerikon.SetParent(f)
	
	f.路斯塔Lusta = BLusta路斯塔
	f.路斯塔Lusta.SetParent(f)
	
	f.奥里瓦Olyva = BOlyva奥里瓦
	f.奥里瓦Olyva.SetParent(f)
	
	f.索尔得亚Soldaia = BSoldaia索尔得亚
	f.索尔得亚Soldaia.SetParent(f)
	
	f.狄奥多西亚Theodosia = BTheodosia狄奥多西亚
	f.狄奥多西亚Theodosia.SetParent(f)
	
}
