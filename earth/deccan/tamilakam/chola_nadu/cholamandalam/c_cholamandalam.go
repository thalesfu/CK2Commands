package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CholamandalamCounty interface {
    feud.County
    BGangaikondacolapuram恒伽军荼朱罗补蓝() 	feud.Barony
    BKannanur甘那罗() 	feud.Barony
    BKumbhakarna宫婆羯罗拏() 	feud.Barony
    BNagapattinam那伽钵亶那() 	feud.Barony
    BSrirangam斯里愣甘() 	feud.Barony
    BTanjavur檀伽瓦() 	feud.Barony
    BTiruccirapalli底哩尸罗钵利() 	feud.Barony
    BUraiyur乌莱尤尔() 	feud.Barony
}

type 朱罗曼荼罗CholamandalamCounty struct {
	feud.BaseCounty
	恒伽军荼朱罗补蓝Gangaikondacolapuram 	feud.Barony
	甘那罗Kannanur 	feud.Barony
	宫婆羯罗拏Kumbhakarna 	feud.Barony
	那伽钵亶那Nagapattinam 	feud.Barony
	斯里愣甘Srirangam 	feud.Barony
	檀伽瓦Tanjavur 	feud.Barony
	底哩尸罗钵利Tiruccirapalli 	feud.Barony
	乌莱尤尔Uraiyur 	feud.Barony
}

func (c *朱罗曼荼罗CholamandalamCounty) BGangaikondacolapuram恒伽军荼朱罗补蓝() feud.Barony {
	return c.恒伽军荼朱罗补蓝Gangaikondacolapuram
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BKannanur甘那罗() feud.Barony {
	return c.甘那罗Kannanur
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BKumbhakarna宫婆羯罗拏() feud.Barony {
	return c.宫婆羯罗拏Kumbhakarna
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BNagapattinam那伽钵亶那() feud.Barony {
	return c.那伽钵亶那Nagapattinam
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BSrirangam斯里愣甘() feud.Barony {
	return c.斯里愣甘Srirangam
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BTanjavur檀伽瓦() feud.Barony {
	return c.檀伽瓦Tanjavur
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BTiruccirapalli底哩尸罗钵利() feud.Barony {
	return c.底哩尸罗钵利Tiruccirapalli
}
    
func (c *朱罗曼荼罗CholamandalamCounty) BUraiyur乌莱尤尔() feud.Barony {
	return c.乌莱尤尔Uraiyur
}
    
var CCholamandalam朱罗曼荼罗 CholamandalamCounty = &朱罗曼荼罗CholamandalamCounty{}

func init() {
	f := CCholamandalam朱罗曼荼罗.(*朱罗曼荼罗CholamandalamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1115",
		Title:     "cholamandalam",
		TitleName: "朱罗曼荼罗",
		TitleCode: "c_cholamandalam",
		Baronies:  map[string]feud.Barony{},
	}

	f.恒伽军荼朱罗补蓝Gangaikondacolapuram = BGangaikondacolapuram恒伽军荼朱罗补蓝
	f.恒伽军荼朱罗补蓝Gangaikondacolapuram.SetParent(f)
	
	f.甘那罗Kannanur = BKannanur甘那罗
	f.甘那罗Kannanur.SetParent(f)
	
	f.宫婆羯罗拏Kumbhakarna = BKumbhakarna宫婆羯罗拏
	f.宫婆羯罗拏Kumbhakarna.SetParent(f)
	
	f.那伽钵亶那Nagapattinam = BNagapattinam那伽钵亶那
	f.那伽钵亶那Nagapattinam.SetParent(f)
	
	f.斯里愣甘Srirangam = BSrirangam斯里愣甘
	f.斯里愣甘Srirangam.SetParent(f)
	
	f.檀伽瓦Tanjavur = BTanjavur檀伽瓦
	f.檀伽瓦Tanjavur.SetParent(f)
	
	f.底哩尸罗钵利Tiruccirapalli = BTiruccirapalli底哩尸罗钵利
	f.底哩尸罗钵利Tiruccirapalli.SetParent(f)
	
	f.乌莱尤尔Uraiyur = BUraiyur乌莱尤尔
	f.乌莱尤尔Uraiyur.SetParent(f)
	
}
