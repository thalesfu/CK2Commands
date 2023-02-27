package siena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SienaCounty interface {
    feud.County
    BColledievaldelsa埃尔萨谷口() 	feud.Barony
    BMontalcino蒙塔尔奇诺() 	feud.Barony
    BMontepulciano蒙特普尔恰诺() 	feud.Barony
    BMonteriggioni蒙特里焦尼() 	feud.Barony
    BPienza皮恩扎() 	feud.Barony
    BSangimignano圣吉米尼亚诺() 	feud.Barony
    BSiena锡耶纳() 	feud.Barony
}

type 锡耶纳SienaCounty struct {
	feud.BaseCounty
	埃尔萨谷口Colledievaldelsa 	feud.Barony
	蒙塔尔奇诺Montalcino 	feud.Barony
	蒙特普尔恰诺Montepulciano 	feud.Barony
	蒙特里焦尼Monteriggioni 	feud.Barony
	皮恩扎Pienza 	feud.Barony
	圣吉米尼亚诺Sangimignano 	feud.Barony
	锡耶纳Siena 	feud.Barony
}

func (c *锡耶纳SienaCounty) BColledievaldelsa埃尔萨谷口() feud.Barony {
	return c.埃尔萨谷口Colledievaldelsa
}
    
func (c *锡耶纳SienaCounty) BMontalcino蒙塔尔奇诺() feud.Barony {
	return c.蒙塔尔奇诺Montalcino
}
    
func (c *锡耶纳SienaCounty) BMontepulciano蒙特普尔恰诺() feud.Barony {
	return c.蒙特普尔恰诺Montepulciano
}
    
func (c *锡耶纳SienaCounty) BMonteriggioni蒙特里焦尼() feud.Barony {
	return c.蒙特里焦尼Monteriggioni
}
    
func (c *锡耶纳SienaCounty) BPienza皮恩扎() feud.Barony {
	return c.皮恩扎Pienza
}
    
func (c *锡耶纳SienaCounty) BSangimignano圣吉米尼亚诺() feud.Barony {
	return c.圣吉米尼亚诺Sangimignano
}
    
func (c *锡耶纳SienaCounty) BSiena锡耶纳() feud.Barony {
	return c.锡耶纳Siena
}
    
var CSiena锡耶纳 SienaCounty = &锡耶纳SienaCounty{}

func init() {
	f := CSiena锡耶纳.(*锡耶纳SienaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "330",
		Title:     "siena",
		TitleName: "锡耶纳",
		TitleCode: "c_siena",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃尔萨谷口Colledievaldelsa = BColledievaldelsa埃尔萨谷口
	f.埃尔萨谷口Colledievaldelsa.SetParent(f)
	
	f.蒙塔尔奇诺Montalcino = BMontalcino蒙塔尔奇诺
	f.蒙塔尔奇诺Montalcino.SetParent(f)
	
	f.蒙特普尔恰诺Montepulciano = BMontepulciano蒙特普尔恰诺
	f.蒙特普尔恰诺Montepulciano.SetParent(f)
	
	f.蒙特里焦尼Monteriggioni = BMonteriggioni蒙特里焦尼
	f.蒙特里焦尼Monteriggioni.SetParent(f)
	
	f.皮恩扎Pienza = BPienza皮恩扎
	f.皮恩扎Pienza.SetParent(f)
	
	f.圣吉米尼亚诺Sangimignano = BSangimignano圣吉米尼亚诺
	f.圣吉米尼亚诺Sangimignano.SetParent(f)
	
	f.锡耶纳Siena = BSiena锡耶纳
	f.锡耶纳Siena.SetParent(f)
	
}
