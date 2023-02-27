package deltuva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeltuvaCounty interface {
    feud.County
    BAriogala阿廖加拉() 	feud.Barony
    BBetygala贝蒂加拉() 	feud.Barony
    BDeltuva德尔图瓦() 	feud.Barony
    BJunigeda尤尼格达() 	feud.Barony
    BSiauliai希奥利艾() 	feud.Barony
    BUpyte乌皮特() 	feud.Barony
    BVelluona韦卢奥纳() 	feud.Barony
}

type 德尔图瓦DeltuvaCounty struct {
	feud.BaseCounty
	阿廖加拉Ariogala 	feud.Barony
	贝蒂加拉Betygala 	feud.Barony
	德尔图瓦Deltuva 	feud.Barony
	尤尼格达Junigeda 	feud.Barony
	希奥利艾Siauliai 	feud.Barony
	乌皮特Upyte 	feud.Barony
	韦卢奥纳Velluona 	feud.Barony
}

func (c *德尔图瓦DeltuvaCounty) BAriogala阿廖加拉() feud.Barony {
	return c.阿廖加拉Ariogala
}
    
func (c *德尔图瓦DeltuvaCounty) BBetygala贝蒂加拉() feud.Barony {
	return c.贝蒂加拉Betygala
}
    
func (c *德尔图瓦DeltuvaCounty) BDeltuva德尔图瓦() feud.Barony {
	return c.德尔图瓦Deltuva
}
    
func (c *德尔图瓦DeltuvaCounty) BJunigeda尤尼格达() feud.Barony {
	return c.尤尼格达Junigeda
}
    
func (c *德尔图瓦DeltuvaCounty) BSiauliai希奥利艾() feud.Barony {
	return c.希奥利艾Siauliai
}
    
func (c *德尔图瓦DeltuvaCounty) BUpyte乌皮特() feud.Barony {
	return c.乌皮特Upyte
}
    
func (c *德尔图瓦DeltuvaCounty) BVelluona韦卢奥纳() feud.Barony {
	return c.韦卢奥纳Velluona
}
    
var CDeltuva德尔图瓦 DeltuvaCounty = &德尔图瓦DeltuvaCounty{}

func init() {
	f := CDeltuva德尔图瓦.(*德尔图瓦DeltuvaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1597",
		Title:     "deltuva",
		TitleName: "德尔图瓦",
		TitleCode: "c_deltuva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿廖加拉Ariogala = BAriogala阿廖加拉
	f.阿廖加拉Ariogala.SetParent(f)
	
	f.贝蒂加拉Betygala = BBetygala贝蒂加拉
	f.贝蒂加拉Betygala.SetParent(f)
	
	f.德尔图瓦Deltuva = BDeltuva德尔图瓦
	f.德尔图瓦Deltuva.SetParent(f)
	
	f.尤尼格达Junigeda = BJunigeda尤尼格达
	f.尤尼格达Junigeda.SetParent(f)
	
	f.希奥利艾Siauliai = BSiauliai希奥利艾
	f.希奥利艾Siauliai.SetParent(f)
	
	f.乌皮特Upyte = BUpyte乌皮特
	f.乌皮特Upyte.SetParent(f)
	
	f.韦卢奥纳Velluona = BVelluona韦卢奥纳
	f.韦卢奥纳Velluona.SetParent(f)
	
}
