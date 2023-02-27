package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShemakhaCounty interface {
    feud.County
    BAhemakha舍马哈() 	feud.Barony
    BAnig阿内赫() 	feud.Barony
    BChiraggala奇拉赫_卡拉() 	feud.Barony
    BKhachmaz哈奇马兹() 	feud.Barony
    BKhil希利() 	feud.Barony
    BMaraza马拉济() 	feud.Barony
    BQuba库巴() 	feud.Barony
    BShikhlar什克拉尔() 	feud.Barony
}

type 舍马哈ShemakhaCounty struct {
	feud.BaseCounty
	舍马哈Ahemakha 	feud.Barony
	阿内赫Anig 	feud.Barony
	奇拉赫_卡拉Chiraggala 	feud.Barony
	哈奇马兹Khachmaz 	feud.Barony
	希利Khil 	feud.Barony
	马拉济Maraza 	feud.Barony
	库巴Quba 	feud.Barony
	什克拉尔Shikhlar 	feud.Barony
}

func (c *舍马哈ShemakhaCounty) BAhemakha舍马哈() feud.Barony {
	return c.舍马哈Ahemakha
}
    
func (c *舍马哈ShemakhaCounty) BAnig阿内赫() feud.Barony {
	return c.阿内赫Anig
}
    
func (c *舍马哈ShemakhaCounty) BChiraggala奇拉赫_卡拉() feud.Barony {
	return c.奇拉赫_卡拉Chiraggala
}
    
func (c *舍马哈ShemakhaCounty) BKhachmaz哈奇马兹() feud.Barony {
	return c.哈奇马兹Khachmaz
}
    
func (c *舍马哈ShemakhaCounty) BKhil希利() feud.Barony {
	return c.希利Khil
}
    
func (c *舍马哈ShemakhaCounty) BMaraza马拉济() feud.Barony {
	return c.马拉济Maraza
}
    
func (c *舍马哈ShemakhaCounty) BQuba库巴() feud.Barony {
	return c.库巴Quba
}
    
func (c *舍马哈ShemakhaCounty) BShikhlar什克拉尔() feud.Barony {
	return c.什克拉尔Shikhlar
}
    
var CShemakha舍马哈 ShemakhaCounty = &舍马哈ShemakhaCounty{}

func init() {
	f := CShemakha舍马哈.(*舍马哈ShemakhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "669",
		Title:     "shemakha",
		TitleName: "舍马哈",
		TitleCode: "c_shemakha",
		Baronies:  map[string]feud.Barony{},
	}

	f.舍马哈Ahemakha = BAhemakha舍马哈
	f.舍马哈Ahemakha.SetParent(f)
	
	f.阿内赫Anig = BAnig阿内赫
	f.阿内赫Anig.SetParent(f)
	
	f.奇拉赫_卡拉Chiraggala = BChiraggala奇拉赫_卡拉
	f.奇拉赫_卡拉Chiraggala.SetParent(f)
	
	f.哈奇马兹Khachmaz = BKhachmaz哈奇马兹
	f.哈奇马兹Khachmaz.SetParent(f)
	
	f.希利Khil = BKhil希利
	f.希利Khil.SetParent(f)
	
	f.马拉济Maraza = BMaraza马拉济
	f.马拉济Maraza.SetParent(f)
	
	f.库巴Quba = BQuba库巴
	f.库巴Quba.SetParent(f)
	
	f.什克拉尔Shikhlar = BShikhlar什克拉尔
	f.什克拉尔Shikhlar.SetParent(f)
	
}
