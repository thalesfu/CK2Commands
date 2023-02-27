package buchan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BuchanCounty interface {
    feud.County
    BAberdeen阿伯丁() 	feud.Barony
    BBanff班夫() 	feud.Barony
    BDeer迪尔() 	feud.Barony
    BEllon埃伦() 	feud.Barony
    BFyvie法维() 	feud.Barony
    BInverurie因弗鲁里() 	feud.Barony
    BKintore金托尔() 	feud.Barony
    BSt_machar圣马哈尔() 	feud.Barony
}

type 巴肯BuchanCounty struct {
	feud.BaseCounty
	阿伯丁Aberdeen 	feud.Barony
	班夫Banff 	feud.Barony
	迪尔Deer 	feud.Barony
	埃伦Ellon 	feud.Barony
	法维Fyvie 	feud.Barony
	因弗鲁里Inverurie 	feud.Barony
	金托尔Kintore 	feud.Barony
	圣马哈尔St_machar 	feud.Barony
}

func (c *巴肯BuchanCounty) BAberdeen阿伯丁() feud.Barony {
	return c.阿伯丁Aberdeen
}
    
func (c *巴肯BuchanCounty) BBanff班夫() feud.Barony {
	return c.班夫Banff
}
    
func (c *巴肯BuchanCounty) BDeer迪尔() feud.Barony {
	return c.迪尔Deer
}
    
func (c *巴肯BuchanCounty) BEllon埃伦() feud.Barony {
	return c.埃伦Ellon
}
    
func (c *巴肯BuchanCounty) BFyvie法维() feud.Barony {
	return c.法维Fyvie
}
    
func (c *巴肯BuchanCounty) BInverurie因弗鲁里() feud.Barony {
	return c.因弗鲁里Inverurie
}
    
func (c *巴肯BuchanCounty) BKintore金托尔() feud.Barony {
	return c.金托尔Kintore
}
    
func (c *巴肯BuchanCounty) BSt_machar圣马哈尔() feud.Barony {
	return c.圣马哈尔St_machar
}
    
var CBuchan巴肯 BuchanCounty = &巴肯BuchanCounty{}

func init() {
	f := CBuchan巴肯.(*巴肯BuchanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "41",
		Title:     "buchan",
		TitleName: "巴肯",
		TitleCode: "c_buchan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯丁Aberdeen = BAberdeen阿伯丁
	f.阿伯丁Aberdeen.SetParent(f)
	
	f.班夫Banff = BBanff班夫
	f.班夫Banff.SetParent(f)
	
	f.迪尔Deer = BDeer迪尔
	f.迪尔Deer.SetParent(f)
	
	f.埃伦Ellon = BEllon埃伦
	f.埃伦Ellon.SetParent(f)
	
	f.法维Fyvie = BFyvie法维
	f.法维Fyvie.SetParent(f)
	
	f.因弗鲁里Inverurie = BInverurie因弗鲁里
	f.因弗鲁里Inverurie.SetParent(f)
	
	f.金托尔Kintore = BKintore金托尔
	f.金托尔Kintore.SetParent(f)
	
	f.圣马哈尔St_machar = BSt_machar圣马哈尔
	f.圣马哈尔St_machar.SetParent(f)
	
}
