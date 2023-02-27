package khokh_serkh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Khokh_serkhCounty interface {
    feud.County
    BDuruu多罗() 	feud.Barony
    BKhokh_serkh呼赫色尔赫() 	feud.Barony
    BKhovd_khokh_serkh科布多() 	feud.Barony
    BTal塔勒() 	feud.Barony
    BTolbo陶勒包() 	feud.Barony
    BTsast察斯特() 	feud.Barony
    BTugal图噶勒() 	feud.Barony
}

type 呼赫色尔赫Khokh_serkhCounty struct {
	feud.BaseCounty
	多罗Duruu 	feud.Barony
	呼赫色尔赫Khokh_serkh 	feud.Barony
	科布多Khovd_khokh_serkh 	feud.Barony
	塔勒Tal 	feud.Barony
	陶勒包Tolbo 	feud.Barony
	察斯特Tsast 	feud.Barony
	图噶勒Tugal 	feud.Barony
}

func (c *呼赫色尔赫Khokh_serkhCounty) BDuruu多罗() feud.Barony {
	return c.多罗Duruu
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BKhokh_serkh呼赫色尔赫() feud.Barony {
	return c.呼赫色尔赫Khokh_serkh
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BKhovd_khokh_serkh科布多() feud.Barony {
	return c.科布多Khovd_khokh_serkh
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BTal塔勒() feud.Barony {
	return c.塔勒Tal
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BTolbo陶勒包() feud.Barony {
	return c.陶勒包Tolbo
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BTsast察斯特() feud.Barony {
	return c.察斯特Tsast
}
    
func (c *呼赫色尔赫Khokh_serkhCounty) BTugal图噶勒() feud.Barony {
	return c.图噶勒Tugal
}
    
var CKhokh_serkh呼赫色尔赫 Khokh_serkhCounty = &呼赫色尔赫Khokh_serkhCounty{}

func init() {
	f := CKhokh_serkh呼赫色尔赫.(*呼赫色尔赫Khokh_serkhCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1909",
		Title:     "khokh_serkh",
		TitleName: "呼赫色尔赫",
		TitleCode: "c_khokh_serkh",
		Baronies:  map[string]feud.Barony{},
	}

	f.多罗Duruu = BDuruu多罗
	f.多罗Duruu.SetParent(f)
	
	f.呼赫色尔赫Khokh_serkh = BKhokh_serkh呼赫色尔赫
	f.呼赫色尔赫Khokh_serkh.SetParent(f)
	
	f.科布多Khovd_khokh_serkh = BKhovd_khokh_serkh科布多
	f.科布多Khovd_khokh_serkh.SetParent(f)
	
	f.塔勒Tal = BTal塔勒
	f.塔勒Tal.SetParent(f)
	
	f.陶勒包Tolbo = BTolbo陶勒包
	f.陶勒包Tolbo.SetParent(f)
	
	f.察斯特Tsast = BTsast察斯特
	f.察斯特Tsast.SetParent(f)
	
	f.图噶勒Tugal = BTugal图噶勒
	f.图噶勒Tugal.SetParent(f)
	
}
