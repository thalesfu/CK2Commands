package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OgliastraCounty interface {
    feud.County
    BBallao巴劳() 	feud.Barony
    BBari_sardo巴里萨尔多() 	feud.Barony
    BJerzu耶尔祖() 	feud.Barony
    BMuravera穆拉韦拉() 	feud.Barony
    BOgliastra奥利亚斯特拉() 	feud.Barony
    BTortoli托尔托利() 	feud.Barony
    BUrzulei乌尔祖莱() 	feud.Barony
}

type 奥利亚斯特拉OgliastraCounty struct {
	feud.BaseCounty
	巴劳Ballao 	feud.Barony
	巴里萨尔多Bari_sardo 	feud.Barony
	耶尔祖Jerzu 	feud.Barony
	穆拉韦拉Muravera 	feud.Barony
	奥利亚斯特拉Ogliastra 	feud.Barony
	托尔托利Tortoli 	feud.Barony
	乌尔祖莱Urzulei 	feud.Barony
}

func (c *奥利亚斯特拉OgliastraCounty) BBallao巴劳() feud.Barony {
	return c.巴劳Ballao
}
    
func (c *奥利亚斯特拉OgliastraCounty) BBari_sardo巴里萨尔多() feud.Barony {
	return c.巴里萨尔多Bari_sardo
}
    
func (c *奥利亚斯特拉OgliastraCounty) BJerzu耶尔祖() feud.Barony {
	return c.耶尔祖Jerzu
}
    
func (c *奥利亚斯特拉OgliastraCounty) BMuravera穆拉韦拉() feud.Barony {
	return c.穆拉韦拉Muravera
}
    
func (c *奥利亚斯特拉OgliastraCounty) BOgliastra奥利亚斯特拉() feud.Barony {
	return c.奥利亚斯特拉Ogliastra
}
    
func (c *奥利亚斯特拉OgliastraCounty) BTortoli托尔托利() feud.Barony {
	return c.托尔托利Tortoli
}
    
func (c *奥利亚斯特拉OgliastraCounty) BUrzulei乌尔祖莱() feud.Barony {
	return c.乌尔祖莱Urzulei
}
    
var COgliastra奥利亚斯特拉 OgliastraCounty = &奥利亚斯特拉OgliastraCounty{}

func init() {
	f := COgliastra奥利亚斯特拉.(*奥利亚斯特拉OgliastraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1560",
		Title:     "ogliastra",
		TitleName: "奥利亚斯特拉",
		TitleCode: "c_ogliastra",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴劳Ballao = BBallao巴劳
	f.巴劳Ballao.SetParent(f)
	
	f.巴里萨尔多Bari_sardo = BBari_sardo巴里萨尔多
	f.巴里萨尔多Bari_sardo.SetParent(f)
	
	f.耶尔祖Jerzu = BJerzu耶尔祖
	f.耶尔祖Jerzu.SetParent(f)
	
	f.穆拉韦拉Muravera = BMuravera穆拉韦拉
	f.穆拉韦拉Muravera.SetParent(f)
	
	f.奥利亚斯特拉Ogliastra = BOgliastra奥利亚斯特拉
	f.奥利亚斯特拉Ogliastra.SetParent(f)
	
	f.托尔托利Tortoli = BTortoli托尔托利
	f.托尔托利Tortoli.SetParent(f)
	
	f.乌尔祖莱Urzulei = BUrzulei乌尔祖莱
	f.乌尔祖莱Urzulei.SetParent(f)
	
}
