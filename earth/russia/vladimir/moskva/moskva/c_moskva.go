package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoskvaCounty interface {
    feud.County
    BIstra伊斯特拉() 	feud.Barony
    BKhimki希姆基() 	feud.Barony
    BKrasnoya克拉斯诺亚() 	feud.Barony
    BLyubertsy柳别尔齐() 	feud.Barony
    BMoskva莫斯科() 	feud.Barony
    BMytishchi梅季希() 	feud.Barony
    BPodolsk波多利斯克() 	feud.Barony
}

type 莫斯科MoskvaCounty struct {
	feud.BaseCounty
	伊斯特拉Istra 	feud.Barony
	希姆基Khimki 	feud.Barony
	克拉斯诺亚Krasnoya 	feud.Barony
	柳别尔齐Lyubertsy 	feud.Barony
	莫斯科Moskva 	feud.Barony
	梅季希Mytishchi 	feud.Barony
	波多利斯克Podolsk 	feud.Barony
}

func (c *莫斯科MoskvaCounty) BIstra伊斯特拉() feud.Barony {
	return c.伊斯特拉Istra
}
    
func (c *莫斯科MoskvaCounty) BKhimki希姆基() feud.Barony {
	return c.希姆基Khimki
}
    
func (c *莫斯科MoskvaCounty) BKrasnoya克拉斯诺亚() feud.Barony {
	return c.克拉斯诺亚Krasnoya
}
    
func (c *莫斯科MoskvaCounty) BLyubertsy柳别尔齐() feud.Barony {
	return c.柳别尔齐Lyubertsy
}
    
func (c *莫斯科MoskvaCounty) BMoskva莫斯科() feud.Barony {
	return c.莫斯科Moskva
}
    
func (c *莫斯科MoskvaCounty) BMytishchi梅季希() feud.Barony {
	return c.梅季希Mytishchi
}
    
func (c *莫斯科MoskvaCounty) BPodolsk波多利斯克() feud.Barony {
	return c.波多利斯克Podolsk
}
    
var CMoskva莫斯科 MoskvaCounty = &莫斯科MoskvaCounty{}

func init() {
	f := CMoskva莫斯科.(*莫斯科MoskvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "575",
		Title:     "moskva",
		TitleName: "莫斯科",
		TitleCode: "c_moskva",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊斯特拉Istra = BIstra伊斯特拉
	f.伊斯特拉Istra.SetParent(f)
	
	f.希姆基Khimki = BKhimki希姆基
	f.希姆基Khimki.SetParent(f)
	
	f.克拉斯诺亚Krasnoya = BKrasnoya克拉斯诺亚
	f.克拉斯诺亚Krasnoya.SetParent(f)
	
	f.柳别尔齐Lyubertsy = BLyubertsy柳别尔齐
	f.柳别尔齐Lyubertsy.SetParent(f)
	
	f.莫斯科Moskva = BMoskva莫斯科
	f.莫斯科Moskva.SetParent(f)
	
	f.梅季希Mytishchi = BMytishchi梅季希
	f.梅季希Mytishchi.SetParent(f)
	
	f.波多利斯克Podolsk = BPodolsk波多利斯克
	f.波多利斯克Podolsk.SetParent(f)
	
}
