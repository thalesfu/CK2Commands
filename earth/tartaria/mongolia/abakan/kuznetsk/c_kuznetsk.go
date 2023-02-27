package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuznetskCounty interface {
    feud.County
    BBelogorsk别洛戈尔斯克() 	feud.Barony
    BChernoye_kuznetsk乔尔诺耶() 	feud.Barony
    BKuznetsk库兹涅茨克() 	feud.Barony
    BMedvezhka_kuznetsk梅德韦日卡() 	feud.Barony
    BMutnyy_kuznetsk穆恃内() 	feud.Barony
    BNaryk纳雷克() 	feud.Barony
    BSarala_kuznetsk萨拉拉() 	feud.Barony
}

type 库兹涅茨克KuznetskCounty struct {
	feud.BaseCounty
	别洛戈尔斯克Belogorsk 	feud.Barony
	乔尔诺耶Chernoye_kuznetsk 	feud.Barony
	库兹涅茨克Kuznetsk 	feud.Barony
	梅德韦日卡Medvezhka_kuznetsk 	feud.Barony
	穆恃内Mutnyy_kuznetsk 	feud.Barony
	纳雷克Naryk 	feud.Barony
	萨拉拉Sarala_kuznetsk 	feud.Barony
}

func (c *库兹涅茨克KuznetskCounty) BBelogorsk别洛戈尔斯克() feud.Barony {
	return c.别洛戈尔斯克Belogorsk
}
    
func (c *库兹涅茨克KuznetskCounty) BChernoye_kuznetsk乔尔诺耶() feud.Barony {
	return c.乔尔诺耶Chernoye_kuznetsk
}
    
func (c *库兹涅茨克KuznetskCounty) BKuznetsk库兹涅茨克() feud.Barony {
	return c.库兹涅茨克Kuznetsk
}
    
func (c *库兹涅茨克KuznetskCounty) BMedvezhka_kuznetsk梅德韦日卡() feud.Barony {
	return c.梅德韦日卡Medvezhka_kuznetsk
}
    
func (c *库兹涅茨克KuznetskCounty) BMutnyy_kuznetsk穆恃内() feud.Barony {
	return c.穆恃内Mutnyy_kuznetsk
}
    
func (c *库兹涅茨克KuznetskCounty) BNaryk纳雷克() feud.Barony {
	return c.纳雷克Naryk
}
    
func (c *库兹涅茨克KuznetskCounty) BSarala_kuznetsk萨拉拉() feud.Barony {
	return c.萨拉拉Sarala_kuznetsk
}
    
var CKuznetsk库兹涅茨克 KuznetskCounty = &库兹涅茨克KuznetskCounty{}

func init() {
	f := CKuznetsk库兹涅茨克.(*库兹涅茨克KuznetskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1901",
		Title:     "kuznetsk",
		TitleName: "库兹涅茨克",
		TitleCode: "c_kuznetsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别洛戈尔斯克Belogorsk = BBelogorsk别洛戈尔斯克
	f.别洛戈尔斯克Belogorsk.SetParent(f)
	
	f.乔尔诺耶Chernoye_kuznetsk = BChernoye_kuznetsk乔尔诺耶
	f.乔尔诺耶Chernoye_kuznetsk.SetParent(f)
	
	f.库兹涅茨克Kuznetsk = BKuznetsk库兹涅茨克
	f.库兹涅茨克Kuznetsk.SetParent(f)
	
	f.梅德韦日卡Medvezhka_kuznetsk = BMedvezhka_kuznetsk梅德韦日卡
	f.梅德韦日卡Medvezhka_kuznetsk.SetParent(f)
	
	f.穆恃内Mutnyy_kuznetsk = BMutnyy_kuznetsk穆恃内
	f.穆恃内Mutnyy_kuznetsk.SetParent(f)
	
	f.纳雷克Naryk = BNaryk纳雷克
	f.纳雷克Naryk.SetParent(f)
	
	f.萨拉拉Sarala_kuznetsk = BSarala_kuznetsk萨拉拉
	f.萨拉拉Sarala_kuznetsk.SetParent(f)
	
}
