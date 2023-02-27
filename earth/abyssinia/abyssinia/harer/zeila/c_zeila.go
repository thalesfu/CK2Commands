package zeila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZeilaCounty interface {
    feud.County
    BAmud阿姆德() 	feud.Barony
    BBorama博拉马() 	feud.Barony
    BDilla迪拉() 	feud.Barony
    BHargeysa哈尔格萨() 	feud.Barony
    BJijiga吉吉加() 	feud.Barony
    BLughaya路哈亚() 	feud.Barony
    BZeila泽拉() 	feud.Barony
}

type 泽拉ZeilaCounty struct {
	feud.BaseCounty
	阿姆德Amud 	feud.Barony
	博拉马Borama 	feud.Barony
	迪拉Dilla 	feud.Barony
	哈尔格萨Hargeysa 	feud.Barony
	吉吉加Jijiga 	feud.Barony
	路哈亚Lughaya 	feud.Barony
	泽拉Zeila 	feud.Barony
}

func (c *泽拉ZeilaCounty) BAmud阿姆德() feud.Barony {
	return c.阿姆德Amud
}
    
func (c *泽拉ZeilaCounty) BBorama博拉马() feud.Barony {
	return c.博拉马Borama
}
    
func (c *泽拉ZeilaCounty) BDilla迪拉() feud.Barony {
	return c.迪拉Dilla
}
    
func (c *泽拉ZeilaCounty) BHargeysa哈尔格萨() feud.Barony {
	return c.哈尔格萨Hargeysa
}
    
func (c *泽拉ZeilaCounty) BJijiga吉吉加() feud.Barony {
	return c.吉吉加Jijiga
}
    
func (c *泽拉ZeilaCounty) BLughaya路哈亚() feud.Barony {
	return c.路哈亚Lughaya
}
    
func (c *泽拉ZeilaCounty) BZeila泽拉() feud.Barony {
	return c.泽拉Zeila
}
    
var CZeila泽拉 ZeilaCounty = &泽拉ZeilaCounty{}

func init() {
	f := CZeila泽拉.(*泽拉ZeilaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1344",
		Title:     "zeila",
		TitleName: "泽拉",
		TitleCode: "c_zeila",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆德Amud = BAmud阿姆德
	f.阿姆德Amud.SetParent(f)
	
	f.博拉马Borama = BBorama博拉马
	f.博拉马Borama.SetParent(f)
	
	f.迪拉Dilla = BDilla迪拉
	f.迪拉Dilla.SetParent(f)
	
	f.哈尔格萨Hargeysa = BHargeysa哈尔格萨
	f.哈尔格萨Hargeysa.SetParent(f)
	
	f.吉吉加Jijiga = BJijiga吉吉加
	f.吉吉加Jijiga.SetParent(f)
	
	f.路哈亚Lughaya = BLughaya路哈亚
	f.路哈亚Lughaya.SetParent(f)
	
	f.泽拉Zeila = BZeila泽拉
	f.泽拉Zeila.SetParent(f)
	
}
