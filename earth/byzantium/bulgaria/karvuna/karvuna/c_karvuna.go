package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarvunaCounty interface {
    feud.County
    BDobrich巴尔奇克() 	feud.Barony
    BKaliakra卡利亚克拉() 	feud.Barony
    BKarvuna卡尔武纳() 	feud.Barony
    BPrezlav普雷斯拉夫() 	feud.Barony
    BSilistria锡利斯特里亚() 	feud.Barony
    BSmyadovo斯米亚多沃() 	feud.Barony
    BVarbitsz沃尔比察() 	feud.Barony
    BVenets韦内茨() 	feud.Barony
}

type 卡尔武纳KarvunaCounty struct {
	feud.BaseCounty
	巴尔奇克Dobrich 	feud.Barony
	卡利亚克拉Kaliakra 	feud.Barony
	卡尔武纳Karvuna 	feud.Barony
	普雷斯拉夫Prezlav 	feud.Barony
	锡利斯特里亚Silistria 	feud.Barony
	斯米亚多沃Smyadovo 	feud.Barony
	沃尔比察Varbitsz 	feud.Barony
	韦内茨Venets 	feud.Barony
}

func (c *卡尔武纳KarvunaCounty) BDobrich巴尔奇克() feud.Barony {
	return c.巴尔奇克Dobrich
}
    
func (c *卡尔武纳KarvunaCounty) BKaliakra卡利亚克拉() feud.Barony {
	return c.卡利亚克拉Kaliakra
}
    
func (c *卡尔武纳KarvunaCounty) BKarvuna卡尔武纳() feud.Barony {
	return c.卡尔武纳Karvuna
}
    
func (c *卡尔武纳KarvunaCounty) BPrezlav普雷斯拉夫() feud.Barony {
	return c.普雷斯拉夫Prezlav
}
    
func (c *卡尔武纳KarvunaCounty) BSilistria锡利斯特里亚() feud.Barony {
	return c.锡利斯特里亚Silistria
}
    
func (c *卡尔武纳KarvunaCounty) BSmyadovo斯米亚多沃() feud.Barony {
	return c.斯米亚多沃Smyadovo
}
    
func (c *卡尔武纳KarvunaCounty) BVarbitsz沃尔比察() feud.Barony {
	return c.沃尔比察Varbitsz
}
    
func (c *卡尔武纳KarvunaCounty) BVenets韦内茨() feud.Barony {
	return c.韦内茨Venets
}
    
var CKarvuna卡尔武纳 KarvunaCounty = &卡尔武纳KarvunaCounty{}

func init() {
	f := CKarvuna卡尔武纳.(*卡尔武纳KarvunaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "509",
		Title:     "karvuna",
		TitleName: "卡尔武纳",
		TitleCode: "c_karvuna",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔奇克Dobrich = BDobrich巴尔奇克
	f.巴尔奇克Dobrich.SetParent(f)
	
	f.卡利亚克拉Kaliakra = BKaliakra卡利亚克拉
	f.卡利亚克拉Kaliakra.SetParent(f)
	
	f.卡尔武纳Karvuna = BKarvuna卡尔武纳
	f.卡尔武纳Karvuna.SetParent(f)
	
	f.普雷斯拉夫Prezlav = BPrezlav普雷斯拉夫
	f.普雷斯拉夫Prezlav.SetParent(f)
	
	f.锡利斯特里亚Silistria = BSilistria锡利斯特里亚
	f.锡利斯特里亚Silistria.SetParent(f)
	
	f.斯米亚多沃Smyadovo = BSmyadovo斯米亚多沃
	f.斯米亚多沃Smyadovo.SetParent(f)
	
	f.沃尔比察Varbitsz = BVarbitsz沃尔比察
	f.沃尔比察Varbitsz.SetParent(f)
	
	f.韦内茨Venets = BVenets韦内茨
	f.韦内茨Venets.SetParent(f)
	
}
