package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MainzCounty interface {
    feud.County
    BDornburg多恩堡() 	feud.Barony
    BEppstein埃普施泰因() 	feud.Barony
    BHanau哈瑙() 	feud.Barony
    BIngelheim英格尔海姆() 	feud.Barony
    BMainz美因茨() 	feud.Barony
    BMannheim曼海姆() 	feud.Barony
    BWeilburg魏尔堡() 	feud.Barony
}

type 美因茨MainzCounty struct {
	feud.BaseCounty
	多恩堡Dornburg 	feud.Barony
	埃普施泰因Eppstein 	feud.Barony
	哈瑙Hanau 	feud.Barony
	英格尔海姆Ingelheim 	feud.Barony
	美因茨Mainz 	feud.Barony
	曼海姆Mannheim 	feud.Barony
	魏尔堡Weilburg 	feud.Barony
}

func (c *美因茨MainzCounty) BDornburg多恩堡() feud.Barony {
	return c.多恩堡Dornburg
}
    
func (c *美因茨MainzCounty) BEppstein埃普施泰因() feud.Barony {
	return c.埃普施泰因Eppstein
}
    
func (c *美因茨MainzCounty) BHanau哈瑙() feud.Barony {
	return c.哈瑙Hanau
}
    
func (c *美因茨MainzCounty) BIngelheim英格尔海姆() feud.Barony {
	return c.英格尔海姆Ingelheim
}
    
func (c *美因茨MainzCounty) BMainz美因茨() feud.Barony {
	return c.美因茨Mainz
}
    
func (c *美因茨MainzCounty) BMannheim曼海姆() feud.Barony {
	return c.曼海姆Mannheim
}
    
func (c *美因茨MainzCounty) BWeilburg魏尔堡() feud.Barony {
	return c.魏尔堡Weilburg
}
    
var CMainz美因茨 MainzCounty = &美因茨MainzCounty{}

func init() {
	f := CMainz美因茨.(*美因茨MainzCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "123",
		Title:     "mainz",
		TitleName: "美因茨",
		TitleCode: "c_mainz",
		Baronies:  map[string]feud.Barony{},
	}

	f.多恩堡Dornburg = BDornburg多恩堡
	f.多恩堡Dornburg.SetParent(f)
	
	f.埃普施泰因Eppstein = BEppstein埃普施泰因
	f.埃普施泰因Eppstein.SetParent(f)
	
	f.哈瑙Hanau = BHanau哈瑙
	f.哈瑙Hanau.SetParent(f)
	
	f.英格尔海姆Ingelheim = BIngelheim英格尔海姆
	f.英格尔海姆Ingelheim.SetParent(f)
	
	f.美因茨Mainz = BMainz美因茨
	f.美因茨Mainz.SetParent(f)
	
	f.曼海姆Mannheim = BMannheim曼海姆
	f.曼海姆Mannheim.SetParent(f)
	
	f.魏尔堡Weilburg = BWeilburg魏尔堡
	f.魏尔堡Weilburg.SetParent(f)
	
}
