package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EuboiaCounty interface {
    feud.County
    BArtemisio阿提米修() 	feud.Barony
    BChalkis哈尔基斯() 	feud.Barony
    BIstiaia伊斯泰雅() 	feud.Barony
    BKarystos卡利斯多斯() 	feud.Barony
    BKymi奇米() 	feud.Barony
    BLilantia利兰迪亚() 	feud.Barony
    BMessapia梅萨皮亚() 	feud.Barony
    BOreoi奥利奥() 	feud.Barony
}

type 优卑亚EuboiaCounty struct {
	feud.BaseCounty
	阿提米修Artemisio 	feud.Barony
	哈尔基斯Chalkis 	feud.Barony
	伊斯泰雅Istiaia 	feud.Barony
	卡利斯多斯Karystos 	feud.Barony
	奇米Kymi 	feud.Barony
	利兰迪亚Lilantia 	feud.Barony
	梅萨皮亚Messapia 	feud.Barony
	奥利奥Oreoi 	feud.Barony
}

func (c *优卑亚EuboiaCounty) BArtemisio阿提米修() feud.Barony {
	return c.阿提米修Artemisio
}
    
func (c *优卑亚EuboiaCounty) BChalkis哈尔基斯() feud.Barony {
	return c.哈尔基斯Chalkis
}
    
func (c *优卑亚EuboiaCounty) BIstiaia伊斯泰雅() feud.Barony {
	return c.伊斯泰雅Istiaia
}
    
func (c *优卑亚EuboiaCounty) BKarystos卡利斯多斯() feud.Barony {
	return c.卡利斯多斯Karystos
}
    
func (c *优卑亚EuboiaCounty) BKymi奇米() feud.Barony {
	return c.奇米Kymi
}
    
func (c *优卑亚EuboiaCounty) BLilantia利兰迪亚() feud.Barony {
	return c.利兰迪亚Lilantia
}
    
func (c *优卑亚EuboiaCounty) BMessapia梅萨皮亚() feud.Barony {
	return c.梅萨皮亚Messapia
}
    
func (c *优卑亚EuboiaCounty) BOreoi奥利奥() feud.Barony {
	return c.奥利奥Oreoi
}
    
var CEuboia优卑亚 EuboiaCounty = &优卑亚EuboiaCounty{}

func init() {
	f := CEuboia优卑亚.(*优卑亚EuboiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "485",
		Title:     "euboia",
		TitleName: "优卑亚",
		TitleCode: "c_euboia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿提米修Artemisio = BArtemisio阿提米修
	f.阿提米修Artemisio.SetParent(f)
	
	f.哈尔基斯Chalkis = BChalkis哈尔基斯
	f.哈尔基斯Chalkis.SetParent(f)
	
	f.伊斯泰雅Istiaia = BIstiaia伊斯泰雅
	f.伊斯泰雅Istiaia.SetParent(f)
	
	f.卡利斯多斯Karystos = BKarystos卡利斯多斯
	f.卡利斯多斯Karystos.SetParent(f)
	
	f.奇米Kymi = BKymi奇米
	f.奇米Kymi.SetParent(f)
	
	f.利兰迪亚Lilantia = BLilantia利兰迪亚
	f.利兰迪亚Lilantia.SetParent(f)
	
	f.梅萨皮亚Messapia = BMessapia梅萨皮亚
	f.梅萨皮亚Messapia.SetParent(f)
	
	f.奥利奥Oreoi = BOreoi奥利奥
	f.奥利奥Oreoi.SetParent(f)
	
}
