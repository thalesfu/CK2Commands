package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KumtagCounty interface {
    feud.County
    BChiting赤亭() 	feud.Barony
    BDajin大金() 	feud.Barony
    BGuojia果佳() 	feud.Barony
    BKumtag库姆塔格() 	feud.Barony
    BLuohu罗护() 	feud.Barony
    BXiaoguoyu小锅峪() 	feud.Barony
    BXiaojian小尖() 	feud.Barony
}

type 库姆塔格KumtagCounty struct {
	feud.BaseCounty
	赤亭Chiting 	feud.Barony
	大金Dajin 	feud.Barony
	果佳Guojia 	feud.Barony
	库姆塔格Kumtag 	feud.Barony
	罗护Luohu 	feud.Barony
	小锅峪Xiaoguoyu 	feud.Barony
	小尖Xiaojian 	feud.Barony
}

func (c *库姆塔格KumtagCounty) BChiting赤亭() feud.Barony {
	return c.赤亭Chiting
}
    
func (c *库姆塔格KumtagCounty) BDajin大金() feud.Barony {
	return c.大金Dajin
}
    
func (c *库姆塔格KumtagCounty) BGuojia果佳() feud.Barony {
	return c.果佳Guojia
}
    
func (c *库姆塔格KumtagCounty) BKumtag库姆塔格() feud.Barony {
	return c.库姆塔格Kumtag
}
    
func (c *库姆塔格KumtagCounty) BLuohu罗护() feud.Barony {
	return c.罗护Luohu
}
    
func (c *库姆塔格KumtagCounty) BXiaoguoyu小锅峪() feud.Barony {
	return c.小锅峪Xiaoguoyu
}
    
func (c *库姆塔格KumtagCounty) BXiaojian小尖() feud.Barony {
	return c.小尖Xiaojian
}
    
var CKumtag库姆塔格 KumtagCounty = &库姆塔格KumtagCounty{}

func init() {
	f := CKumtag库姆塔格.(*库姆塔格KumtagCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1518",
		Title:     "kumtag",
		TitleName: "库姆塔格",
		TitleCode: "c_kumtag",
		Baronies:  map[string]feud.Barony{},
	}

	f.赤亭Chiting = BChiting赤亭
	f.赤亭Chiting.SetParent(f)
	
	f.大金Dajin = BDajin大金
	f.大金Dajin.SetParent(f)
	
	f.果佳Guojia = BGuojia果佳
	f.果佳Guojia.SetParent(f)
	
	f.库姆塔格Kumtag = BKumtag库姆塔格
	f.库姆塔格Kumtag.SetParent(f)
	
	f.罗护Luohu = BLuohu罗护
	f.罗护Luohu.SetParent(f)
	
	f.小锅峪Xiaoguoyu = BXiaoguoyu小锅峪
	f.小锅峪Xiaoguoyu.SetParent(f)
	
	f.小尖Xiaojian = BXiaojian小尖
	f.小尖Xiaojian.SetParent(f)
	
}
