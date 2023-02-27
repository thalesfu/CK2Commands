package betpaqdala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BetpaqdalaCounty interface {
    feud.County
    BBetpaqdala别特帕克达拉() 	feud.Barony
    BBurybaytal布雷尔拜塔尔() 	feud.Barony
    BKuygan奎甘() 	feud.Barony
    BMirnyy米尔内() 	feud.Barony
    BMynaral梅纳拉尔() 	feud.Barony
    BShiganak希加纳克() 	feud.Barony
    BUlken于尔肯() 	feud.Barony
}

type 别特帕克达拉BetpaqdalaCounty struct {
	feud.BaseCounty
	别特帕克达拉Betpaqdala 	feud.Barony
	布雷尔拜塔尔Burybaytal 	feud.Barony
	奎甘Kuygan 	feud.Barony
	米尔内Mirnyy 	feud.Barony
	梅纳拉尔Mynaral 	feud.Barony
	希加纳克Shiganak 	feud.Barony
	于尔肯Ulken 	feud.Barony
}

func (c *别特帕克达拉BetpaqdalaCounty) BBetpaqdala别特帕克达拉() feud.Barony {
	return c.别特帕克达拉Betpaqdala
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BBurybaytal布雷尔拜塔尔() feud.Barony {
	return c.布雷尔拜塔尔Burybaytal
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BKuygan奎甘() feud.Barony {
	return c.奎甘Kuygan
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BMirnyy米尔内() feud.Barony {
	return c.米尔内Mirnyy
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BMynaral梅纳拉尔() feud.Barony {
	return c.梅纳拉尔Mynaral
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BShiganak希加纳克() feud.Barony {
	return c.希加纳克Shiganak
}
    
func (c *别特帕克达拉BetpaqdalaCounty) BUlken于尔肯() feud.Barony {
	return c.于尔肯Ulken
}
    
var CBetpaqdala别特帕克达拉 BetpaqdalaCounty = &别特帕克达拉BetpaqdalaCounty{}

func init() {
	f := CBetpaqdala别特帕克达拉.(*别特帕克达拉BetpaqdalaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1867",
		Title:     "betpaqdala",
		TitleName: "别特帕克达拉",
		TitleCode: "c_betpaqdala",
		Baronies:  map[string]feud.Barony{},
	}

	f.别特帕克达拉Betpaqdala = BBetpaqdala别特帕克达拉
	f.别特帕克达拉Betpaqdala.SetParent(f)
	
	f.布雷尔拜塔尔Burybaytal = BBurybaytal布雷尔拜塔尔
	f.布雷尔拜塔尔Burybaytal.SetParent(f)
	
	f.奎甘Kuygan = BKuygan奎甘
	f.奎甘Kuygan.SetParent(f)
	
	f.米尔内Mirnyy = BMirnyy米尔内
	f.米尔内Mirnyy.SetParent(f)
	
	f.梅纳拉尔Mynaral = BMynaral梅纳拉尔
	f.梅纳拉尔Mynaral.SetParent(f)
	
	f.希加纳克Shiganak = BShiganak希加纳克
	f.希加纳克Shiganak.SetParent(f)
	
	f.于尔肯Ulken = BUlken于尔肯
	f.于尔肯Ulken.SetParent(f)
	
}
