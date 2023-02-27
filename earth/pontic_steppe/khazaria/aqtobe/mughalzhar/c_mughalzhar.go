package mughalzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MughalzharCounty interface {
    feud.County
    BBolgarka博尔加尔卡() 	feud.Barony
    BKandyagash坎德阿加什() 	feud.Barony
    BKenzhaly肯扎雷() 	feud.Barony
    BKopa科帕() 	feud.Barony
    BMughalzhar穆加尔扎尔() 	feud.Barony
    BTemir铁米尔() 	feud.Barony
    BUil乌伊尔() 	feud.Barony
}

type 穆加尔扎尔MughalzharCounty struct {
	feud.BaseCounty
	博尔加尔卡Bolgarka 	feud.Barony
	坎德阿加什Kandyagash 	feud.Barony
	肯扎雷Kenzhaly 	feud.Barony
	科帕Kopa 	feud.Barony
	穆加尔扎尔Mughalzhar 	feud.Barony
	铁米尔Temir 	feud.Barony
	乌伊尔Uil 	feud.Barony
}

func (c *穆加尔扎尔MughalzharCounty) BBolgarka博尔加尔卡() feud.Barony {
	return c.博尔加尔卡Bolgarka
}
    
func (c *穆加尔扎尔MughalzharCounty) BKandyagash坎德阿加什() feud.Barony {
	return c.坎德阿加什Kandyagash
}
    
func (c *穆加尔扎尔MughalzharCounty) BKenzhaly肯扎雷() feud.Barony {
	return c.肯扎雷Kenzhaly
}
    
func (c *穆加尔扎尔MughalzharCounty) BKopa科帕() feud.Barony {
	return c.科帕Kopa
}
    
func (c *穆加尔扎尔MughalzharCounty) BMughalzhar穆加尔扎尔() feud.Barony {
	return c.穆加尔扎尔Mughalzhar
}
    
func (c *穆加尔扎尔MughalzharCounty) BTemir铁米尔() feud.Barony {
	return c.铁米尔Temir
}
    
func (c *穆加尔扎尔MughalzharCounty) BUil乌伊尔() feud.Barony {
	return c.乌伊尔Uil
}
    
var CMughalzhar穆加尔扎尔 MughalzharCounty = &穆加尔扎尔MughalzharCounty{}

func init() {
	f := CMughalzhar穆加尔扎尔.(*穆加尔扎尔MughalzharCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1851",
		Title:     "mughalzhar",
		TitleName: "穆加尔扎尔",
		TitleCode: "c_mughalzhar",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔加尔卡Bolgarka = BBolgarka博尔加尔卡
	f.博尔加尔卡Bolgarka.SetParent(f)
	
	f.坎德阿加什Kandyagash = BKandyagash坎德阿加什
	f.坎德阿加什Kandyagash.SetParent(f)
	
	f.肯扎雷Kenzhaly = BKenzhaly肯扎雷
	f.肯扎雷Kenzhaly.SetParent(f)
	
	f.科帕Kopa = BKopa科帕
	f.科帕Kopa.SetParent(f)
	
	f.穆加尔扎尔Mughalzhar = BMughalzhar穆加尔扎尔
	f.穆加尔扎尔Mughalzhar.SetParent(f)
	
	f.铁米尔Temir = BTemir铁米尔
	f.铁米尔Temir.SetParent(f)
	
	f.乌伊尔Uil = BUil乌伊尔
	f.乌伊尔Uil.SetParent(f)
	
}
