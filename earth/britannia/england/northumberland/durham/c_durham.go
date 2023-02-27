package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DurhamCounty interface {
    feud.County
    BAuckland奥克兰() 	feud.Barony
    BDurham达勒姆() 	feud.Barony
    BGateshead盖茨黑德() 	feud.Barony
    BHartlepool哈特尔浦() 	feud.Barony
    BJarrow贾罗() 	feud.Barony
    BRaby雷比() 	feud.Barony
    BSt_cuthbert圣卡斯伯特() 	feud.Barony
}

type 达勒姆DurhamCounty struct {
	feud.BaseCounty
	奥克兰Auckland 	feud.Barony
	达勒姆Durham 	feud.Barony
	盖茨黑德Gateshead 	feud.Barony
	哈特尔浦Hartlepool 	feud.Barony
	贾罗Jarrow 	feud.Barony
	雷比Raby 	feud.Barony
	圣卡斯伯特St_cuthbert 	feud.Barony
}

func (c *达勒姆DurhamCounty) BAuckland奥克兰() feud.Barony {
	return c.奥克兰Auckland
}
    
func (c *达勒姆DurhamCounty) BDurham达勒姆() feud.Barony {
	return c.达勒姆Durham
}
    
func (c *达勒姆DurhamCounty) BGateshead盖茨黑德() feud.Barony {
	return c.盖茨黑德Gateshead
}
    
func (c *达勒姆DurhamCounty) BHartlepool哈特尔浦() feud.Barony {
	return c.哈特尔浦Hartlepool
}
    
func (c *达勒姆DurhamCounty) BJarrow贾罗() feud.Barony {
	return c.贾罗Jarrow
}
    
func (c *达勒姆DurhamCounty) BRaby雷比() feud.Barony {
	return c.雷比Raby
}
    
func (c *达勒姆DurhamCounty) BSt_cuthbert圣卡斯伯特() feud.Barony {
	return c.圣卡斯伯特St_cuthbert
}
    
var CDurham达勒姆 DurhamCounty = &达勒姆DurhamCounty{}

func init() {
	f := CDurham达勒姆.(*达勒姆DurhamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "56",
		Title:     "durham",
		TitleName: "达勒姆",
		TitleCode: "c_durham",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥克兰Auckland = BAuckland奥克兰
	f.奥克兰Auckland.SetParent(f)
	
	f.达勒姆Durham = BDurham达勒姆
	f.达勒姆Durham.SetParent(f)
	
	f.盖茨黑德Gateshead = BGateshead盖茨黑德
	f.盖茨黑德Gateshead.SetParent(f)
	
	f.哈特尔浦Hartlepool = BHartlepool哈特尔浦
	f.哈特尔浦Hartlepool.SetParent(f)
	
	f.贾罗Jarrow = BJarrow贾罗
	f.贾罗Jarrow.SetParent(f)
	
	f.雷比Raby = BRaby雷比
	f.雷比Raby.SetParent(f)
	
	f.圣卡斯伯特St_cuthbert = BSt_cuthbert圣卡斯伯特
	f.圣卡斯伯特St_cuthbert.SetParent(f)
	
}
