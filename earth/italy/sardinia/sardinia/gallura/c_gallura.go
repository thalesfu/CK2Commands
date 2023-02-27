package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GalluraCounty interface {
    feud.County
    BAggius阿朱斯() 	feud.Barony
    BBicinara比西纳拉() 	feud.Barony
    BDorgali多尔加利() 	feud.Barony
    BGaltelli加尔泰利() 	feud.Barony
    BLungone伦贡() 	feud.Barony
    BNuoro努奥罗() 	feud.Barony
    BOlbia奥尔比亚() 	feud.Barony
    BPosada波萨达() 	feud.Barony
}

type 加卢拉GalluraCounty struct {
	feud.BaseCounty
	阿朱斯Aggius 	feud.Barony
	比西纳拉Bicinara 	feud.Barony
	多尔加利Dorgali 	feud.Barony
	加尔泰利Galtelli 	feud.Barony
	伦贡Lungone 	feud.Barony
	努奥罗Nuoro 	feud.Barony
	奥尔比亚Olbia 	feud.Barony
	波萨达Posada 	feud.Barony
}

func (c *加卢拉GalluraCounty) BAggius阿朱斯() feud.Barony {
	return c.阿朱斯Aggius
}
    
func (c *加卢拉GalluraCounty) BBicinara比西纳拉() feud.Barony {
	return c.比西纳拉Bicinara
}
    
func (c *加卢拉GalluraCounty) BDorgali多尔加利() feud.Barony {
	return c.多尔加利Dorgali
}
    
func (c *加卢拉GalluraCounty) BGaltelli加尔泰利() feud.Barony {
	return c.加尔泰利Galtelli
}
    
func (c *加卢拉GalluraCounty) BLungone伦贡() feud.Barony {
	return c.伦贡Lungone
}
    
func (c *加卢拉GalluraCounty) BNuoro努奥罗() feud.Barony {
	return c.努奥罗Nuoro
}
    
func (c *加卢拉GalluraCounty) BOlbia奥尔比亚() feud.Barony {
	return c.奥尔比亚Olbia
}
    
func (c *加卢拉GalluraCounty) BPosada波萨达() feud.Barony {
	return c.波萨达Posada
}
    
var CGallura加卢拉 GalluraCounty = &加卢拉GalluraCounty{}

func init() {
	f := CGallura加卢拉.(*加卢拉GalluraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1574",
		Title:     "gallura",
		TitleName: "加卢拉",
		TitleCode: "c_gallura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿朱斯Aggius = BAggius阿朱斯
	f.阿朱斯Aggius.SetParent(f)
	
	f.比西纳拉Bicinara = BBicinara比西纳拉
	f.比西纳拉Bicinara.SetParent(f)
	
	f.多尔加利Dorgali = BDorgali多尔加利
	f.多尔加利Dorgali.SetParent(f)
	
	f.加尔泰利Galtelli = BGaltelli加尔泰利
	f.加尔泰利Galtelli.SetParent(f)
	
	f.伦贡Lungone = BLungone伦贡
	f.伦贡Lungone.SetParent(f)
	
	f.努奥罗Nuoro = BNuoro努奥罗
	f.努奥罗Nuoro.SetParent(f)
	
	f.奥尔比亚Olbia = BOlbia奥尔比亚
	f.奥尔比亚Olbia.SetParent(f)
	
	f.波萨达Posada = BPosada波萨达
	f.波萨达Posada.SetParent(f)
	
}
