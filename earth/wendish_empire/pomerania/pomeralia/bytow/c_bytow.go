package bytow

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BytowCounty interface {
    feud.County
    BBytow贝图夫() 	feud.Barony
    BCarbero卡尔贝罗() 	feud.Barony
    BLebe莱贝() 	feud.Barony
    BLebork伦堡() 	feud.Barony
    BLupawa武帕瓦() 	feud.Barony
    BParchowy帕尔霍维() 	feud.Barony
    BSwolow斯沃沃夫() 	feud.Barony
}

type 贝图夫BytowCounty struct {
	feud.BaseCounty
	贝图夫Bytow 	feud.Barony
	卡尔贝罗Carbero 	feud.Barony
	莱贝Lebe 	feud.Barony
	伦堡Lebork 	feud.Barony
	武帕瓦Lupawa 	feud.Barony
	帕尔霍维Parchowy 	feud.Barony
	斯沃沃夫Swolow 	feud.Barony
}

func (c *贝图夫BytowCounty) BBytow贝图夫() feud.Barony {
	return c.贝图夫Bytow
}
    
func (c *贝图夫BytowCounty) BCarbero卡尔贝罗() feud.Barony {
	return c.卡尔贝罗Carbero
}
    
func (c *贝图夫BytowCounty) BLebe莱贝() feud.Barony {
	return c.莱贝Lebe
}
    
func (c *贝图夫BytowCounty) BLebork伦堡() feud.Barony {
	return c.伦堡Lebork
}
    
func (c *贝图夫BytowCounty) BLupawa武帕瓦() feud.Barony {
	return c.武帕瓦Lupawa
}
    
func (c *贝图夫BytowCounty) BParchowy帕尔霍维() feud.Barony {
	return c.帕尔霍维Parchowy
}
    
func (c *贝图夫BytowCounty) BSwolow斯沃沃夫() feud.Barony {
	return c.斯沃沃夫Swolow
}
    
var CBytow贝图夫 BytowCounty = &贝图夫BytowCounty{}

func init() {
	f := CBytow贝图夫.(*贝图夫BytowCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1587",
		Title:     "bytow",
		TitleName: "贝图夫",
		TitleCode: "c_bytow",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝图夫Bytow = BBytow贝图夫
	f.贝图夫Bytow.SetParent(f)
	
	f.卡尔贝罗Carbero = BCarbero卡尔贝罗
	f.卡尔贝罗Carbero.SetParent(f)
	
	f.莱贝Lebe = BLebe莱贝
	f.莱贝Lebe.SetParent(f)
	
	f.伦堡Lebork = BLebork伦堡
	f.伦堡Lebork.SetParent(f)
	
	f.武帕瓦Lupawa = BLupawa武帕瓦
	f.武帕瓦Lupawa.SetParent(f)
	
	f.帕尔霍维Parchowy = BParchowy帕尔霍维
	f.帕尔霍维Parchowy.SetParent(f)
	
	f.斯沃沃夫Swolow = BSwolow斯沃沃夫
	f.斯沃沃夫Swolow.SetParent(f)
	
}
