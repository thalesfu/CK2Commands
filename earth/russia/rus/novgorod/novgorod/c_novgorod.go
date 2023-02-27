package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NovgorodCounty interface {
    feud.County
    BBatetsky巴捷茨基() 	feud.Barony
    BChudovo丘多沃() 	feud.Barony
    BLuga卢加() 	feud.Barony
    BNovgorod诺夫哥罗德() 	feud.Barony
    BPankovka潘科夫卡() 	feud.Barony
    BShimsk希姆斯克() 	feud.Barony
    BSoltsy索利齐() 	feud.Barony
}

type 诺夫哥罗德NovgorodCounty struct {
	feud.BaseCounty
	巴捷茨基Batetsky 	feud.Barony
	丘多沃Chudovo 	feud.Barony
	卢加Luga 	feud.Barony
	诺夫哥罗德Novgorod 	feud.Barony
	潘科夫卡Pankovka 	feud.Barony
	希姆斯克Shimsk 	feud.Barony
	索利齐Soltsy 	feud.Barony
}

func (c *诺夫哥罗德NovgorodCounty) BBatetsky巴捷茨基() feud.Barony {
	return c.巴捷茨基Batetsky
}
    
func (c *诺夫哥罗德NovgorodCounty) BChudovo丘多沃() feud.Barony {
	return c.丘多沃Chudovo
}
    
func (c *诺夫哥罗德NovgorodCounty) BLuga卢加() feud.Barony {
	return c.卢加Luga
}
    
func (c *诺夫哥罗德NovgorodCounty) BNovgorod诺夫哥罗德() feud.Barony {
	return c.诺夫哥罗德Novgorod
}
    
func (c *诺夫哥罗德NovgorodCounty) BPankovka潘科夫卡() feud.Barony {
	return c.潘科夫卡Pankovka
}
    
func (c *诺夫哥罗德NovgorodCounty) BShimsk希姆斯克() feud.Barony {
	return c.希姆斯克Shimsk
}
    
func (c *诺夫哥罗德NovgorodCounty) BSoltsy索利齐() feud.Barony {
	return c.索利齐Soltsy
}
    
var CNovgorod诺夫哥罗德 NovgorodCounty = &诺夫哥罗德NovgorodCounty{}

func init() {
	f := CNovgorod诺夫哥罗德.(*诺夫哥罗德NovgorodCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "414",
		Title:     "novgorod",
		TitleName: "诺夫哥罗德",
		TitleCode: "c_novgorod",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴捷茨基Batetsky = BBatetsky巴捷茨基
	f.巴捷茨基Batetsky.SetParent(f)
	
	f.丘多沃Chudovo = BChudovo丘多沃
	f.丘多沃Chudovo.SetParent(f)
	
	f.卢加Luga = BLuga卢加
	f.卢加Luga.SetParent(f)
	
	f.诺夫哥罗德Novgorod = BNovgorod诺夫哥罗德
	f.诺夫哥罗德Novgorod.SetParent(f)
	
	f.潘科夫卡Pankovka = BPankovka潘科夫卡
	f.潘科夫卡Pankovka.SetParent(f)
	
	f.希姆斯克Shimsk = BShimsk希姆斯克
	f.希姆斯克Shimsk.SetParent(f)
	
	f.索利齐Soltsy = BSoltsy索利齐
	f.索利齐Soltsy.SetParent(f)
	
}
