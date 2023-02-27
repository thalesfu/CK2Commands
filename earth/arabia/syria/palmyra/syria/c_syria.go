package syria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SyriaCounty interface {
    feud.County
    BAdra阿德拉() 	feud.Barony
    BAmrah阿玛拉赫() 	feud.Barony
    BBaly巴鲁() 	feud.Barony
    BBuraq布拉格() 	feud.Barony
    BHayjanah海贾纳() 	feud.Barony
    BJarba扎巴() 	feud.Barony
    BOtaybah奥太巴() 	feud.Barony
    BQasmiye喀斯米耶() 	feud.Barony
}

type 布斯拉SyriaCounty struct {
	feud.BaseCounty
	阿德拉Adra 	feud.Barony
	阿玛拉赫Amrah 	feud.Barony
	巴鲁Baly 	feud.Barony
	布拉格Buraq 	feud.Barony
	海贾纳Hayjanah 	feud.Barony
	扎巴Jarba 	feud.Barony
	奥太巴Otaybah 	feud.Barony
	喀斯米耶Qasmiye 	feud.Barony
}

func (c *布斯拉SyriaCounty) BAdra阿德拉() feud.Barony {
	return c.阿德拉Adra
}
    
func (c *布斯拉SyriaCounty) BAmrah阿玛拉赫() feud.Barony {
	return c.阿玛拉赫Amrah
}
    
func (c *布斯拉SyriaCounty) BBaly巴鲁() feud.Barony {
	return c.巴鲁Baly
}
    
func (c *布斯拉SyriaCounty) BBuraq布拉格() feud.Barony {
	return c.布拉格Buraq
}
    
func (c *布斯拉SyriaCounty) BHayjanah海贾纳() feud.Barony {
	return c.海贾纳Hayjanah
}
    
func (c *布斯拉SyriaCounty) BJarba扎巴() feud.Barony {
	return c.扎巴Jarba
}
    
func (c *布斯拉SyriaCounty) BOtaybah奥太巴() feud.Barony {
	return c.奥太巴Otaybah
}
    
func (c *布斯拉SyriaCounty) BQasmiye喀斯米耶() feud.Barony {
	return c.喀斯米耶Qasmiye
}
    
var CSyria布斯拉 SyriaCounty = &布斯拉SyriaCounty{}

func init() {
	f := CSyria布斯拉.(*布斯拉SyriaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "727",
		Title:     "syria",
		TitleName: "布斯拉",
		TitleCode: "c_syria",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德拉Adra = BAdra阿德拉
	f.阿德拉Adra.SetParent(f)
	
	f.阿玛拉赫Amrah = BAmrah阿玛拉赫
	f.阿玛拉赫Amrah.SetParent(f)
	
	f.巴鲁Baly = BBaly巴鲁
	f.巴鲁Baly.SetParent(f)
	
	f.布拉格Buraq = BBuraq布拉格
	f.布拉格Buraq.SetParent(f)
	
	f.海贾纳Hayjanah = BHayjanah海贾纳
	f.海贾纳Hayjanah.SetParent(f)
	
	f.扎巴Jarba = BJarba扎巴
	f.扎巴Jarba.SetParent(f)
	
	f.奥太巴Otaybah = BOtaybah奥太巴
	f.奥太巴Otaybah.SetParent(f)
	
	f.喀斯米耶Qasmiye = BQasmiye喀斯米耶
	f.喀斯米耶Qasmiye.SetParent(f)
	
}
