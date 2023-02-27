package doti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DotiCounty interface {
    feud.County
    BDadeldhura达代尔图拉() 	feud.Barony
    BDailekh代莱克() 	feud.Barony
    BDhamboji昙菩耆() 	feud.Barony
    BDullu杜鲁() 	feud.Barony
    BGularia古拉里亚() 	feud.Barony
    BKhalanga伽兰迦() 	feud.Barony
    BMahendranagar摩醯因陀罗城() 	feud.Barony
}

type 多蒂DotiCounty struct {
	feud.BaseCounty
	达代尔图拉Dadeldhura 	feud.Barony
	代莱克Dailekh 	feud.Barony
	昙菩耆Dhamboji 	feud.Barony
	杜鲁Dullu 	feud.Barony
	古拉里亚Gularia 	feud.Barony
	伽兰迦Khalanga 	feud.Barony
	摩醯因陀罗城Mahendranagar 	feud.Barony
}

func (c *多蒂DotiCounty) BDadeldhura达代尔图拉() feud.Barony {
	return c.达代尔图拉Dadeldhura
}
    
func (c *多蒂DotiCounty) BDailekh代莱克() feud.Barony {
	return c.代莱克Dailekh
}
    
func (c *多蒂DotiCounty) BDhamboji昙菩耆() feud.Barony {
	return c.昙菩耆Dhamboji
}
    
func (c *多蒂DotiCounty) BDullu杜鲁() feud.Barony {
	return c.杜鲁Dullu
}
    
func (c *多蒂DotiCounty) BGularia古拉里亚() feud.Barony {
	return c.古拉里亚Gularia
}
    
func (c *多蒂DotiCounty) BKhalanga伽兰迦() feud.Barony {
	return c.伽兰迦Khalanga
}
    
func (c *多蒂DotiCounty) BMahendranagar摩醯因陀罗城() feud.Barony {
	return c.摩醯因陀罗城Mahendranagar
}
    
var CDoti多蒂 DotiCounty = &多蒂DotiCounty{}

func init() {
	f := CDoti多蒂.(*多蒂DotiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1473",
		Title:     "doti",
		TitleName: "多蒂",
		TitleCode: "c_doti",
		Baronies:  map[string]feud.Barony{},
	}

	f.达代尔图拉Dadeldhura = BDadeldhura达代尔图拉
	f.达代尔图拉Dadeldhura.SetParent(f)
	
	f.代莱克Dailekh = BDailekh代莱克
	f.代莱克Dailekh.SetParent(f)
	
	f.昙菩耆Dhamboji = BDhamboji昙菩耆
	f.昙菩耆Dhamboji.SetParent(f)
	
	f.杜鲁Dullu = BDullu杜鲁
	f.杜鲁Dullu.SetParent(f)
	
	f.古拉里亚Gularia = BGularia古拉里亚
	f.古拉里亚Gularia.SetParent(f)
	
	f.伽兰迦Khalanga = BKhalanga伽兰迦
	f.伽兰迦Khalanga.SetParent(f)
	
	f.摩醯因陀罗城Mahendranagar = BMahendranagar摩醯因陀罗城
	f.摩醯因陀罗城Mahendranagar.SetParent(f)
	
}
