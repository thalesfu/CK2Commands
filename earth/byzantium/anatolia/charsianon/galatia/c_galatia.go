package galatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalatiaCounty interface {
    feud.County
    BAsponia阿斯波尼亚() 	feud.Barony
    BCarissa卡里萨() 	feud.Barony
    BGarsaura加苏拉() 	feud.Barony
    BKaracaviran卡拉贾维兰() 	feud.Barony
    BKochisar科奇希萨尔() 	feud.Barony
    BMikissos莫基索斯() 	feud.Barony
    BTavia泰维亚() 	feud.Barony
}

type 加拉太GalatiaCounty struct {
	feud.BaseCounty
	阿斯波尼亚Asponia 	feud.Barony
	卡里萨Carissa 	feud.Barony
	加苏拉Garsaura 	feud.Barony
	卡拉贾维兰Karacaviran 	feud.Barony
	科奇希萨尔Kochisar 	feud.Barony
	莫基索斯Mikissos 	feud.Barony
	泰维亚Tavia 	feud.Barony
}

func (c *加拉太GalatiaCounty) BAsponia阿斯波尼亚() feud.Barony {
	return c.阿斯波尼亚Asponia
}
    
func (c *加拉太GalatiaCounty) BCarissa卡里萨() feud.Barony {
	return c.卡里萨Carissa
}
    
func (c *加拉太GalatiaCounty) BGarsaura加苏拉() feud.Barony {
	return c.加苏拉Garsaura
}
    
func (c *加拉太GalatiaCounty) BKaracaviran卡拉贾维兰() feud.Barony {
	return c.卡拉贾维兰Karacaviran
}
    
func (c *加拉太GalatiaCounty) BKochisar科奇希萨尔() feud.Barony {
	return c.科奇希萨尔Kochisar
}
    
func (c *加拉太GalatiaCounty) BMikissos莫基索斯() feud.Barony {
	return c.莫基索斯Mikissos
}
    
func (c *加拉太GalatiaCounty) BTavia泰维亚() feud.Barony {
	return c.泰维亚Tavia
}
    
var CGalatia加拉太 GalatiaCounty = &加拉太GalatiaCounty{}

func init() {
	f := CGalatia加拉太.(*加拉太GalatiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "752",
		Title:     "galatia",
		TitleName: "加拉太",
		TitleCode: "c_galatia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯波尼亚Asponia = BAsponia阿斯波尼亚
	f.阿斯波尼亚Asponia.SetParent(f)
	
	f.卡里萨Carissa = BCarissa卡里萨
	f.卡里萨Carissa.SetParent(f)
	
	f.加苏拉Garsaura = BGarsaura加苏拉
	f.加苏拉Garsaura.SetParent(f)
	
	f.卡拉贾维兰Karacaviran = BKaracaviran卡拉贾维兰
	f.卡拉贾维兰Karacaviran.SetParent(f)
	
	f.科奇希萨尔Kochisar = BKochisar科奇希萨尔
	f.科奇希萨尔Kochisar.SetParent(f)
	
	f.莫基索斯Mikissos = BMikissos莫基索斯
	f.莫基索斯Mikissos.SetParent(f)
	
	f.泰维亚Tavia = BTavia泰维亚
	f.泰维亚Tavia.SetParent(f)
	
}
