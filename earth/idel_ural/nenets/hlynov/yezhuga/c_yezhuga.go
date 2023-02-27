package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YezhugaCounty interface {
    feud.County
    BPilisa皮利萨() 	feud.Barony
    BPinega皮涅加() 	feud.Barony
    BShirokoye希罗科耶() 	feud.Barony
    BVyva维瓦() 	feud.Barony
    BYeyuga叶尤加() 	feud.Barony
    BYezhuga叶茹加() 	feud.Barony
    BYula尤拉() 	feud.Barony
}

type 皮涅加YezhugaCounty struct {
	feud.BaseCounty
	皮利萨Pilisa 	feud.Barony
	皮涅加Pinega 	feud.Barony
	希罗科耶Shirokoye 	feud.Barony
	维瓦Vyva 	feud.Barony
	叶尤加Yeyuga 	feud.Barony
	叶茹加Yezhuga 	feud.Barony
	尤拉Yula 	feud.Barony
}

func (c *皮涅加YezhugaCounty) BPilisa皮利萨() feud.Barony {
	return c.皮利萨Pilisa
}
    
func (c *皮涅加YezhugaCounty) BPinega皮涅加() feud.Barony {
	return c.皮涅加Pinega
}
    
func (c *皮涅加YezhugaCounty) BShirokoye希罗科耶() feud.Barony {
	return c.希罗科耶Shirokoye
}
    
func (c *皮涅加YezhugaCounty) BVyva维瓦() feud.Barony {
	return c.维瓦Vyva
}
    
func (c *皮涅加YezhugaCounty) BYeyuga叶尤加() feud.Barony {
	return c.叶尤加Yeyuga
}
    
func (c *皮涅加YezhugaCounty) BYezhuga叶茹加() feud.Barony {
	return c.叶茹加Yezhuga
}
    
func (c *皮涅加YezhugaCounty) BYula尤拉() feud.Barony {
	return c.尤拉Yula
}
    
var CYezhuga皮涅加 YezhugaCounty = &皮涅加YezhugaCounty{}

func init() {
	f := CYezhuga皮涅加.(*皮涅加YezhugaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1824",
		Title:     "yezhuga",
		TitleName: "皮涅加",
		TitleCode: "c_yezhuga",
		Baronies:  map[string]feud.Barony{},
	}

	f.皮利萨Pilisa = BPilisa皮利萨
	f.皮利萨Pilisa.SetParent(f)
	
	f.皮涅加Pinega = BPinega皮涅加
	f.皮涅加Pinega.SetParent(f)
	
	f.希罗科耶Shirokoye = BShirokoye希罗科耶
	f.希罗科耶Shirokoye.SetParent(f)
	
	f.维瓦Vyva = BVyva维瓦
	f.维瓦Vyva.SetParent(f)
	
	f.叶尤加Yeyuga = BYeyuga叶尤加
	f.叶尤加Yeyuga.SetParent(f)
	
	f.叶茹加Yezhuga = BYezhuga叶茹加
	f.叶茹加Yezhuga.SetParent(f)
	
	f.尤拉Yula = BYula尤拉
	f.尤拉Yula.SetParent(f)
	
}
