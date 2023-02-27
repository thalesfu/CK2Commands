package al_bichri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_bichriCounty interface {
    feud.County
    BAbukamal阿布卡迈勒() 	feud.Barony
    BBichri比基利() 	feud.Barony
    BDeiralzur代尔祖尔() 	feud.Barony
    BMayadin迈亚丁() 	feud.Barony
    BMhaymidah穆哈米达() 	feud.Barony
    BOsrhoene奥斯霍尼() 	feud.Barony
    BShamiyya沙米亚() 	feud.Barony
    BSirhi锡尔() 	feud.Barony
}

type 比什里Al_bichriCounty struct {
	feud.BaseCounty
	阿布卡迈勒Abukamal 	feud.Barony
	比基利Bichri 	feud.Barony
	代尔祖尔Deiralzur 	feud.Barony
	迈亚丁Mayadin 	feud.Barony
	穆哈米达Mhaymidah 	feud.Barony
	奥斯霍尼Osrhoene 	feud.Barony
	沙米亚Shamiyya 	feud.Barony
	锡尔Sirhi 	feud.Barony
}

func (c *比什里Al_bichriCounty) BAbukamal阿布卡迈勒() feud.Barony {
	return c.阿布卡迈勒Abukamal
}
    
func (c *比什里Al_bichriCounty) BBichri比基利() feud.Barony {
	return c.比基利Bichri
}
    
func (c *比什里Al_bichriCounty) BDeiralzur代尔祖尔() feud.Barony {
	return c.代尔祖尔Deiralzur
}
    
func (c *比什里Al_bichriCounty) BMayadin迈亚丁() feud.Barony {
	return c.迈亚丁Mayadin
}
    
func (c *比什里Al_bichriCounty) BMhaymidah穆哈米达() feud.Barony {
	return c.穆哈米达Mhaymidah
}
    
func (c *比什里Al_bichriCounty) BOsrhoene奥斯霍尼() feud.Barony {
	return c.奥斯霍尼Osrhoene
}
    
func (c *比什里Al_bichriCounty) BShamiyya沙米亚() feud.Barony {
	return c.沙米亚Shamiyya
}
    
func (c *比什里Al_bichriCounty) BSirhi锡尔() feud.Barony {
	return c.锡尔Sirhi
}
    
var CAl_bichri比什里 Al_bichriCounty = &比什里Al_bichriCounty{}

func init() {
	f := CAl_bichri比什里.(*比什里Al_bichriCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "710",
		Title:     "al_bichri",
		TitleName: "比什里",
		TitleCode: "c_al_bichri",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布卡迈勒Abukamal = BAbukamal阿布卡迈勒
	f.阿布卡迈勒Abukamal.SetParent(f)
	
	f.比基利Bichri = BBichri比基利
	f.比基利Bichri.SetParent(f)
	
	f.代尔祖尔Deiralzur = BDeiralzur代尔祖尔
	f.代尔祖尔Deiralzur.SetParent(f)
	
	f.迈亚丁Mayadin = BMayadin迈亚丁
	f.迈亚丁Mayadin.SetParent(f)
	
	f.穆哈米达Mhaymidah = BMhaymidah穆哈米达
	f.穆哈米达Mhaymidah.SetParent(f)
	
	f.奥斯霍尼Osrhoene = BOsrhoene奥斯霍尼
	f.奥斯霍尼Osrhoene.SetParent(f)
	
	f.沙米亚Shamiyya = BShamiyya沙米亚
	f.沙米亚Shamiyya.SetParent(f)
	
	f.锡尔Sirhi = BSirhi锡尔
	f.锡尔Sirhi.SetParent(f)
	
}
