package jaen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JaenCounty interface {
    feud.County
    BAndujar安杜哈尔() 	feud.Barony
    BBaeza巴埃萨() 	feud.Barony
    BHuescar韦斯卡尔() 	feud.Barony
    BJaen哈恩() 	feud.Barony
    BQuesada克萨达() 	feud.Barony
    BSegura塞古拉() 	feud.Barony
    BTiscar蒂斯卡尔() 	feud.Barony
}

type 哈恩JaenCounty struct {
	feud.BaseCounty
	安杜哈尔Andujar 	feud.Barony
	巴埃萨Baeza 	feud.Barony
	韦斯卡尔Huescar 	feud.Barony
	哈恩Jaen 	feud.Barony
	克萨达Quesada 	feud.Barony
	塞古拉Segura 	feud.Barony
	蒂斯卡尔Tiscar 	feud.Barony
}

func (c *哈恩JaenCounty) BAndujar安杜哈尔() feud.Barony {
	return c.安杜哈尔Andujar
}
    
func (c *哈恩JaenCounty) BBaeza巴埃萨() feud.Barony {
	return c.巴埃萨Baeza
}
    
func (c *哈恩JaenCounty) BHuescar韦斯卡尔() feud.Barony {
	return c.韦斯卡尔Huescar
}
    
func (c *哈恩JaenCounty) BJaen哈恩() feud.Barony {
	return c.哈恩Jaen
}
    
func (c *哈恩JaenCounty) BQuesada克萨达() feud.Barony {
	return c.克萨达Quesada
}
    
func (c *哈恩JaenCounty) BSegura塞古拉() feud.Barony {
	return c.塞古拉Segura
}
    
func (c *哈恩JaenCounty) BTiscar蒂斯卡尔() feud.Barony {
	return c.蒂斯卡尔Tiscar
}
    
var CJaen哈恩 JaenCounty = &哈恩JaenCounty{}

func init() {
	f := CJaen哈恩.(*哈恩JaenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1974",
		Title:     "jaen",
		TitleName: "哈恩",
		TitleCode: "c_jaen",
		Baronies:  map[string]feud.Barony{},
	}

	f.安杜哈尔Andujar = BAndujar安杜哈尔
	f.安杜哈尔Andujar.SetParent(f)
	
	f.巴埃萨Baeza = BBaeza巴埃萨
	f.巴埃萨Baeza.SetParent(f)
	
	f.韦斯卡尔Huescar = BHuescar韦斯卡尔
	f.韦斯卡尔Huescar.SetParent(f)
	
	f.哈恩Jaen = BJaen哈恩
	f.哈恩Jaen.SetParent(f)
	
	f.克萨达Quesada = BQuesada克萨达
	f.克萨达Quesada.SetParent(f)
	
	f.塞古拉Segura = BSegura塞古拉
	f.塞古拉Segura.SetParent(f)
	
	f.蒂斯卡尔Tiscar = BTiscar蒂斯卡尔
	f.蒂斯卡尔Tiscar.SetParent(f)
	
}
