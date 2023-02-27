package seleukeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SeleukeiaCounty interface {
    feud.County
    BAnemurium阿内穆里翁() 	feud.Barony
    BCorycus科律科斯() 	feud.Barony
    BDalisandus达利桑佐斯() 	feud.Barony
    BGermanak日耳曼纳克() 	feud.Barony
    BIrenopolis伊雷诺波利斯() 	feud.Barony
    BNinica尼尼察() 	feud.Barony
    BSeleukeia塞琉西亚() 	feud.Barony
    BSelinus塞利努斯() 	feud.Barony
}

type 塞琉西亚SeleukeiaCounty struct {
	feud.BaseCounty
	阿内穆里翁Anemurium 	feud.Barony
	科律科斯Corycus 	feud.Barony
	达利桑佐斯Dalisandus 	feud.Barony
	日耳曼纳克Germanak 	feud.Barony
	伊雷诺波利斯Irenopolis 	feud.Barony
	尼尼察Ninica 	feud.Barony
	塞琉西亚Seleukeia 	feud.Barony
	塞利努斯Selinus 	feud.Barony
}

func (c *塞琉西亚SeleukeiaCounty) BAnemurium阿内穆里翁() feud.Barony {
	return c.阿内穆里翁Anemurium
}
    
func (c *塞琉西亚SeleukeiaCounty) BCorycus科律科斯() feud.Barony {
	return c.科律科斯Corycus
}
    
func (c *塞琉西亚SeleukeiaCounty) BDalisandus达利桑佐斯() feud.Barony {
	return c.达利桑佐斯Dalisandus
}
    
func (c *塞琉西亚SeleukeiaCounty) BGermanak日耳曼纳克() feud.Barony {
	return c.日耳曼纳克Germanak
}
    
func (c *塞琉西亚SeleukeiaCounty) BIrenopolis伊雷诺波利斯() feud.Barony {
	return c.伊雷诺波利斯Irenopolis
}
    
func (c *塞琉西亚SeleukeiaCounty) BNinica尼尼察() feud.Barony {
	return c.尼尼察Ninica
}
    
func (c *塞琉西亚SeleukeiaCounty) BSeleukeia塞琉西亚() feud.Barony {
	return c.塞琉西亚Seleukeia
}
    
func (c *塞琉西亚SeleukeiaCounty) BSelinus塞利努斯() feud.Barony {
	return c.塞利努斯Selinus
}
    
var CSeleukeia塞琉西亚 SeleukeiaCounty = &塞琉西亚SeleukeiaCounty{}

func init() {
	f := CSeleukeia塞琉西亚.(*塞琉西亚SeleukeiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "758",
		Title:     "seleukeia",
		TitleName: "塞琉西亚",
		TitleCode: "c_seleukeia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿内穆里翁Anemurium = BAnemurium阿内穆里翁
	f.阿内穆里翁Anemurium.SetParent(f)
	
	f.科律科斯Corycus = BCorycus科律科斯
	f.科律科斯Corycus.SetParent(f)
	
	f.达利桑佐斯Dalisandus = BDalisandus达利桑佐斯
	f.达利桑佐斯Dalisandus.SetParent(f)
	
	f.日耳曼纳克Germanak = BGermanak日耳曼纳克
	f.日耳曼纳克Germanak.SetParent(f)
	
	f.伊雷诺波利斯Irenopolis = BIrenopolis伊雷诺波利斯
	f.伊雷诺波利斯Irenopolis.SetParent(f)
	
	f.尼尼察Ninica = BNinica尼尼察
	f.尼尼察Ninica.SetParent(f)
	
	f.塞琉西亚Seleukeia = BSeleukeia塞琉西亚
	f.塞琉西亚Seleukeia.SetParent(f)
	
	f.塞利努斯Selinus = BSelinus塞利努斯
	f.塞利努斯Selinus.SetParent(f)
	
}
