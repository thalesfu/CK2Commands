package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShirvanCounty interface {
    feud.County
    BAbsheron阿普舍龙() 	feud.Barony
    BAltiagay阿勒提干() 	feud.Barony
    BBaku巴库() 	feud.Barony
    BKhizirzinda希济尔津达() 	feud.Barony
    BLankaran连科兰() 	feud.Barony
    BSalyan萨利杨() 	feud.Barony
    BShakhriyar沙赫里亚尔() 	feud.Barony
    BShirvan希尔凡() 	feud.Barony
}

type 希尔凡ShirvanCounty struct {
	feud.BaseCounty
	阿普舍龙Absheron 	feud.Barony
	阿勒提干Altiagay 	feud.Barony
	巴库Baku 	feud.Barony
	希济尔津达Khizirzinda 	feud.Barony
	连科兰Lankaran 	feud.Barony
	萨利杨Salyan 	feud.Barony
	沙赫里亚尔Shakhriyar 	feud.Barony
	希尔凡Shirvan 	feud.Barony
}

func (c *希尔凡ShirvanCounty) BAbsheron阿普舍龙() feud.Barony {
	return c.阿普舍龙Absheron
}
    
func (c *希尔凡ShirvanCounty) BAltiagay阿勒提干() feud.Barony {
	return c.阿勒提干Altiagay
}
    
func (c *希尔凡ShirvanCounty) BBaku巴库() feud.Barony {
	return c.巴库Baku
}
    
func (c *希尔凡ShirvanCounty) BKhizirzinda希济尔津达() feud.Barony {
	return c.希济尔津达Khizirzinda
}
    
func (c *希尔凡ShirvanCounty) BLankaran连科兰() feud.Barony {
	return c.连科兰Lankaran
}
    
func (c *希尔凡ShirvanCounty) BSalyan萨利杨() feud.Barony {
	return c.萨利杨Salyan
}
    
func (c *希尔凡ShirvanCounty) BShakhriyar沙赫里亚尔() feud.Barony {
	return c.沙赫里亚尔Shakhriyar
}
    
func (c *希尔凡ShirvanCounty) BShirvan希尔凡() feud.Barony {
	return c.希尔凡Shirvan
}
    
var CShirvan希尔凡 ShirvanCounty = &希尔凡ShirvanCounty{}

func init() {
	f := CShirvan希尔凡.(*希尔凡ShirvanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "668",
		Title:     "shirvan",
		TitleName: "希尔凡",
		TitleCode: "c_shirvan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普舍龙Absheron = BAbsheron阿普舍龙
	f.阿普舍龙Absheron.SetParent(f)
	
	f.阿勒提干Altiagay = BAltiagay阿勒提干
	f.阿勒提干Altiagay.SetParent(f)
	
	f.巴库Baku = BBaku巴库
	f.巴库Baku.SetParent(f)
	
	f.希济尔津达Khizirzinda = BKhizirzinda希济尔津达
	f.希济尔津达Khizirzinda.SetParent(f)
	
	f.连科兰Lankaran = BLankaran连科兰
	f.连科兰Lankaran.SetParent(f)
	
	f.萨利杨Salyan = BSalyan萨利杨
	f.萨利杨Salyan.SetParent(f)
	
	f.沙赫里亚尔Shakhriyar = BShakhriyar沙赫里亚尔
	f.沙赫里亚尔Shakhriyar.SetParent(f)
	
	f.希尔凡Shirvan = BShirvan希尔凡
	f.希尔凡Shirvan.SetParent(f)
	
}
