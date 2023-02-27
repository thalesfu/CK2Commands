package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FerraraCounty interface {
    feud.County
    BBondeno邦代诺() 	feud.Barony
    BCodigoro科迪戈罗() 	feud.Barony
    BCommacchio科马基奥() 	feud.Barony
    BCopparo科帕罗() 	feud.Barony
    BFerrara费拉拉() 	feud.Barony
    BOcchiobello奥基奥贝洛() 	feud.Barony
    BTresigallo特雷西加洛() 	feud.Barony
    BVoghiera沃吉耶拉() 	feud.Barony
}

type 费拉拉FerraraCounty struct {
	feud.BaseCounty
	邦代诺Bondeno 	feud.Barony
	科迪戈罗Codigoro 	feud.Barony
	科马基奥Commacchio 	feud.Barony
	科帕罗Copparo 	feud.Barony
	费拉拉Ferrara 	feud.Barony
	奥基奥贝洛Occhiobello 	feud.Barony
	特雷西加洛Tresigallo 	feud.Barony
	沃吉耶拉Voghiera 	feud.Barony
}

func (c *费拉拉FerraraCounty) BBondeno邦代诺() feud.Barony {
	return c.邦代诺Bondeno
}
    
func (c *费拉拉FerraraCounty) BCodigoro科迪戈罗() feud.Barony {
	return c.科迪戈罗Codigoro
}
    
func (c *费拉拉FerraraCounty) BCommacchio科马基奥() feud.Barony {
	return c.科马基奥Commacchio
}
    
func (c *费拉拉FerraraCounty) BCopparo科帕罗() feud.Barony {
	return c.科帕罗Copparo
}
    
func (c *费拉拉FerraraCounty) BFerrara费拉拉() feud.Barony {
	return c.费拉拉Ferrara
}
    
func (c *费拉拉FerraraCounty) BOcchiobello奥基奥贝洛() feud.Barony {
	return c.奥基奥贝洛Occhiobello
}
    
func (c *费拉拉FerraraCounty) BTresigallo特雷西加洛() feud.Barony {
	return c.特雷西加洛Tresigallo
}
    
func (c *费拉拉FerraraCounty) BVoghiera沃吉耶拉() feud.Barony {
	return c.沃吉耶拉Voghiera
}
    
var CFerrara费拉拉 FerraraCounty = &费拉拉FerraraCounty{}

func init() {
	f := CFerrara费拉拉.(*费拉拉FerraraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "353",
		Title:     "ferrara",
		TitleName: "费拉拉",
		TitleCode: "c_ferrara",
		Baronies:  map[string]feud.Barony{},
	}

	f.邦代诺Bondeno = BBondeno邦代诺
	f.邦代诺Bondeno.SetParent(f)
	
	f.科迪戈罗Codigoro = BCodigoro科迪戈罗
	f.科迪戈罗Codigoro.SetParent(f)
	
	f.科马基奥Commacchio = BCommacchio科马基奥
	f.科马基奥Commacchio.SetParent(f)
	
	f.科帕罗Copparo = BCopparo科帕罗
	f.科帕罗Copparo.SetParent(f)
	
	f.费拉拉Ferrara = BFerrara费拉拉
	f.费拉拉Ferrara.SetParent(f)
	
	f.奥基奥贝洛Occhiobello = BOcchiobello奥基奥贝洛
	f.奥基奥贝洛Occhiobello.SetParent(f)
	
	f.特雷西加洛Tresigallo = BTresigallo特雷西加洛
	f.特雷西加洛Tresigallo.SetParent(f)
	
	f.沃吉耶拉Voghiera = BVoghiera沃吉耶拉
	f.沃吉耶拉Voghiera.SetParent(f)
	
}
