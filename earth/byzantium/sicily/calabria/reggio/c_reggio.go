package reggio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ReggioCounty interface {
    feud.County
    BBova博瓦() 	feud.Barony
    BGerace杰拉切() 	feud.Barony
    BMileto米莱托() 	feud.Barony
    BNicotera尼科泰拉() 	feud.Barony
    BReggio雷焦() 	feud.Barony
    BSiderno锡代诺() 	feud.Barony
    BTropea特罗佩阿() 	feud.Barony
}

type 雷焦ReggioCounty struct {
	feud.BaseCounty
	博瓦Bova 	feud.Barony
	杰拉切Gerace 	feud.Barony
	米莱托Mileto 	feud.Barony
	尼科泰拉Nicotera 	feud.Barony
	雷焦Reggio 	feud.Barony
	锡代诺Siderno 	feud.Barony
	特罗佩阿Tropea 	feud.Barony
}

func (c *雷焦ReggioCounty) BBova博瓦() feud.Barony {
	return c.博瓦Bova
}
    
func (c *雷焦ReggioCounty) BGerace杰拉切() feud.Barony {
	return c.杰拉切Gerace
}
    
func (c *雷焦ReggioCounty) BMileto米莱托() feud.Barony {
	return c.米莱托Mileto
}
    
func (c *雷焦ReggioCounty) BNicotera尼科泰拉() feud.Barony {
	return c.尼科泰拉Nicotera
}
    
func (c *雷焦ReggioCounty) BReggio雷焦() feud.Barony {
	return c.雷焦Reggio
}
    
func (c *雷焦ReggioCounty) BSiderno锡代诺() feud.Barony {
	return c.锡代诺Siderno
}
    
func (c *雷焦ReggioCounty) BTropea特罗佩阿() feud.Barony {
	return c.特罗佩阿Tropea
}
    
var CReggio雷焦 ReggioCounty = &雷焦ReggioCounty{}

func init() {
	f := CReggio雷焦.(*雷焦ReggioCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "338",
		Title:     "reggio",
		TitleName: "雷焦",
		TitleCode: "c_reggio",
		Baronies:  map[string]feud.Barony{},
	}

	f.博瓦Bova = BBova博瓦
	f.博瓦Bova.SetParent(f)
	
	f.杰拉切Gerace = BGerace杰拉切
	f.杰拉切Gerace.SetParent(f)
	
	f.米莱托Mileto = BMileto米莱托
	f.米莱托Mileto.SetParent(f)
	
	f.尼科泰拉Nicotera = BNicotera尼科泰拉
	f.尼科泰拉Nicotera.SetParent(f)
	
	f.雷焦Reggio = BReggio雷焦
	f.雷焦Reggio.SetParent(f)
	
	f.锡代诺Siderno = BSiderno锡代诺
	f.锡代诺Siderno.SetParent(f)
	
	f.特罗佩阿Tropea = BTropea特罗佩阿
	f.特罗佩阿Tropea.SetParent(f)
	
}
