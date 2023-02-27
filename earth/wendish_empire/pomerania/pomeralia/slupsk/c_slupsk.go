package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SlupskCounty interface {
    feud.County
    BBelgard贝尔加德() 	feud.Barony
    BColberg科尔贝格() 	feud.Barony
    BCorlin格林() 	feud.Barony
    BGerwin格温() 	feud.Barony
    BRugenwalde吕根瓦尔德() 	feud.Barony
    BSchlawe施拉沃() 	feud.Barony
    BSlupsk斯武普斯克() 	feud.Barony
    BUstka乌斯特卡() 	feud.Barony
}

type 斯武普斯克SlupskCounty struct {
	feud.BaseCounty
	贝尔加德Belgard 	feud.Barony
	科尔贝格Colberg 	feud.Barony
	格林Corlin 	feud.Barony
	格温Gerwin 	feud.Barony
	吕根瓦尔德Rugenwalde 	feud.Barony
	施拉沃Schlawe 	feud.Barony
	斯武普斯克Slupsk 	feud.Barony
	乌斯特卡Ustka 	feud.Barony
}

func (c *斯武普斯克SlupskCounty) BBelgard贝尔加德() feud.Barony {
	return c.贝尔加德Belgard
}
    
func (c *斯武普斯克SlupskCounty) BColberg科尔贝格() feud.Barony {
	return c.科尔贝格Colberg
}
    
func (c *斯武普斯克SlupskCounty) BCorlin格林() feud.Barony {
	return c.格林Corlin
}
    
func (c *斯武普斯克SlupskCounty) BGerwin格温() feud.Barony {
	return c.格温Gerwin
}
    
func (c *斯武普斯克SlupskCounty) BRugenwalde吕根瓦尔德() feud.Barony {
	return c.吕根瓦尔德Rugenwalde
}
    
func (c *斯武普斯克SlupskCounty) BSchlawe施拉沃() feud.Barony {
	return c.施拉沃Schlawe
}
    
func (c *斯武普斯克SlupskCounty) BSlupsk斯武普斯克() feud.Barony {
	return c.斯武普斯克Slupsk
}
    
func (c *斯武普斯克SlupskCounty) BUstka乌斯特卡() feud.Barony {
	return c.乌斯特卡Ustka
}
    
var CSlupsk斯武普斯克 SlupskCounty = &斯武普斯克SlupskCounty{}

func init() {
	f := CSlupsk斯武普斯克.(*斯武普斯克SlupskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "367",
		Title:     "slupsk",
		TitleName: "斯武普斯克",
		TitleCode: "c_slupsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔加德Belgard = BBelgard贝尔加德
	f.贝尔加德Belgard.SetParent(f)
	
	f.科尔贝格Colberg = BColberg科尔贝格
	f.科尔贝格Colberg.SetParent(f)
	
	f.格林Corlin = BCorlin格林
	f.格林Corlin.SetParent(f)
	
	f.格温Gerwin = BGerwin格温
	f.格温Gerwin.SetParent(f)
	
	f.吕根瓦尔德Rugenwalde = BRugenwalde吕根瓦尔德
	f.吕根瓦尔德Rugenwalde.SetParent(f)
	
	f.施拉沃Schlawe = BSchlawe施拉沃
	f.施拉沃Schlawe.SetParent(f)
	
	f.斯武普斯克Slupsk = BSlupsk斯武普斯克
	f.斯武普斯克Slupsk.SetParent(f)
	
	f.乌斯特卡Ustka = BUstka乌斯特卡
	f.乌斯特卡Ustka.SetParent(f)
	
}
