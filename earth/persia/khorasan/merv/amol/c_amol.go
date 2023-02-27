package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AmolCounty interface {
    feud.County
    BAkhsisak阿赫斯萨克() 	feud.Barony
    BAmol阿莫勒() 	feud.Barony
    BDautly多特利() 	feud.Barony
    BKenekesyr凯内克斯尔() 	feud.Barony
    BMegejik梅盖吉() 	feud.Barony
    BMirzabek米尔扎别克() 	feud.Barony
    BZamm扎木木() 	feud.Barony
}

type 阿莫勒AmolCounty struct {
	feud.BaseCounty
	阿赫斯萨克Akhsisak 	feud.Barony
	阿莫勒Amol 	feud.Barony
	多特利Dautly 	feud.Barony
	凯内克斯尔Kenekesyr 	feud.Barony
	梅盖吉Megejik 	feud.Barony
	米尔扎别克Mirzabek 	feud.Barony
	扎木木Zamm 	feud.Barony
}

func (c *阿莫勒AmolCounty) BAkhsisak阿赫斯萨克() feud.Barony {
	return c.阿赫斯萨克Akhsisak
}
    
func (c *阿莫勒AmolCounty) BAmol阿莫勒() feud.Barony {
	return c.阿莫勒Amol
}
    
func (c *阿莫勒AmolCounty) BDautly多特利() feud.Barony {
	return c.多特利Dautly
}
    
func (c *阿莫勒AmolCounty) BKenekesyr凯内克斯尔() feud.Barony {
	return c.凯内克斯尔Kenekesyr
}
    
func (c *阿莫勒AmolCounty) BMegejik梅盖吉() feud.Barony {
	return c.梅盖吉Megejik
}
    
func (c *阿莫勒AmolCounty) BMirzabek米尔扎别克() feud.Barony {
	return c.米尔扎别克Mirzabek
}
    
func (c *阿莫勒AmolCounty) BZamm扎木木() feud.Barony {
	return c.扎木木Zamm
}
    
var CAmol阿莫勒 AmolCounty = &阿莫勒AmolCounty{}

func init() {
	f := CAmol阿莫勒.(*阿莫勒AmolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1541",
		Title:     "amol",
		TitleName: "阿莫勒",
		TitleCode: "c_amol",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫斯萨克Akhsisak = BAkhsisak阿赫斯萨克
	f.阿赫斯萨克Akhsisak.SetParent(f)
	
	f.阿莫勒Amol = BAmol阿莫勒
	f.阿莫勒Amol.SetParent(f)
	
	f.多特利Dautly = BDautly多特利
	f.多特利Dautly.SetParent(f)
	
	f.凯内克斯尔Kenekesyr = BKenekesyr凯内克斯尔
	f.凯内克斯尔Kenekesyr.SetParent(f)
	
	f.梅盖吉Megejik = BMegejik梅盖吉
	f.梅盖吉Megejik.SetParent(f)
	
	f.米尔扎别克Mirzabek = BMirzabek米尔扎别克
	f.米尔扎别克Mirzabek.SetParent(f)
	
	f.扎木木Zamm = BZamm扎木木
	f.扎木木Zamm.SetParent(f)
	
}
