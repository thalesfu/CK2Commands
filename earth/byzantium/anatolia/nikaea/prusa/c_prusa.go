package prusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PrusaCounty interface {
    feud.County
    BAdrastea阿德拉斯提亚() 	feud.Barony
    BApamea阿帕米亚() 	feud.Barony
    BDarieium达里埃翁() 	feud.Barony
    BDocimium多西米翁() 	feud.Barony
    BMiletopolis米莱托波利斯() 	feud.Barony
    BPelopia佩洛皮亚() 	feud.Barony
    BPrusa普鲁萨() 	feud.Barony
    BThyatira提亚提拉() 	feud.Barony
}

type 普鲁萨PrusaCounty struct {
	feud.BaseCounty
	阿德拉斯提亚Adrastea 	feud.Barony
	阿帕米亚Apamea 	feud.Barony
	达里埃翁Darieium 	feud.Barony
	多西米翁Docimium 	feud.Barony
	米莱托波利斯Miletopolis 	feud.Barony
	佩洛皮亚Pelopia 	feud.Barony
	普鲁萨Prusa 	feud.Barony
	提亚提拉Thyatira 	feud.Barony
}

func (c *普鲁萨PrusaCounty) BAdrastea阿德拉斯提亚() feud.Barony {
	return c.阿德拉斯提亚Adrastea
}
    
func (c *普鲁萨PrusaCounty) BApamea阿帕米亚() feud.Barony {
	return c.阿帕米亚Apamea
}
    
func (c *普鲁萨PrusaCounty) BDarieium达里埃翁() feud.Barony {
	return c.达里埃翁Darieium
}
    
func (c *普鲁萨PrusaCounty) BDocimium多西米翁() feud.Barony {
	return c.多西米翁Docimium
}
    
func (c *普鲁萨PrusaCounty) BMiletopolis米莱托波利斯() feud.Barony {
	return c.米莱托波利斯Miletopolis
}
    
func (c *普鲁萨PrusaCounty) BPelopia佩洛皮亚() feud.Barony {
	return c.佩洛皮亚Pelopia
}
    
func (c *普鲁萨PrusaCounty) BPrusa普鲁萨() feud.Barony {
	return c.普鲁萨Prusa
}
    
func (c *普鲁萨PrusaCounty) BThyatira提亚提拉() feud.Barony {
	return c.提亚提拉Thyatira
}
    
var CPrusa普鲁萨 PrusaCounty = &普鲁萨PrusaCounty{}

func init() {
	f := CPrusa普鲁萨.(*普鲁萨PrusaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "742",
		Title:     "prusa",
		TitleName: "普鲁萨",
		TitleCode: "c_prusa",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德拉斯提亚Adrastea = BAdrastea阿德拉斯提亚
	f.阿德拉斯提亚Adrastea.SetParent(f)
	
	f.阿帕米亚Apamea = BApamea阿帕米亚
	f.阿帕米亚Apamea.SetParent(f)
	
	f.达里埃翁Darieium = BDarieium达里埃翁
	f.达里埃翁Darieium.SetParent(f)
	
	f.多西米翁Docimium = BDocimium多西米翁
	f.多西米翁Docimium.SetParent(f)
	
	f.米莱托波利斯Miletopolis = BMiletopolis米莱托波利斯
	f.米莱托波利斯Miletopolis.SetParent(f)
	
	f.佩洛皮亚Pelopia = BPelopia佩洛皮亚
	f.佩洛皮亚Pelopia.SetParent(f)
	
	f.普鲁萨Prusa = BPrusa普鲁萨
	f.普鲁萨Prusa.SetParent(f)
	
	f.提亚提拉Thyatira = BThyatira提亚提拉
	f.提亚提拉Thyatira.SetParent(f)
	
}
