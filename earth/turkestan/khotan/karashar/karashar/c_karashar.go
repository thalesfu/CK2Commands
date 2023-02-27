package karashar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarasharCounty interface {
    feud.County
    BKarashar喀喇沙尔() 	feud.Barony
    BKorla坤闾() 	feud.Barony
    BPaohsinchi鲍兴集() 	feud.Barony
    BShanguo山国() 	feud.Barony
    BTiemenguan铁门关() 	feud.Barony
    BWeili尉犁() 	feud.Barony
    BWeixu危须() 	feud.Barony
}

type 喀喇沙尔KarasharCounty struct {
	feud.BaseCounty
	喀喇沙尔Karashar 	feud.Barony
	坤闾Korla 	feud.Barony
	鲍兴集Paohsinchi 	feud.Barony
	山国Shanguo 	feud.Barony
	铁门关Tiemenguan 	feud.Barony
	尉犁Weili 	feud.Barony
	危须Weixu 	feud.Barony
}

func (c *喀喇沙尔KarasharCounty) BKarashar喀喇沙尔() feud.Barony {
	return c.喀喇沙尔Karashar
}
    
func (c *喀喇沙尔KarasharCounty) BKorla坤闾() feud.Barony {
	return c.坤闾Korla
}
    
func (c *喀喇沙尔KarasharCounty) BPaohsinchi鲍兴集() feud.Barony {
	return c.鲍兴集Paohsinchi
}
    
func (c *喀喇沙尔KarasharCounty) BShanguo山国() feud.Barony {
	return c.山国Shanguo
}
    
func (c *喀喇沙尔KarasharCounty) BTiemenguan铁门关() feud.Barony {
	return c.铁门关Tiemenguan
}
    
func (c *喀喇沙尔KarasharCounty) BWeili尉犁() feud.Barony {
	return c.尉犁Weili
}
    
func (c *喀喇沙尔KarasharCounty) BWeixu危须() feud.Barony {
	return c.危须Weixu
}
    
var CKarashar喀喇沙尔 KarasharCounty = &喀喇沙尔KarasharCounty{}

func init() {
	f := CKarashar喀喇沙尔.(*喀喇沙尔KarasharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1443",
		Title:     "karashar",
		TitleName: "喀喇沙尔",
		TitleCode: "c_karashar",
		Baronies:  map[string]feud.Barony{},
	}

	f.喀喇沙尔Karashar = BKarashar喀喇沙尔
	f.喀喇沙尔Karashar.SetParent(f)
	
	f.坤闾Korla = BKorla坤闾
	f.坤闾Korla.SetParent(f)
	
	f.鲍兴集Paohsinchi = BPaohsinchi鲍兴集
	f.鲍兴集Paohsinchi.SetParent(f)
	
	f.山国Shanguo = BShanguo山国
	f.山国Shanguo.SetParent(f)
	
	f.铁门关Tiemenguan = BTiemenguan铁门关
	f.铁门关Tiemenguan.SetParent(f)
	
	f.尉犁Weili = BWeili尉犁
	f.尉犁Weili.SetParent(f)
	
	f.危须Weixu = BWeixu危须
	f.危须Weixu.SetParent(f)
	
}
