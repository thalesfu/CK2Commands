package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JiuquanCounty interface {
    feud.County
    BFulu福禄() 	feud.Barony
    BGuantou观头() 	feud.Barony
    BJiuquan酒泉() 	feud.Barony
    BKongtong崆峒() 	feud.Barony
    BLongdi陇狄() 	feud.Barony
    BSuzhou肃州() 	feud.Barony
    BXianshi仙石() 	feud.Barony
}

type 酒泉JiuquanCounty struct {
	feud.BaseCounty
	福禄Fulu 	feud.Barony
	观头Guantou 	feud.Barony
	酒泉Jiuquan 	feud.Barony
	崆峒Kongtong 	feud.Barony
	陇狄Longdi 	feud.Barony
	肃州Suzhou 	feud.Barony
	仙石Xianshi 	feud.Barony
}

func (c *酒泉JiuquanCounty) BFulu福禄() feud.Barony {
	return c.福禄Fulu
}
    
func (c *酒泉JiuquanCounty) BGuantou观头() feud.Barony {
	return c.观头Guantou
}
    
func (c *酒泉JiuquanCounty) BJiuquan酒泉() feud.Barony {
	return c.酒泉Jiuquan
}
    
func (c *酒泉JiuquanCounty) BKongtong崆峒() feud.Barony {
	return c.崆峒Kongtong
}
    
func (c *酒泉JiuquanCounty) BLongdi陇狄() feud.Barony {
	return c.陇狄Longdi
}
    
func (c *酒泉JiuquanCounty) BSuzhou肃州() feud.Barony {
	return c.肃州Suzhou
}
    
func (c *酒泉JiuquanCounty) BXianshi仙石() feud.Barony {
	return c.仙石Xianshi
}
    
var CJiuquan酒泉 JiuquanCounty = &酒泉JiuquanCounty{}

func init() {
	f := CJiuquan酒泉.(*酒泉JiuquanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1513",
		Title:     "jiuquan",
		TitleName: "酒泉",
		TitleCode: "c_jiuquan",
		Baronies:  map[string]feud.Barony{},
	}

	f.福禄Fulu = BFulu福禄
	f.福禄Fulu.SetParent(f)
	
	f.观头Guantou = BGuantou观头
	f.观头Guantou.SetParent(f)
	
	f.酒泉Jiuquan = BJiuquan酒泉
	f.酒泉Jiuquan.SetParent(f)
	
	f.崆峒Kongtong = BKongtong崆峒
	f.崆峒Kongtong.SetParent(f)
	
	f.陇狄Longdi = BLongdi陇狄
	f.陇狄Longdi.SetParent(f)
	
	f.肃州Suzhou = BSuzhou肃州
	f.肃州Suzhou.SetParent(f)
	
	f.仙石Xianshi = BXianshi仙石
	f.仙石Xianshi.SetParent(f)
	
}
