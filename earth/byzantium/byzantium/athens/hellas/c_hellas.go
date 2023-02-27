package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HellasCounty interface {
    feud.County
    BAmfissa阿姆菲萨() 	feud.Barony
    BAmphissa安菲萨() 	feud.Barony
    BItea柳城() 	feud.Barony
    BKastrinitsi卡斯特里尼齐() 	feud.Barony
    BLevadhia莱瓦贾() 	feud.Barony
    BMarkrynia马尔克利尼亚() 	feud.Barony
    BThebes底比斯() 	feud.Barony
}

type 底比斯HellasCounty struct {
	feud.BaseCounty
	阿姆菲萨Amfissa 	feud.Barony
	安菲萨Amphissa 	feud.Barony
	柳城Itea 	feud.Barony
	卡斯特里尼齐Kastrinitsi 	feud.Barony
	莱瓦贾Levadhia 	feud.Barony
	马尔克利尼亚Markrynia 	feud.Barony
	底比斯Thebes 	feud.Barony
}

func (c *底比斯HellasCounty) BAmfissa阿姆菲萨() feud.Barony {
	return c.阿姆菲萨Amfissa
}
    
func (c *底比斯HellasCounty) BAmphissa安菲萨() feud.Barony {
	return c.安菲萨Amphissa
}
    
func (c *底比斯HellasCounty) BItea柳城() feud.Barony {
	return c.柳城Itea
}
    
func (c *底比斯HellasCounty) BKastrinitsi卡斯特里尼齐() feud.Barony {
	return c.卡斯特里尼齐Kastrinitsi
}
    
func (c *底比斯HellasCounty) BLevadhia莱瓦贾() feud.Barony {
	return c.莱瓦贾Levadhia
}
    
func (c *底比斯HellasCounty) BMarkrynia马尔克利尼亚() feud.Barony {
	return c.马尔克利尼亚Markrynia
}
    
func (c *底比斯HellasCounty) BThebes底比斯() feud.Barony {
	return c.底比斯Thebes
}
    
var CHellas底比斯 HellasCounty = &底比斯HellasCounty{}

func init() {
	f := CHellas底比斯.(*底比斯HellasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "475",
		Title:     "hellas",
		TitleName: "底比斯",
		TitleCode: "c_hellas",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆菲萨Amfissa = BAmfissa阿姆菲萨
	f.阿姆菲萨Amfissa.SetParent(f)
	
	f.安菲萨Amphissa = BAmphissa安菲萨
	f.安菲萨Amphissa.SetParent(f)
	
	f.柳城Itea = BItea柳城
	f.柳城Itea.SetParent(f)
	
	f.卡斯特里尼齐Kastrinitsi = BKastrinitsi卡斯特里尼齐
	f.卡斯特里尼齐Kastrinitsi.SetParent(f)
	
	f.莱瓦贾Levadhia = BLevadhia莱瓦贾
	f.莱瓦贾Levadhia.SetParent(f)
	
	f.马尔克利尼亚Markrynia = BMarkrynia马尔克利尼亚
	f.马尔克利尼亚Markrynia.SetParent(f)
	
	f.底比斯Thebes = BThebes底比斯
	f.底比斯Thebes.SetParent(f)
	
}
