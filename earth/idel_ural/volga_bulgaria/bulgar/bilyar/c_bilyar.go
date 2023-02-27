package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BilyarCounty interface {
    feud.County
    BBilyar比利亚尔() 	feud.Barony
    BBogulma布古利马() 	feud.Barony
    BBryakhimov布里亚希莫夫() 	feud.Barony
    BCukataw茹凯陶() 	feud.Barony
    BDzhalil贾利利() 	feud.Barony
    BNurlat努尔拉特() 	feud.Barony
    BTukhchin图克钦() 	feud.Barony
    BUrussu乌鲁苏() 	feud.Barony
}

type 比利亚尔BilyarCounty struct {
	feud.BaseCounty
	比利亚尔Bilyar 	feud.Barony
	布古利马Bogulma 	feud.Barony
	布里亚希莫夫Bryakhimov 	feud.Barony
	茹凯陶Cukataw 	feud.Barony
	贾利利Dzhalil 	feud.Barony
	努尔拉特Nurlat 	feud.Barony
	图克钦Tukhchin 	feud.Barony
	乌鲁苏Urussu 	feud.Barony
}

func (c *比利亚尔BilyarCounty) BBilyar比利亚尔() feud.Barony {
	return c.比利亚尔Bilyar
}
    
func (c *比利亚尔BilyarCounty) BBogulma布古利马() feud.Barony {
	return c.布古利马Bogulma
}
    
func (c *比利亚尔BilyarCounty) BBryakhimov布里亚希莫夫() feud.Barony {
	return c.布里亚希莫夫Bryakhimov
}
    
func (c *比利亚尔BilyarCounty) BCukataw茹凯陶() feud.Barony {
	return c.茹凯陶Cukataw
}
    
func (c *比利亚尔BilyarCounty) BDzhalil贾利利() feud.Barony {
	return c.贾利利Dzhalil
}
    
func (c *比利亚尔BilyarCounty) BNurlat努尔拉特() feud.Barony {
	return c.努尔拉特Nurlat
}
    
func (c *比利亚尔BilyarCounty) BTukhchin图克钦() feud.Barony {
	return c.图克钦Tukhchin
}
    
func (c *比利亚尔BilyarCounty) BUrussu乌鲁苏() feud.Barony {
	return c.乌鲁苏Urussu
}
    
var CBilyar比利亚尔 BilyarCounty = &比利亚尔BilyarCounty{}

func init() {
	f := CBilyar比利亚尔.(*比利亚尔BilyarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "613",
		Title:     "bilyar",
		TitleName: "比利亚尔",
		TitleCode: "c_bilyar",
		Baronies:  map[string]feud.Barony{},
	}

	f.比利亚尔Bilyar = BBilyar比利亚尔
	f.比利亚尔Bilyar.SetParent(f)
	
	f.布古利马Bogulma = BBogulma布古利马
	f.布古利马Bogulma.SetParent(f)
	
	f.布里亚希莫夫Bryakhimov = BBryakhimov布里亚希莫夫
	f.布里亚希莫夫Bryakhimov.SetParent(f)
	
	f.茹凯陶Cukataw = BCukataw茹凯陶
	f.茹凯陶Cukataw.SetParent(f)
	
	f.贾利利Dzhalil = BDzhalil贾利利
	f.贾利利Dzhalil.SetParent(f)
	
	f.努尔拉特Nurlat = BNurlat努尔拉特
	f.努尔拉特Nurlat.SetParent(f)
	
	f.图克钦Tukhchin = BTukhchin图克钦
	f.图克钦Tukhchin.SetParent(f)
	
	f.乌鲁苏Urussu = BUrussu乌鲁苏
	f.乌鲁苏Urussu.SetParent(f)
	
}
