package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZawilaCounty interface {
    feud.County
    BEt_tuebia图韦比亚() 	feud.Barony
    BJikheira贾哈拉() 	feud.Barony
    BPetras_maior大佩特拉斯() 	feud.Barony
    BSciaua沙瓦() 	feud.Barony
    BZalha泽尔哈() 	feud.Barony
    BZawila宰维莱() 	feud.Barony
    BZizau杰扎乌() 	feud.Barony
}

type 宰维莱ZawilaCounty struct {
	feud.BaseCounty
	图韦比亚Et_tuebia 	feud.Barony
	贾哈拉Jikheira 	feud.Barony
	大佩特拉斯Petras_maior 	feud.Barony
	沙瓦Sciaua 	feud.Barony
	泽尔哈Zalha 	feud.Barony
	宰维莱Zawila 	feud.Barony
	杰扎乌Zizau 	feud.Barony
}

func (c *宰维莱ZawilaCounty) BEt_tuebia图韦比亚() feud.Barony {
	return c.图韦比亚Et_tuebia
}
    
func (c *宰维莱ZawilaCounty) BJikheira贾哈拉() feud.Barony {
	return c.贾哈拉Jikheira
}
    
func (c *宰维莱ZawilaCounty) BPetras_maior大佩特拉斯() feud.Barony {
	return c.大佩特拉斯Petras_maior
}
    
func (c *宰维莱ZawilaCounty) BSciaua沙瓦() feud.Barony {
	return c.沙瓦Sciaua
}
    
func (c *宰维莱ZawilaCounty) BZalha泽尔哈() feud.Barony {
	return c.泽尔哈Zalha
}
    
func (c *宰维莱ZawilaCounty) BZawila宰维莱() feud.Barony {
	return c.宰维莱Zawila
}
    
func (c *宰维莱ZawilaCounty) BZizau杰扎乌() feud.Barony {
	return c.杰扎乌Zizau
}
    
var CZawila宰维莱 ZawilaCounty = &宰维莱ZawilaCounty{}

func init() {
	f := CZawila宰维莱.(*宰维莱ZawilaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1731",
		Title:     "zawila",
		TitleName: "宰维莱",
		TitleCode: "c_zawila",
		Baronies:  map[string]feud.Barony{},
	}

	f.图韦比亚Et_tuebia = BEt_tuebia图韦比亚
	f.图韦比亚Et_tuebia.SetParent(f)
	
	f.贾哈拉Jikheira = BJikheira贾哈拉
	f.贾哈拉Jikheira.SetParent(f)
	
	f.大佩特拉斯Petras_maior = BPetras_maior大佩特拉斯
	f.大佩特拉斯Petras_maior.SetParent(f)
	
	f.沙瓦Sciaua = BSciaua沙瓦
	f.沙瓦Sciaua.SetParent(f)
	
	f.泽尔哈Zalha = BZalha泽尔哈
	f.泽尔哈Zalha.SetParent(f)
	
	f.宰维莱Zawila = BZawila宰维莱
	f.宰维莱Zawila.SetParent(f)
	
	f.杰扎乌Zizau = BZizau杰扎乌
	f.杰扎乌Zizau.SetParent(f)
	
}
