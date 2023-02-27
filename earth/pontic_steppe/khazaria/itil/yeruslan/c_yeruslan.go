package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YeruslanCounty interface {
    feud.County
    BBorsi博尔西() 	feud.Barony
    BImash伊玛什() 	feud.Barony
    BMereke梅列克() 	feud.Barony
    BOrda奥尔达() 	feud.Barony
    BTaygara泰加拉() 	feud.Barony
    BTungysh通格什() 	feud.Barony
    BYeruslan叶鲁斯兰() 	feud.Barony
}

type 叶鲁斯兰YeruslanCounty struct {
	feud.BaseCounty
	博尔西Borsi 	feud.Barony
	伊玛什Imash 	feud.Barony
	梅列克Mereke 	feud.Barony
	奥尔达Orda 	feud.Barony
	泰加拉Taygara 	feud.Barony
	通格什Tungysh 	feud.Barony
	叶鲁斯兰Yeruslan 	feud.Barony
}

func (c *叶鲁斯兰YeruslanCounty) BBorsi博尔西() feud.Barony {
	return c.博尔西Borsi
}
    
func (c *叶鲁斯兰YeruslanCounty) BImash伊玛什() feud.Barony {
	return c.伊玛什Imash
}
    
func (c *叶鲁斯兰YeruslanCounty) BMereke梅列克() feud.Barony {
	return c.梅列克Mereke
}
    
func (c *叶鲁斯兰YeruslanCounty) BOrda奥尔达() feud.Barony {
	return c.奥尔达Orda
}
    
func (c *叶鲁斯兰YeruslanCounty) BTaygara泰加拉() feud.Barony {
	return c.泰加拉Taygara
}
    
func (c *叶鲁斯兰YeruslanCounty) BTungysh通格什() feud.Barony {
	return c.通格什Tungysh
}
    
func (c *叶鲁斯兰YeruslanCounty) BYeruslan叶鲁斯兰() feud.Barony {
	return c.叶鲁斯兰Yeruslan
}
    
var CYeruslan叶鲁斯兰 YeruslanCounty = &叶鲁斯兰YeruslanCounty{}

func init() {
	f := CYeruslan叶鲁斯兰.(*叶鲁斯兰YeruslanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1815",
		Title:     "yeruslan",
		TitleName: "叶鲁斯兰",
		TitleCode: "c_yeruslan",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔西Borsi = BBorsi博尔西
	f.博尔西Borsi.SetParent(f)
	
	f.伊玛什Imash = BImash伊玛什
	f.伊玛什Imash.SetParent(f)
	
	f.梅列克Mereke = BMereke梅列克
	f.梅列克Mereke.SetParent(f)
	
	f.奥尔达Orda = BOrda奥尔达
	f.奥尔达Orda.SetParent(f)
	
	f.泰加拉Taygara = BTaygara泰加拉
	f.泰加拉Taygara.SetParent(f)
	
	f.通格什Tungysh = BTungysh通格什
	f.通格什Tungysh.SetParent(f)
	
	f.叶鲁斯兰Yeruslan = BYeruslan叶鲁斯兰
	f.叶鲁斯兰Yeruslan.SetParent(f)
	
}
