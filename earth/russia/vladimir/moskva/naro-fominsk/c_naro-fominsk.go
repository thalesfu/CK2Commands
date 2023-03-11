package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Naro-fominskCounty interface {
    feud.County
    BChukovo丘科沃() 	feud.Barony
    BKlimovka克利莫夫卡() 	feud.Barony
    BKubinka库宾卡() 	feud.Barony
    BLopasnya洛帕斯尼亚() 	feud.Barony
    BNaro-fominsk纳罗福明斯克() 	feud.Barony
    BTroitsk特罗伊茨克() 	feud.Barony
    BVoronovo沃罗诺沃() 	feud.Barony
}

type 纳罗福明斯克Naro-fominskCounty struct {
	feud.BaseCounty
	丘科沃Chukovo 	feud.Barony
	克利莫夫卡Klimovka 	feud.Barony
	库宾卡Kubinka 	feud.Barony
	洛帕斯尼亚Lopasnya 	feud.Barony
	纳罗福明斯克Naro-fominsk 	feud.Barony
	特罗伊茨克Troitsk 	feud.Barony
	沃罗诺沃Voronovo 	feud.Barony
}

func (c *纳罗福明斯克Naro-fominskCounty) BChukovo丘科沃() feud.Barony {
	return c.丘科沃Chukovo
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BKlimovka克利莫夫卡() feud.Barony {
	return c.克利莫夫卡Klimovka
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BKubinka库宾卡() feud.Barony {
	return c.库宾卡Kubinka
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BLopasnya洛帕斯尼亚() feud.Barony {
	return c.洛帕斯尼亚Lopasnya
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BNaro-fominsk纳罗福明斯克() feud.Barony {
	return c.纳罗福明斯克Naro-fominsk
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BTroitsk特罗伊茨克() feud.Barony {
	return c.特罗伊茨克Troitsk
}
    
func (c *纳罗福明斯克Naro-fominskCounty) BVoronovo沃罗诺沃() feud.Barony {
	return c.沃罗诺沃Voronovo
}
    
var CNaro-fominsk纳罗福明斯克 Naro-fominskCounty = &纳罗福明斯克Naro-fominskCounty{}

func init() {
	f := CNaro-fominsk纳罗福明斯克.(*纳罗福明斯克Naro-fominskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1677",
		Title:     "naro-fominsk",
		TitleName: "纳罗福明斯克",
		TitleCode: "c_naro-fominsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.丘科沃Chukovo = BChukovo丘科沃
	f.丘科沃Chukovo.SetParent(f)
	
	f.克利莫夫卡Klimovka = BKlimovka克利莫夫卡
	f.克利莫夫卡Klimovka.SetParent(f)
	
	f.库宾卡Kubinka = BKubinka库宾卡
	f.库宾卡Kubinka.SetParent(f)
	
	f.洛帕斯尼亚Lopasnya = BLopasnya洛帕斯尼亚
	f.洛帕斯尼亚Lopasnya.SetParent(f)
	
	f.纳罗福明斯克Naro-fominsk = BNaro-fominsk纳罗福明斯克
	f.纳罗福明斯克Naro-fominsk.SetParent(f)
	
	f.特罗伊茨克Troitsk = BTroitsk特罗伊茨克
	f.特罗伊茨克Troitsk.SetParent(f)
	
	f.沃罗诺沃Voronovo = BVoronovo沃罗诺沃
	f.沃罗诺沃Voronovo.SetParent(f)
	
}
