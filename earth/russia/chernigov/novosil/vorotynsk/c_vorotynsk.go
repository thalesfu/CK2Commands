package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VorotynskCounty interface {
    feud.County
    BChern切尔尼() 	feud.Barony
    BMchensk姆岑斯克() 	feud.Barony
    BOdoyev奥多耶夫() 	feud.Barony
    BPlavsk普拉夫斯克() 	feud.Barony
    BSomovo索莫沃() 	feud.Barony
    BSuvorova苏沃罗瓦() 	feud.Barony
    BVorotynsk沃罗滕斯克() 	feud.Barony
}

type 沃罗滕斯克VorotynskCounty struct {
	feud.BaseCounty
	切尔尼Chern 	feud.Barony
	姆岑斯克Mchensk 	feud.Barony
	奥多耶夫Odoyev 	feud.Barony
	普拉夫斯克Plavsk 	feud.Barony
	索莫沃Somovo 	feud.Barony
	苏沃罗瓦Suvorova 	feud.Barony
	沃罗滕斯克Vorotynsk 	feud.Barony
}

func (c *沃罗滕斯克VorotynskCounty) BChern切尔尼() feud.Barony {
	return c.切尔尼Chern
}
    
func (c *沃罗滕斯克VorotynskCounty) BMchensk姆岑斯克() feud.Barony {
	return c.姆岑斯克Mchensk
}
    
func (c *沃罗滕斯克VorotynskCounty) BOdoyev奥多耶夫() feud.Barony {
	return c.奥多耶夫Odoyev
}
    
func (c *沃罗滕斯克VorotynskCounty) BPlavsk普拉夫斯克() feud.Barony {
	return c.普拉夫斯克Plavsk
}
    
func (c *沃罗滕斯克VorotynskCounty) BSomovo索莫沃() feud.Barony {
	return c.索莫沃Somovo
}
    
func (c *沃罗滕斯克VorotynskCounty) BSuvorova苏沃罗瓦() feud.Barony {
	return c.苏沃罗瓦Suvorova
}
    
func (c *沃罗滕斯克VorotynskCounty) BVorotynsk沃罗滕斯克() feud.Barony {
	return c.沃罗滕斯克Vorotynsk
}
    
var CVorotynsk沃罗滕斯克 VorotynskCounty = &沃罗滕斯克VorotynskCounty{}

func init() {
	f := CVorotynsk沃罗滕斯克.(*沃罗滕斯克VorotynskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1650",
		Title:     "vorotynsk",
		TitleName: "沃罗滕斯克",
		TitleCode: "c_vorotynsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔尼Chern = BChern切尔尼
	f.切尔尼Chern.SetParent(f)
	
	f.姆岑斯克Mchensk = BMchensk姆岑斯克
	f.姆岑斯克Mchensk.SetParent(f)
	
	f.奥多耶夫Odoyev = BOdoyev奥多耶夫
	f.奥多耶夫Odoyev.SetParent(f)
	
	f.普拉夫斯克Plavsk = BPlavsk普拉夫斯克
	f.普拉夫斯克Plavsk.SetParent(f)
	
	f.索莫沃Somovo = BSomovo索莫沃
	f.索莫沃Somovo.SetParent(f)
	
	f.苏沃罗瓦Suvorova = BSuvorova苏沃罗瓦
	f.苏沃罗瓦Suvorova.SetParent(f)
	
	f.沃罗滕斯克Vorotynsk = BVorotynsk沃罗滕斯克
	f.沃罗滕斯克Vorotynsk.SetParent(f)
	
}
