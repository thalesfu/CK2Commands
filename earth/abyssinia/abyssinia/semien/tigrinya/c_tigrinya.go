package tigrinya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TigrinyaCounty interface {
    feud.County
    BAdenna阿丹那() 	feud.Barony
    BAdi_ramets阿迪拉梅兹() 	feud.Barony
    BAmba_maderiya阿姆巴马德里亚() 	feud.Barony
    BBilamba比拉巴() 	feud.Barony
    BGolonco哥隆可() 	feud.Barony
    BTigrinya提格利尼亚() 	feud.Barony
    BZagamat札加马特() 	feud.Barony
}

type 提格利尼亚TigrinyaCounty struct {
	feud.BaseCounty
	阿丹那Adenna 	feud.Barony
	阿迪拉梅兹Adi_ramets 	feud.Barony
	阿姆巴马德里亚Amba_maderiya 	feud.Barony
	比拉巴Bilamba 	feud.Barony
	哥隆可Golonco 	feud.Barony
	提格利尼亚Tigrinya 	feud.Barony
	札加马特Zagamat 	feud.Barony
}

func (c *提格利尼亚TigrinyaCounty) BAdenna阿丹那() feud.Barony {
	return c.阿丹那Adenna
}
    
func (c *提格利尼亚TigrinyaCounty) BAdi_ramets阿迪拉梅兹() feud.Barony {
	return c.阿迪拉梅兹Adi_ramets
}
    
func (c *提格利尼亚TigrinyaCounty) BAmba_maderiya阿姆巴马德里亚() feud.Barony {
	return c.阿姆巴马德里亚Amba_maderiya
}
    
func (c *提格利尼亚TigrinyaCounty) BBilamba比拉巴() feud.Barony {
	return c.比拉巴Bilamba
}
    
func (c *提格利尼亚TigrinyaCounty) BGolonco哥隆可() feud.Barony {
	return c.哥隆可Golonco
}
    
func (c *提格利尼亚TigrinyaCounty) BTigrinya提格利尼亚() feud.Barony {
	return c.提格利尼亚Tigrinya
}
    
func (c *提格利尼亚TigrinyaCounty) BZagamat札加马特() feud.Barony {
	return c.札加马特Zagamat
}
    
var CTigrinya提格利尼亚 TigrinyaCounty = &提格利尼亚TigrinyaCounty{}

func init() {
	f := CTigrinya提格利尼亚.(*提格利尼亚TigrinyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1437",
		Title:     "tigrinya",
		TitleName: "提格利尼亚",
		TitleCode: "c_tigrinya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿丹那Adenna = BAdenna阿丹那
	f.阿丹那Adenna.SetParent(f)
	
	f.阿迪拉梅兹Adi_ramets = BAdi_ramets阿迪拉梅兹
	f.阿迪拉梅兹Adi_ramets.SetParent(f)
	
	f.阿姆巴马德里亚Amba_maderiya = BAmba_maderiya阿姆巴马德里亚
	f.阿姆巴马德里亚Amba_maderiya.SetParent(f)
	
	f.比拉巴Bilamba = BBilamba比拉巴
	f.比拉巴Bilamba.SetParent(f)
	
	f.哥隆可Golonco = BGolonco哥隆可
	f.哥隆可Golonco.SetParent(f)
	
	f.提格利尼亚Tigrinya = BTigrinya提格利尼亚
	f.提格利尼亚Tigrinya.SetParent(f)
	
	f.札加马特Zagamat = BZagamat札加马特
	f.札加马特Zagamat.SetParent(f)
	
}
