package zaranj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZaranjCounty interface {
    feud.County
    BBonjar博尼亚尔() 	feud.Barony
    BMilak米拉克() 	feud.Barony
    BNimeh尼梅希() 	feud.Barony
    BRudbar鲁德巴尔() 	feud.Barony
    BZabol扎博勒() 	feud.Barony
    BZahak扎哈克() 	feud.Barony
    BZaranj疾陵() 	feud.Barony
}

type 疾陵ZaranjCounty struct {
	feud.BaseCounty
	博尼亚尔Bonjar 	feud.Barony
	米拉克Milak 	feud.Barony
	尼梅希Nimeh 	feud.Barony
	鲁德巴尔Rudbar 	feud.Barony
	扎博勒Zabol 	feud.Barony
	扎哈克Zahak 	feud.Barony
	疾陵Zaranj 	feud.Barony
}

func (c *疾陵ZaranjCounty) BBonjar博尼亚尔() feud.Barony {
	return c.博尼亚尔Bonjar
}
    
func (c *疾陵ZaranjCounty) BMilak米拉克() feud.Barony {
	return c.米拉克Milak
}
    
func (c *疾陵ZaranjCounty) BNimeh尼梅希() feud.Barony {
	return c.尼梅希Nimeh
}
    
func (c *疾陵ZaranjCounty) BRudbar鲁德巴尔() feud.Barony {
	return c.鲁德巴尔Rudbar
}
    
func (c *疾陵ZaranjCounty) BZabol扎博勒() feud.Barony {
	return c.扎博勒Zabol
}
    
func (c *疾陵ZaranjCounty) BZahak扎哈克() feud.Barony {
	return c.扎哈克Zahak
}
    
func (c *疾陵ZaranjCounty) BZaranj疾陵() feud.Barony {
	return c.疾陵Zaranj
}
    
var CZaranj疾陵 ZaranjCounty = &疾陵ZaranjCounty{}

func init() {
	f := CZaranj疾陵.(*疾陵ZaranjCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1275",
		Title:     "zaranj",
		TitleName: "疾陵",
		TitleCode: "c_zaranj",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尼亚尔Bonjar = BBonjar博尼亚尔
	f.博尼亚尔Bonjar.SetParent(f)
	
	f.米拉克Milak = BMilak米拉克
	f.米拉克Milak.SetParent(f)
	
	f.尼梅希Nimeh = BNimeh尼梅希
	f.尼梅希Nimeh.SetParent(f)
	
	f.鲁德巴尔Rudbar = BRudbar鲁德巴尔
	f.鲁德巴尔Rudbar.SetParent(f)
	
	f.扎博勒Zabol = BZabol扎博勒
	f.扎博勒Zabol.SetParent(f)
	
	f.扎哈克Zahak = BZahak扎哈克
	f.扎哈克Zahak.SetParent(f)
	
	f.疾陵Zaranj = BZaranj疾陵
	f.疾陵Zaranj.SetParent(f)
	
}
