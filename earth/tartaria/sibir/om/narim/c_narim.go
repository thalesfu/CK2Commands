package narim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NarimCounty interface {
    feud.County
    BAruz阿鲁兹() 	feud.Barony
    BAskiz阿斯基兹() 	feud.Barony
    BAziz阿齐兹() 	feud.Barony
    BKolta科尔塔() 	feud.Barony
    BNarim纳林() 	feud.Barony
    BUros乌罗斯() 	feud.Barony
    BYag雅格() 	feud.Barony
}

type 纳林NarimCounty struct {
	feud.BaseCounty
	阿鲁兹Aruz 	feud.Barony
	阿斯基兹Askiz 	feud.Barony
	阿齐兹Aziz 	feud.Barony
	科尔塔Kolta 	feud.Barony
	纳林Narim 	feud.Barony
	乌罗斯Uros 	feud.Barony
	雅格Yag 	feud.Barony
}

func (c *纳林NarimCounty) BAruz阿鲁兹() feud.Barony {
	return c.阿鲁兹Aruz
}
    
func (c *纳林NarimCounty) BAskiz阿斯基兹() feud.Barony {
	return c.阿斯基兹Askiz
}
    
func (c *纳林NarimCounty) BAziz阿齐兹() feud.Barony {
	return c.阿齐兹Aziz
}
    
func (c *纳林NarimCounty) BKolta科尔塔() feud.Barony {
	return c.科尔塔Kolta
}
    
func (c *纳林NarimCounty) BNarim纳林() feud.Barony {
	return c.纳林Narim
}
    
func (c *纳林NarimCounty) BUros乌罗斯() feud.Barony {
	return c.乌罗斯Uros
}
    
func (c *纳林NarimCounty) BYag雅格() feud.Barony {
	return c.雅格Yag
}
    
var CNarim纳林 NarimCounty = &纳林NarimCounty{}

func init() {
	f := CNarim纳林.(*纳林NarimCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1205",
		Title:     "narim",
		TitleName: "纳林",
		TitleCode: "c_narim",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿鲁兹Aruz = BAruz阿鲁兹
	f.阿鲁兹Aruz.SetParent(f)
	
	f.阿斯基兹Askiz = BAskiz阿斯基兹
	f.阿斯基兹Askiz.SetParent(f)
	
	f.阿齐兹Aziz = BAziz阿齐兹
	f.阿齐兹Aziz.SetParent(f)
	
	f.科尔塔Kolta = BKolta科尔塔
	f.科尔塔Kolta.SetParent(f)
	
	f.纳林Narim = BNarim纳林
	f.纳林Narim.SetParent(f)
	
	f.乌罗斯Uros = BUros乌罗斯
	f.乌罗斯Uros.SetParent(f)
	
	f.雅格Yag = BYag雅格
	f.雅格Yag.SetParent(f)
	
}
