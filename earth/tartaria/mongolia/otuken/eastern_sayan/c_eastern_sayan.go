package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Eastern_sayanCounty interface {
    feud.County
    BArzhan阿尔然() 	feud.Barony
    BCherbi切尔比() 	feud.Barony
    BEastern_sayan萨彦() 	feud.Barony
    BErzhey叶尔热伊() 	feud.Barony
    BKundustug昆杜斯图格() 	feud.Barony
    BMezhegey梅热格伊() 	feud.Barony
    BSukpak苏克帕克() 	feud.Barony
}

type 东萨彦Eastern_sayanCounty struct {
	feud.BaseCounty
	阿尔然Arzhan 	feud.Barony
	切尔比Cherbi 	feud.Barony
	萨彦Eastern_sayan 	feud.Barony
	叶尔热伊Erzhey 	feud.Barony
	昆杜斯图格Kundustug 	feud.Barony
	梅热格伊Mezhegey 	feud.Barony
	苏克帕克Sukpak 	feud.Barony
}

func (c *东萨彦Eastern_sayanCounty) BArzhan阿尔然() feud.Barony {
	return c.阿尔然Arzhan
}
    
func (c *东萨彦Eastern_sayanCounty) BCherbi切尔比() feud.Barony {
	return c.切尔比Cherbi
}
    
func (c *东萨彦Eastern_sayanCounty) BEastern_sayan萨彦() feud.Barony {
	return c.萨彦Eastern_sayan
}
    
func (c *东萨彦Eastern_sayanCounty) BErzhey叶尔热伊() feud.Barony {
	return c.叶尔热伊Erzhey
}
    
func (c *东萨彦Eastern_sayanCounty) BKundustug昆杜斯图格() feud.Barony {
	return c.昆杜斯图格Kundustug
}
    
func (c *东萨彦Eastern_sayanCounty) BMezhegey梅热格伊() feud.Barony {
	return c.梅热格伊Mezhegey
}
    
func (c *东萨彦Eastern_sayanCounty) BSukpak苏克帕克() feud.Barony {
	return c.苏克帕克Sukpak
}
    
var CEastern_sayan东萨彦 Eastern_sayanCounty = &东萨彦Eastern_sayanCounty{}

func init() {
	f := CEastern_sayan东萨彦.(*东萨彦Eastern_sayanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1902",
		Title:     "eastern_sayan",
		TitleName: "东萨彦",
		TitleCode: "c_eastern_sayan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔然Arzhan = BArzhan阿尔然
	f.阿尔然Arzhan.SetParent(f)
	
	f.切尔比Cherbi = BCherbi切尔比
	f.切尔比Cherbi.SetParent(f)
	
	f.萨彦Eastern_sayan = BEastern_sayan萨彦
	f.萨彦Eastern_sayan.SetParent(f)
	
	f.叶尔热伊Erzhey = BErzhey叶尔热伊
	f.叶尔热伊Erzhey.SetParent(f)
	
	f.昆杜斯图格Kundustug = BKundustug昆杜斯图格
	f.昆杜斯图格Kundustug.SetParent(f)
	
	f.梅热格伊Mezhegey = BMezhegey梅热格伊
	f.梅热格伊Mezhegey.SetParent(f)
	
	f.苏克帕克Sukpak = BSukpak苏克帕克
	f.苏克帕克Sukpak.SetParent(f)
	
}
