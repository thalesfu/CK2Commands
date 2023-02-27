package nobatia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NobatiaCounty interface {
    feud.County
    BAbri阿卜里() 	feud.Barony
    BAkasha阿卡沙() 	feud.Barony
    BDelgo代勒古() 	feud.Barony
    BKulubnarti库勒卜() 	feud.Barony
    BSemna塞姆纳() 	feud.Barony
    BSoleb叟勒布() 	feud.Barony
    BWawa瓦瓦() 	feud.Barony
}

type 诺巴提亚NobatiaCounty struct {
	feud.BaseCounty
	阿卜里Abri 	feud.Barony
	阿卡沙Akasha 	feud.Barony
	代勒古Delgo 	feud.Barony
	库勒卜Kulubnarti 	feud.Barony
	塞姆纳Semna 	feud.Barony
	叟勒布Soleb 	feud.Barony
	瓦瓦Wawa 	feud.Barony
}

func (c *诺巴提亚NobatiaCounty) BAbri阿卜里() feud.Barony {
	return c.阿卜里Abri
}
    
func (c *诺巴提亚NobatiaCounty) BAkasha阿卡沙() feud.Barony {
	return c.阿卡沙Akasha
}
    
func (c *诺巴提亚NobatiaCounty) BDelgo代勒古() feud.Barony {
	return c.代勒古Delgo
}
    
func (c *诺巴提亚NobatiaCounty) BKulubnarti库勒卜() feud.Barony {
	return c.库勒卜Kulubnarti
}
    
func (c *诺巴提亚NobatiaCounty) BSemna塞姆纳() feud.Barony {
	return c.塞姆纳Semna
}
    
func (c *诺巴提亚NobatiaCounty) BSoleb叟勒布() feud.Barony {
	return c.叟勒布Soleb
}
    
func (c *诺巴提亚NobatiaCounty) BWawa瓦瓦() feud.Barony {
	return c.瓦瓦Wawa
}
    
var CNobatia诺巴提亚 NobatiaCounty = &诺巴提亚NobatiaCounty{}

func init() {
	f := CNobatia诺巴提亚.(*诺巴提亚NobatiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1322",
		Title:     "nobatia",
		TitleName: "诺巴提亚",
		TitleCode: "c_nobatia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿卜里Abri = BAbri阿卜里
	f.阿卜里Abri.SetParent(f)
	
	f.阿卡沙Akasha = BAkasha阿卡沙
	f.阿卡沙Akasha.SetParent(f)
	
	f.代勒古Delgo = BDelgo代勒古
	f.代勒古Delgo.SetParent(f)
	
	f.库勒卜Kulubnarti = BKulubnarti库勒卜
	f.库勒卜Kulubnarti.SetParent(f)
	
	f.塞姆纳Semna = BSemna塞姆纳
	f.塞姆纳Semna.SetParent(f)
	
	f.叟勒布Soleb = BSoleb叟勒布
	f.叟勒布Soleb.SetParent(f)
	
	f.瓦瓦Wawa = BWawa瓦瓦
	f.瓦瓦Wawa.SetParent(f)
	
}
