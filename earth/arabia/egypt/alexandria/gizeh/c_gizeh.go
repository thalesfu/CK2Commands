package gizeh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GizehCounty interface {
    feud.County
    BAbughorob阿布吉拉伯() 	feud.Barony
    BAburowash阿布鲁韦斯() 	feud.Barony
    BAbusir阿布西尔() 	feud.Barony
    BDashur代赫舒尔() 	feud.Barony
    BEllisht利什特() 	feud.Barony
    BGizeh吉萨() 	feud.Barony
    BSaqqara萨卡拉() 	feud.Barony
    BZawyetalayran扎维耶奥亚因() 	feud.Barony
}

type 法尤姆GizehCounty struct {
	feud.BaseCounty
	阿布吉拉伯Abughorob 	feud.Barony
	阿布鲁韦斯Aburowash 	feud.Barony
	阿布西尔Abusir 	feud.Barony
	代赫舒尔Dashur 	feud.Barony
	利什特Ellisht 	feud.Barony
	吉萨Gizeh 	feud.Barony
	萨卡拉Saqqara 	feud.Barony
	扎维耶奥亚因Zawyetalayran 	feud.Barony
}

func (c *法尤姆GizehCounty) BAbughorob阿布吉拉伯() feud.Barony {
	return c.阿布吉拉伯Abughorob
}
    
func (c *法尤姆GizehCounty) BAburowash阿布鲁韦斯() feud.Barony {
	return c.阿布鲁韦斯Aburowash
}
    
func (c *法尤姆GizehCounty) BAbusir阿布西尔() feud.Barony {
	return c.阿布西尔Abusir
}
    
func (c *法尤姆GizehCounty) BDashur代赫舒尔() feud.Barony {
	return c.代赫舒尔Dashur
}
    
func (c *法尤姆GizehCounty) BEllisht利什特() feud.Barony {
	return c.利什特Ellisht
}
    
func (c *法尤姆GizehCounty) BGizeh吉萨() feud.Barony {
	return c.吉萨Gizeh
}
    
func (c *法尤姆GizehCounty) BSaqqara萨卡拉() feud.Barony {
	return c.萨卡拉Saqqara
}
    
func (c *法尤姆GizehCounty) BZawyetalayran扎维耶奥亚因() feud.Barony {
	return c.扎维耶奥亚因Zawyetalayran
}
    
var CGizeh法尤姆 GizehCounty = &法尤姆GizehCounty{}

func init() {
	f := CGizeh法尤姆.(*法尤姆GizehCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "800",
		Title:     "gizeh",
		TitleName: "法尤姆",
		TitleCode: "c_gizeh",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布吉拉伯Abughorob = BAbughorob阿布吉拉伯
	f.阿布吉拉伯Abughorob.SetParent(f)
	
	f.阿布鲁韦斯Aburowash = BAburowash阿布鲁韦斯
	f.阿布鲁韦斯Aburowash.SetParent(f)
	
	f.阿布西尔Abusir = BAbusir阿布西尔
	f.阿布西尔Abusir.SetParent(f)
	
	f.代赫舒尔Dashur = BDashur代赫舒尔
	f.代赫舒尔Dashur.SetParent(f)
	
	f.利什特Ellisht = BEllisht利什特
	f.利什特Ellisht.SetParent(f)
	
	f.吉萨Gizeh = BGizeh吉萨
	f.吉萨Gizeh.SetParent(f)
	
	f.萨卡拉Saqqara = BSaqqara萨卡拉
	f.萨卡拉Saqqara.SetParent(f)
	
	f.扎维耶奥亚因Zawyetalayran = BZawyetalayran扎维耶奥亚因
	f.扎维耶奥亚因Zawyetalayran.SetParent(f)
	
}
