package north_dvina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type North_dvinaCounty interface {
    feud.County
    BAntonievosiysky安托尼耶沃_西斯基() 	feud.Barony
    BArchangelsk阿尔汉格尔斯克() 	feud.Barony
    BKholmogory霍尔莫戈雷() 	feud.Barony
    BKoryazhma科里亚日马() 	feud.Barony
    BNikolokorelski尼科洛_科列尔斯基() 	feud.Barony
    BNovokholmogory新霍尔莫戈雷() 	feud.Barony
    BSolvychegodsk索里维切戈茨克() 	feud.Barony
    BUsolsk乌索利斯克() 	feud.Barony
}

type 德维纳North_dvinaCounty struct {
	feud.BaseCounty
	安托尼耶沃_西斯基Antonievosiysky 	feud.Barony
	阿尔汉格尔斯克Archangelsk 	feud.Barony
	霍尔莫戈雷Kholmogory 	feud.Barony
	科里亚日马Koryazhma 	feud.Barony
	尼科洛_科列尔斯基Nikolokorelski 	feud.Barony
	新霍尔莫戈雷Novokholmogory 	feud.Barony
	索里维切戈茨克Solvychegodsk 	feud.Barony
	乌索利斯克Usolsk 	feud.Barony
}

func (c *德维纳North_dvinaCounty) BAntonievosiysky安托尼耶沃_西斯基() feud.Barony {
	return c.安托尼耶沃_西斯基Antonievosiysky
}
    
func (c *德维纳North_dvinaCounty) BArchangelsk阿尔汉格尔斯克() feud.Barony {
	return c.阿尔汉格尔斯克Archangelsk
}
    
func (c *德维纳North_dvinaCounty) BKholmogory霍尔莫戈雷() feud.Barony {
	return c.霍尔莫戈雷Kholmogory
}
    
func (c *德维纳North_dvinaCounty) BKoryazhma科里亚日马() feud.Barony {
	return c.科里亚日马Koryazhma
}
    
func (c *德维纳North_dvinaCounty) BNikolokorelski尼科洛_科列尔斯基() feud.Barony {
	return c.尼科洛_科列尔斯基Nikolokorelski
}
    
func (c *德维纳North_dvinaCounty) BNovokholmogory新霍尔莫戈雷() feud.Barony {
	return c.新霍尔莫戈雷Novokholmogory
}
    
func (c *德维纳North_dvinaCounty) BSolvychegodsk索里维切戈茨克() feud.Barony {
	return c.索里维切戈茨克Solvychegodsk
}
    
func (c *德维纳North_dvinaCounty) BUsolsk乌索利斯克() feud.Barony {
	return c.乌索利斯克Usolsk
}
    
var CNorth_dvina德维纳 North_dvinaCounty = &德维纳North_dvinaCounty{}

func init() {
	f := CNorth_dvina德维纳.(*德维纳North_dvinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "395",
		Title:     "north_dvina",
		TitleName: "德维纳",
		TitleCode: "c_north_dvina",
		Baronies:  map[string]feud.Barony{},
	}

	f.安托尼耶沃_西斯基Antonievosiysky = BAntonievosiysky安托尼耶沃_西斯基
	f.安托尼耶沃_西斯基Antonievosiysky.SetParent(f)
	
	f.阿尔汉格尔斯克Archangelsk = BArchangelsk阿尔汉格尔斯克
	f.阿尔汉格尔斯克Archangelsk.SetParent(f)
	
	f.霍尔莫戈雷Kholmogory = BKholmogory霍尔莫戈雷
	f.霍尔莫戈雷Kholmogory.SetParent(f)
	
	f.科里亚日马Koryazhma = BKoryazhma科里亚日马
	f.科里亚日马Koryazhma.SetParent(f)
	
	f.尼科洛_科列尔斯基Nikolokorelski = BNikolokorelski尼科洛_科列尔斯基
	f.尼科洛_科列尔斯基Nikolokorelski.SetParent(f)
	
	f.新霍尔莫戈雷Novokholmogory = BNovokholmogory新霍尔莫戈雷
	f.新霍尔莫戈雷Novokholmogory.SetParent(f)
	
	f.索里维切戈茨克Solvychegodsk = BSolvychegodsk索里维切戈茨克
	f.索里维切戈茨克Solvychegodsk.SetParent(f)
	
	f.乌索利斯克Usolsk = BUsolsk乌索利斯克
	f.乌索利斯克Usolsk.SetParent(f)
	
}
