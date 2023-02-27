package dunkheger

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DunkhegerCounty interface {
    feud.County
    BDawuqiao大坞桥() 	feud.Barony
    BDunkheger冬和格尔() 	feud.Barony
    BHengling横岭() 	feud.Barony
    BHou_pulei后蒲类() 	feud.Barony
    BMajiajie马家街() 	feud.Barony
    BMori木垒() 	feud.Barony
    BYulishi郁立师() 	feud.Barony
}

type 冬和格尔DunkhegerCounty struct {
	feud.BaseCounty
	大坞桥Dawuqiao 	feud.Barony
	冬和格尔Dunkheger 	feud.Barony
	横岭Hengling 	feud.Barony
	后蒲类Hou_pulei 	feud.Barony
	马家街Majiajie 	feud.Barony
	木垒Mori 	feud.Barony
	郁立师Yulishi 	feud.Barony
}

func (c *冬和格尔DunkhegerCounty) BDawuqiao大坞桥() feud.Barony {
	return c.大坞桥Dawuqiao
}
    
func (c *冬和格尔DunkhegerCounty) BDunkheger冬和格尔() feud.Barony {
	return c.冬和格尔Dunkheger
}
    
func (c *冬和格尔DunkhegerCounty) BHengling横岭() feud.Barony {
	return c.横岭Hengling
}
    
func (c *冬和格尔DunkhegerCounty) BHou_pulei后蒲类() feud.Barony {
	return c.后蒲类Hou_pulei
}
    
func (c *冬和格尔DunkhegerCounty) BMajiajie马家街() feud.Barony {
	return c.马家街Majiajie
}
    
func (c *冬和格尔DunkhegerCounty) BMori木垒() feud.Barony {
	return c.木垒Mori
}
    
func (c *冬和格尔DunkhegerCounty) BYulishi郁立师() feud.Barony {
	return c.郁立师Yulishi
}
    
var CDunkheger冬和格尔 DunkhegerCounty = &冬和格尔DunkhegerCounty{}

func init() {
	f := CDunkheger冬和格尔.(*冬和格尔DunkhegerCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1517",
		Title:     "dunkheger",
		TitleName: "冬和格尔",
		TitleCode: "c_dunkheger",
		Baronies:  map[string]feud.Barony{},
	}

	f.大坞桥Dawuqiao = BDawuqiao大坞桥
	f.大坞桥Dawuqiao.SetParent(f)
	
	f.冬和格尔Dunkheger = BDunkheger冬和格尔
	f.冬和格尔Dunkheger.SetParent(f)
	
	f.横岭Hengling = BHengling横岭
	f.横岭Hengling.SetParent(f)
	
	f.后蒲类Hou_pulei = BHou_pulei后蒲类
	f.后蒲类Hou_pulei.SetParent(f)
	
	f.马家街Majiajie = BMajiajie马家街
	f.马家街Majiajie.SetParent(f)
	
	f.木垒Mori = BMori木垒
	f.木垒Mori.SetParent(f)
	
	f.郁立师Yulishi = BYulishi郁立师
	f.郁立师Yulishi.SetParent(f)
	
}
