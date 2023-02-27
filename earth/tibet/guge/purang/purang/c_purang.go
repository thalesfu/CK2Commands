package purang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PurangCounty interface {
    feud.County
    BDulung多隆() 	feud.Barony
    BKailash鸡罗娑() 	feud.Barony
    BKhorzhak科迦() 	feud.Barony
    BPurang布让() 	feud.Barony
    BSimbiling贤柏林() 	feud.Barony
    BTeglakar达拉喀() 	feud.Barony
    BZhangja生结() 	feud.Barony
}

type 布让PurangCounty struct {
	feud.BaseCounty
	多隆Dulung 	feud.Barony
	鸡罗娑Kailash 	feud.Barony
	科迦Khorzhak 	feud.Barony
	布让Purang 	feud.Barony
	贤柏林Simbiling 	feud.Barony
	达拉喀Teglakar 	feud.Barony
	生结Zhangja 	feud.Barony
}

func (c *布让PurangCounty) BDulung多隆() feud.Barony {
	return c.多隆Dulung
}
    
func (c *布让PurangCounty) BKailash鸡罗娑() feud.Barony {
	return c.鸡罗娑Kailash
}
    
func (c *布让PurangCounty) BKhorzhak科迦() feud.Barony {
	return c.科迦Khorzhak
}
    
func (c *布让PurangCounty) BPurang布让() feud.Barony {
	return c.布让Purang
}
    
func (c *布让PurangCounty) BSimbiling贤柏林() feud.Barony {
	return c.贤柏林Simbiling
}
    
func (c *布让PurangCounty) BTeglakar达拉喀() feud.Barony {
	return c.达拉喀Teglakar
}
    
func (c *布让PurangCounty) BZhangja生结() feud.Barony {
	return c.生结Zhangja
}
    
var CPurang布让 PurangCounty = &布让PurangCounty{}

func init() {
	f := CPurang布让.(*布让PurangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1491",
		Title:     "purang",
		TitleName: "布让",
		TitleCode: "c_purang",
		Baronies:  map[string]feud.Barony{},
	}

	f.多隆Dulung = BDulung多隆
	f.多隆Dulung.SetParent(f)
	
	f.鸡罗娑Kailash = BKailash鸡罗娑
	f.鸡罗娑Kailash.SetParent(f)
	
	f.科迦Khorzhak = BKhorzhak科迦
	f.科迦Khorzhak.SetParent(f)
	
	f.布让Purang = BPurang布让
	f.布让Purang.SetParent(f)
	
	f.贤柏林Simbiling = BSimbiling贤柏林
	f.贤柏林Simbiling.SetParent(f)
	
	f.达拉喀Teglakar = BTeglakar达拉喀
	f.达拉喀Teglakar.SetParent(f)
	
	f.生结Zhangja = BZhangja生结
	f.生结Zhangja.SetParent(f)
	
}
