package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DunhuangCounty interface {
    feud.County
    BDunhuang敦煌() 	feud.Barony
    BFangpan方盘() 	feud.Barony
    BHongliuwan红柳湾() 	feud.Barony
    BMogao莫高() 	feud.Barony
    BXuanquanzhi悬泉置() 	feud.Barony
    BYueyaquan月牙泉() 	feud.Barony
    BYumenquan玉门泉() 	feud.Barony
}

type 敦煌DunhuangCounty struct {
	feud.BaseCounty
	敦煌Dunhuang 	feud.Barony
	方盘Fangpan 	feud.Barony
	红柳湾Hongliuwan 	feud.Barony
	莫高Mogao 	feud.Barony
	悬泉置Xuanquanzhi 	feud.Barony
	月牙泉Yueyaquan 	feud.Barony
	玉门泉Yumenquan 	feud.Barony
}

func (c *敦煌DunhuangCounty) BDunhuang敦煌() feud.Barony {
	return c.敦煌Dunhuang
}
    
func (c *敦煌DunhuangCounty) BFangpan方盘() feud.Barony {
	return c.方盘Fangpan
}
    
func (c *敦煌DunhuangCounty) BHongliuwan红柳湾() feud.Barony {
	return c.红柳湾Hongliuwan
}
    
func (c *敦煌DunhuangCounty) BMogao莫高() feud.Barony {
	return c.莫高Mogao
}
    
func (c *敦煌DunhuangCounty) BXuanquanzhi悬泉置() feud.Barony {
	return c.悬泉置Xuanquanzhi
}
    
func (c *敦煌DunhuangCounty) BYueyaquan月牙泉() feud.Barony {
	return c.月牙泉Yueyaquan
}
    
func (c *敦煌DunhuangCounty) BYumenquan玉门泉() feud.Barony {
	return c.玉门泉Yumenquan
}
    
var CDunhuang敦煌 DunhuangCounty = &敦煌DunhuangCounty{}

func init() {
	f := CDunhuang敦煌.(*敦煌DunhuangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1448",
		Title:     "dunhuang",
		TitleName: "敦煌",
		TitleCode: "c_dunhuang",
		Baronies:  map[string]feud.Barony{},
	}

	f.敦煌Dunhuang = BDunhuang敦煌
	f.敦煌Dunhuang.SetParent(f)
	
	f.方盘Fangpan = BFangpan方盘
	f.方盘Fangpan.SetParent(f)
	
	f.红柳湾Hongliuwan = BHongliuwan红柳湾
	f.红柳湾Hongliuwan.SetParent(f)
	
	f.莫高Mogao = BMogao莫高
	f.莫高Mogao.SetParent(f)
	
	f.悬泉置Xuanquanzhi = BXuanquanzhi悬泉置
	f.悬泉置Xuanquanzhi.SetParent(f)
	
	f.月牙泉Yueyaquan = BYueyaquan月牙泉
	f.月牙泉Yueyaquan.SetParent(f)
	
	f.玉门泉Yumenquan = BYumenquan玉门泉
	f.玉门泉Yumenquan.SetParent(f)
	
}
