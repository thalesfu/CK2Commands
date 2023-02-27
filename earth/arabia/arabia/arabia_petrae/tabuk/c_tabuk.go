package tabuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TabukCounty interface {
    feud.County
    BAbu_ujayyijat阿布欧杰伊贾特() 	feud.Barony
    BDuba杜巴() 	feud.Barony
    BGabouk加布克() 	feud.Barony
    BMariyiat马里伊特() 	feud.Barony
    BMuwaylih穆韦利赫() 	feud.Barony
    BShaghab舍盖卜() 	feud.Barony
    BSharmah舍尔迈() 	feud.Barony
    BTabuk塔布克() 	feud.Barony
}

type 塔布克TabukCounty struct {
	feud.BaseCounty
	阿布欧杰伊贾特Abu_ujayyijat 	feud.Barony
	杜巴Duba 	feud.Barony
	加布克Gabouk 	feud.Barony
	马里伊特Mariyiat 	feud.Barony
	穆韦利赫Muwaylih 	feud.Barony
	舍盖卜Shaghab 	feud.Barony
	舍尔迈Sharmah 	feud.Barony
	塔布克Tabuk 	feud.Barony
}

func (c *塔布克TabukCounty) BAbu_ujayyijat阿布欧杰伊贾特() feud.Barony {
	return c.阿布欧杰伊贾特Abu_ujayyijat
}
    
func (c *塔布克TabukCounty) BDuba杜巴() feud.Barony {
	return c.杜巴Duba
}
    
func (c *塔布克TabukCounty) BGabouk加布克() feud.Barony {
	return c.加布克Gabouk
}
    
func (c *塔布克TabukCounty) BMariyiat马里伊特() feud.Barony {
	return c.马里伊特Mariyiat
}
    
func (c *塔布克TabukCounty) BMuwaylih穆韦利赫() feud.Barony {
	return c.穆韦利赫Muwaylih
}
    
func (c *塔布克TabukCounty) BShaghab舍盖卜() feud.Barony {
	return c.舍盖卜Shaghab
}
    
func (c *塔布克TabukCounty) BSharmah舍尔迈() feud.Barony {
	return c.舍尔迈Sharmah
}
    
func (c *塔布克TabukCounty) BTabuk塔布克() feud.Barony {
	return c.塔布克Tabuk
}
    
var CTabuk塔布克 TabukCounty = &塔布克TabukCounty{}

func init() {
	f := CTabuk塔布克.(*塔布克TabukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "721",
		Title:     "tabuk",
		TitleName: "塔布克",
		TitleCode: "c_tabuk",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布欧杰伊贾特Abu_ujayyijat = BAbu_ujayyijat阿布欧杰伊贾特
	f.阿布欧杰伊贾特Abu_ujayyijat.SetParent(f)
	
	f.杜巴Duba = BDuba杜巴
	f.杜巴Duba.SetParent(f)
	
	f.加布克Gabouk = BGabouk加布克
	f.加布克Gabouk.SetParent(f)
	
	f.马里伊特Mariyiat = BMariyiat马里伊特
	f.马里伊特Mariyiat.SetParent(f)
	
	f.穆韦利赫Muwaylih = BMuwaylih穆韦利赫
	f.穆韦利赫Muwaylih.SetParent(f)
	
	f.舍盖卜Shaghab = BShaghab舍盖卜
	f.舍盖卜Shaghab.SetParent(f)
	
	f.舍尔迈Sharmah = BSharmah舍尔迈
	f.舍尔迈Sharmah.SetParent(f)
	
	f.塔布克Tabuk = BTabuk塔布克
	f.塔布克Tabuk.SetParent(f)
	
}
