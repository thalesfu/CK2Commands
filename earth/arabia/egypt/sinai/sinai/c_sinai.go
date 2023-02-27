package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SinaiCounty interface {
    feud.County
    BAttur图尔() 	feud.Barony
    BDahab宰海卜() 	feud.Barony
    BJabalsamra扎拜萨马拉() 	feud.Barony
    BMamlah马穆拉() 	feud.Barony
    BNabq奈卜格() 	feud.Barony
    BNuweiba努韦巴() 	feud.Barony
    BSharmelsheikh沙姆沙伊赫() 	feud.Barony
    BSinbarqa拜尔盖() 	feud.Barony
}

type 西奈SinaiCounty struct {
	feud.BaseCounty
	图尔Attur 	feud.Barony
	宰海卜Dahab 	feud.Barony
	扎拜萨马拉Jabalsamra 	feud.Barony
	马穆拉Mamlah 	feud.Barony
	奈卜格Nabq 	feud.Barony
	努韦巴Nuweiba 	feud.Barony
	沙姆沙伊赫Sharmelsheikh 	feud.Barony
	拜尔盖Sinbarqa 	feud.Barony
}

func (c *西奈SinaiCounty) BAttur图尔() feud.Barony {
	return c.图尔Attur
}
    
func (c *西奈SinaiCounty) BDahab宰海卜() feud.Barony {
	return c.宰海卜Dahab
}
    
func (c *西奈SinaiCounty) BJabalsamra扎拜萨马拉() feud.Barony {
	return c.扎拜萨马拉Jabalsamra
}
    
func (c *西奈SinaiCounty) BMamlah马穆拉() feud.Barony {
	return c.马穆拉Mamlah
}
    
func (c *西奈SinaiCounty) BNabq奈卜格() feud.Barony {
	return c.奈卜格Nabq
}
    
func (c *西奈SinaiCounty) BNuweiba努韦巴() feud.Barony {
	return c.努韦巴Nuweiba
}
    
func (c *西奈SinaiCounty) BSharmelsheikh沙姆沙伊赫() feud.Barony {
	return c.沙姆沙伊赫Sharmelsheikh
}
    
func (c *西奈SinaiCounty) BSinbarqa拜尔盖() feud.Barony {
	return c.拜尔盖Sinbarqa
}
    
var CSinai西奈 SinaiCounty = &西奈SinaiCounty{}

func init() {
	f := CSinai西奈.(*西奈SinaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "785",
		Title:     "sinai",
		TitleName: "西奈",
		TitleCode: "c_sinai",
		Baronies:  map[string]feud.Barony{},
	}

	f.图尔Attur = BAttur图尔
	f.图尔Attur.SetParent(f)
	
	f.宰海卜Dahab = BDahab宰海卜
	f.宰海卜Dahab.SetParent(f)
	
	f.扎拜萨马拉Jabalsamra = BJabalsamra扎拜萨马拉
	f.扎拜萨马拉Jabalsamra.SetParent(f)
	
	f.马穆拉Mamlah = BMamlah马穆拉
	f.马穆拉Mamlah.SetParent(f)
	
	f.奈卜格Nabq = BNabq奈卜格
	f.奈卜格Nabq.SetParent(f)
	
	f.努韦巴Nuweiba = BNuweiba努韦巴
	f.努韦巴Nuweiba.SetParent(f)
	
	f.沙姆沙伊赫Sharmelsheikh = BSharmelsheikh沙姆沙伊赫
	f.沙姆沙伊赫Sharmelsheikh.SetParent(f)
	
	f.拜尔盖Sinbarqa = BSinbarqa拜尔盖
	f.拜尔盖Sinbarqa.SetParent(f)
	
}
