package neuchatel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NeuchatelCounty interface {
    feud.County
    BAvenches阿旺什() 	feud.Barony
    BColombier科隆比耶() 	feud.Barony
    BErguel埃尔盖尔() 	feud.Barony
    BEstavayer埃斯塔瓦耶() 	feud.Barony
    BLabonneville拉博讷维尔() 	feud.Barony
    BLalocle勒洛克勒() 	feud.Barony
    BNeuchatel纳沙泰尔() 	feud.Barony
    BNeuveville讷维尔() 	feud.Barony
    BStimier圣伊米耶() 	feud.Barony
    BValangin瓦朗然() 	feud.Barony
}

type 纳沙泰尔NeuchatelCounty struct {
	feud.BaseCounty
	阿旺什Avenches 	feud.Barony
	科隆比耶Colombier 	feud.Barony
	埃尔盖尔Erguel 	feud.Barony
	埃斯塔瓦耶Estavayer 	feud.Barony
	拉博讷维尔Labonneville 	feud.Barony
	勒洛克勒Lalocle 	feud.Barony
	纳沙泰尔Neuchatel 	feud.Barony
	讷维尔Neuveville 	feud.Barony
	圣伊米耶Stimier 	feud.Barony
	瓦朗然Valangin 	feud.Barony
}

func (c *纳沙泰尔NeuchatelCounty) BAvenches阿旺什() feud.Barony {
	return c.阿旺什Avenches
}
    
func (c *纳沙泰尔NeuchatelCounty) BColombier科隆比耶() feud.Barony {
	return c.科隆比耶Colombier
}
    
func (c *纳沙泰尔NeuchatelCounty) BErguel埃尔盖尔() feud.Barony {
	return c.埃尔盖尔Erguel
}
    
func (c *纳沙泰尔NeuchatelCounty) BEstavayer埃斯塔瓦耶() feud.Barony {
	return c.埃斯塔瓦耶Estavayer
}
    
func (c *纳沙泰尔NeuchatelCounty) BLabonneville拉博讷维尔() feud.Barony {
	return c.拉博讷维尔Labonneville
}
    
func (c *纳沙泰尔NeuchatelCounty) BLalocle勒洛克勒() feud.Barony {
	return c.勒洛克勒Lalocle
}
    
func (c *纳沙泰尔NeuchatelCounty) BNeuchatel纳沙泰尔() feud.Barony {
	return c.纳沙泰尔Neuchatel
}
    
func (c *纳沙泰尔NeuchatelCounty) BNeuveville讷维尔() feud.Barony {
	return c.讷维尔Neuveville
}
    
func (c *纳沙泰尔NeuchatelCounty) BStimier圣伊米耶() feud.Barony {
	return c.圣伊米耶Stimier
}
    
func (c *纳沙泰尔NeuchatelCounty) BValangin瓦朗然() feud.Barony {
	return c.瓦朗然Valangin
}
    
var CNeuchatel纳沙泰尔 NeuchatelCounty = &纳沙泰尔NeuchatelCounty{}

func init() {
	f := CNeuchatel纳沙泰尔.(*纳沙泰尔NeuchatelCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "241",
		Title:     "neuchatel",
		TitleName: "纳沙泰尔",
		TitleCode: "c_neuchatel",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿旺什Avenches = BAvenches阿旺什
	f.阿旺什Avenches.SetParent(f)
	
	f.科隆比耶Colombier = BColombier科隆比耶
	f.科隆比耶Colombier.SetParent(f)
	
	f.埃尔盖尔Erguel = BErguel埃尔盖尔
	f.埃尔盖尔Erguel.SetParent(f)
	
	f.埃斯塔瓦耶Estavayer = BEstavayer埃斯塔瓦耶
	f.埃斯塔瓦耶Estavayer.SetParent(f)
	
	f.拉博讷维尔Labonneville = BLabonneville拉博讷维尔
	f.拉博讷维尔Labonneville.SetParent(f)
	
	f.勒洛克勒Lalocle = BLalocle勒洛克勒
	f.勒洛克勒Lalocle.SetParent(f)
	
	f.纳沙泰尔Neuchatel = BNeuchatel纳沙泰尔
	f.纳沙泰尔Neuchatel.SetParent(f)
	
	f.讷维尔Neuveville = BNeuveville讷维尔
	f.讷维尔Neuveville.SetParent(f)
	
	f.圣伊米耶Stimier = BStimier圣伊米耶
	f.圣伊米耶Stimier.SetParent(f)
	
	f.瓦朗然Valangin = BValangin瓦朗然
	f.瓦朗然Valangin.SetParent(f)
	
}
