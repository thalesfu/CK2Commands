package polotsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PolotskCounty interface {
    feud.County
    BDouza多尔扎() 	feud.Barony
    BHaradok戈罗多克() 	feud.Barony
    BLeskovichi列斯科维奇() 	feud.Barony
    BObol奥博利() 	feud.Barony
    BPolotsk波洛茨克() 	feud.Barony
    BRassony罗索内() 	feud.Barony
    BSirotino西罗季诺() 	feud.Barony
}

type 波洛茨克PolotskCounty struct {
	feud.BaseCounty
	多尔扎Douza 	feud.Barony
	戈罗多克Haradok 	feud.Barony
	列斯科维奇Leskovichi 	feud.Barony
	奥博利Obol 	feud.Barony
	波洛茨克Polotsk 	feud.Barony
	罗索内Rassony 	feud.Barony
	西罗季诺Sirotino 	feud.Barony
}

func (c *波洛茨克PolotskCounty) BDouza多尔扎() feud.Barony {
	return c.多尔扎Douza
}
    
func (c *波洛茨克PolotskCounty) BHaradok戈罗多克() feud.Barony {
	return c.戈罗多克Haradok
}
    
func (c *波洛茨克PolotskCounty) BLeskovichi列斯科维奇() feud.Barony {
	return c.列斯科维奇Leskovichi
}
    
func (c *波洛茨克PolotskCounty) BObol奥博利() feud.Barony {
	return c.奥博利Obol
}
    
func (c *波洛茨克PolotskCounty) BPolotsk波洛茨克() feud.Barony {
	return c.波洛茨克Polotsk
}
    
func (c *波洛茨克PolotskCounty) BRassony罗索内() feud.Barony {
	return c.罗索内Rassony
}
    
func (c *波洛茨克PolotskCounty) BSirotino西罗季诺() feud.Barony {
	return c.西罗季诺Sirotino
}
    
var CPolotsk波洛茨克 PolotskCounty = &波洛茨克PolotskCounty{}

func init() {
	f := CPolotsk波洛茨克.(*波洛茨克PolotskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "419",
		Title:     "polotsk",
		TitleName: "波洛茨克",
		TitleCode: "c_polotsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.多尔扎Douza = BDouza多尔扎
	f.多尔扎Douza.SetParent(f)
	
	f.戈罗多克Haradok = BHaradok戈罗多克
	f.戈罗多克Haradok.SetParent(f)
	
	f.列斯科维奇Leskovichi = BLeskovichi列斯科维奇
	f.列斯科维奇Leskovichi.SetParent(f)
	
	f.奥博利Obol = BObol奥博利
	f.奥博利Obol.SetParent(f)
	
	f.波洛茨克Polotsk = BPolotsk波洛茨克
	f.波洛茨克Polotsk.SetParent(f)
	
	f.罗索内Rassony = BRassony罗索内
	f.罗索内Rassony.SetParent(f)
	
	f.西罗季诺Sirotino = BSirotino西罗季诺
	f.西罗季诺Sirotino.SetParent(f)
	
}
