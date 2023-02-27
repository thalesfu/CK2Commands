package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedjerdaCounty interface {
    feud.County
    BDougga沙格() 	feud.Barony
    BElkrib凯里卜() 	feud.Barony
    BLekef卡夫() 	feud.Barony
    BMedjerda迈杰尔达() 	feud.Barony
    BQaafur喀阿富尔() 	feud.Barony
    BSidimubarak西迪穆巴拉克() 	feud.Barony
    BSiliana锡勒亚奈() 	feud.Barony
}

type 迈杰尔达MedjerdaCounty struct {
	feud.BaseCounty
	沙格Dougga 	feud.Barony
	凯里卜Elkrib 	feud.Barony
	卡夫Lekef 	feud.Barony
	迈杰尔达Medjerda 	feud.Barony
	喀阿富尔Qaafur 	feud.Barony
	西迪穆巴拉克Sidimubarak 	feud.Barony
	锡勒亚奈Siliana 	feud.Barony
}

func (c *迈杰尔达MedjerdaCounty) BDougga沙格() feud.Barony {
	return c.沙格Dougga
}
    
func (c *迈杰尔达MedjerdaCounty) BElkrib凯里卜() feud.Barony {
	return c.凯里卜Elkrib
}
    
func (c *迈杰尔达MedjerdaCounty) BLekef卡夫() feud.Barony {
	return c.卡夫Lekef
}
    
func (c *迈杰尔达MedjerdaCounty) BMedjerda迈杰尔达() feud.Barony {
	return c.迈杰尔达Medjerda
}
    
func (c *迈杰尔达MedjerdaCounty) BQaafur喀阿富尔() feud.Barony {
	return c.喀阿富尔Qaafur
}
    
func (c *迈杰尔达MedjerdaCounty) BSidimubarak西迪穆巴拉克() feud.Barony {
	return c.西迪穆巴拉克Sidimubarak
}
    
func (c *迈杰尔达MedjerdaCounty) BSiliana锡勒亚奈() feud.Barony {
	return c.锡勒亚奈Siliana
}
    
var CMedjerda迈杰尔达 MedjerdaCounty = &迈杰尔达MedjerdaCounty{}

func init() {
	f := CMedjerda迈杰尔达.(*迈杰尔达MedjerdaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "818",
		Title:     "medjerda",
		TitleName: "迈杰尔达",
		TitleCode: "c_medjerda",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙格Dougga = BDougga沙格
	f.沙格Dougga.SetParent(f)
	
	f.凯里卜Elkrib = BElkrib凯里卜
	f.凯里卜Elkrib.SetParent(f)
	
	f.卡夫Lekef = BLekef卡夫
	f.卡夫Lekef.SetParent(f)
	
	f.迈杰尔达Medjerda = BMedjerda迈杰尔达
	f.迈杰尔达Medjerda.SetParent(f)
	
	f.喀阿富尔Qaafur = BQaafur喀阿富尔
	f.喀阿富尔Qaafur.SetParent(f)
	
	f.西迪穆巴拉克Sidimubarak = BSidimubarak西迪穆巴拉克
	f.西迪穆巴拉克Sidimubarak.SetParent(f)
	
	f.锡勒亚奈Siliana = BSiliana锡勒亚奈
	f.锡勒亚奈Siliana.SetParent(f)
	
}
