package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TalasCounty interface {
    feud.County
    BAkchulak阿克丘拉克() 	feud.Barony
    BAsa阿萨() 	feud.Barony
    BMerke灭尔基() 	feud.Barony
    BMyrzatay梅尔扎泰() 	feud.Barony
    BSheker舍克尔() 	feud.Barony
    BShelji希勒吉() 	feud.Barony
    BTaraz怛逻斯() 	feud.Barony
}

type 怛罗斯TalasCounty struct {
	feud.BaseCounty
	阿克丘拉克Akchulak 	feud.Barony
	阿萨Asa 	feud.Barony
	灭尔基Merke 	feud.Barony
	梅尔扎泰Myrzatay 	feud.Barony
	舍克尔Sheker 	feud.Barony
	希勒吉Shelji 	feud.Barony
	怛逻斯Taraz 	feud.Barony
}

func (c *怛罗斯TalasCounty) BAkchulak阿克丘拉克() feud.Barony {
	return c.阿克丘拉克Akchulak
}
    
func (c *怛罗斯TalasCounty) BAsa阿萨() feud.Barony {
	return c.阿萨Asa
}
    
func (c *怛罗斯TalasCounty) BMerke灭尔基() feud.Barony {
	return c.灭尔基Merke
}
    
func (c *怛罗斯TalasCounty) BMyrzatay梅尔扎泰() feud.Barony {
	return c.梅尔扎泰Myrzatay
}
    
func (c *怛罗斯TalasCounty) BSheker舍克尔() feud.Barony {
	return c.舍克尔Sheker
}
    
func (c *怛罗斯TalasCounty) BShelji希勒吉() feud.Barony {
	return c.希勒吉Shelji
}
    
func (c *怛罗斯TalasCounty) BTaraz怛逻斯() feud.Barony {
	return c.怛逻斯Taraz
}
    
var CTalas怛罗斯 TalasCounty = &怛罗斯TalasCounty{}

func init() {
	f := CTalas怛罗斯.(*怛罗斯TalasCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1795",
		Title:     "talas",
		TitleName: "怛罗斯",
		TitleCode: "c_talas",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克丘拉克Akchulak = BAkchulak阿克丘拉克
	f.阿克丘拉克Akchulak.SetParent(f)
	
	f.阿萨Asa = BAsa阿萨
	f.阿萨Asa.SetParent(f)
	
	f.灭尔基Merke = BMerke灭尔基
	f.灭尔基Merke.SetParent(f)
	
	f.梅尔扎泰Myrzatay = BMyrzatay梅尔扎泰
	f.梅尔扎泰Myrzatay.SetParent(f)
	
	f.舍克尔Sheker = BSheker舍克尔
	f.舍克尔Sheker.SetParent(f)
	
	f.希勒吉Shelji = BShelji希勒吉
	f.希勒吉Shelji.SetParent(f)
	
	f.怛逻斯Taraz = BTaraz怛逻斯
	f.怛逻斯Taraz.SetParent(f)
	
}
