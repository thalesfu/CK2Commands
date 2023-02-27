package ross

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RossCounty interface {
    feud.County
    BApplecross阿普尔克罗斯() 	feud.Barony
    BAvoch奥赫() 	feud.Barony
    BCromarty克罗默蒂() 	feud.Barony
    BDingwall丁沃尔() 	feud.Barony
    BFearn弗恩() 	feud.Barony
    BFortrose福特罗斯() 	feud.Barony
    BRosemarkie罗斯马基() 	feud.Barony
    BTain泰恩() 	feud.Barony
}

type 罗斯RossCounty struct {
	feud.BaseCounty
	阿普尔克罗斯Applecross 	feud.Barony
	奥赫Avoch 	feud.Barony
	克罗默蒂Cromarty 	feud.Barony
	丁沃尔Dingwall 	feud.Barony
	弗恩Fearn 	feud.Barony
	福特罗斯Fortrose 	feud.Barony
	罗斯马基Rosemarkie 	feud.Barony
	泰恩Tain 	feud.Barony
}

func (c *罗斯RossCounty) BApplecross阿普尔克罗斯() feud.Barony {
	return c.阿普尔克罗斯Applecross
}
    
func (c *罗斯RossCounty) BAvoch奥赫() feud.Barony {
	return c.奥赫Avoch
}
    
func (c *罗斯RossCounty) BCromarty克罗默蒂() feud.Barony {
	return c.克罗默蒂Cromarty
}
    
func (c *罗斯RossCounty) BDingwall丁沃尔() feud.Barony {
	return c.丁沃尔Dingwall
}
    
func (c *罗斯RossCounty) BFearn弗恩() feud.Barony {
	return c.弗恩Fearn
}
    
func (c *罗斯RossCounty) BFortrose福特罗斯() feud.Barony {
	return c.福特罗斯Fortrose
}
    
func (c *罗斯RossCounty) BRosemarkie罗斯马基() feud.Barony {
	return c.罗斯马基Rosemarkie
}
    
func (c *罗斯RossCounty) BTain泰恩() feud.Barony {
	return c.泰恩Tain
}
    
var CRoss罗斯 RossCounty = &罗斯RossCounty{}

func init() {
	f := CRoss罗斯.(*罗斯RossCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "39",
		Title:     "ross",
		TitleName: "罗斯",
		TitleCode: "c_ross",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿普尔克罗斯Applecross = BApplecross阿普尔克罗斯
	f.阿普尔克罗斯Applecross.SetParent(f)
	
	f.奥赫Avoch = BAvoch奥赫
	f.奥赫Avoch.SetParent(f)
	
	f.克罗默蒂Cromarty = BCromarty克罗默蒂
	f.克罗默蒂Cromarty.SetParent(f)
	
	f.丁沃尔Dingwall = BDingwall丁沃尔
	f.丁沃尔Dingwall.SetParent(f)
	
	f.弗恩Fearn = BFearn弗恩
	f.弗恩Fearn.SetParent(f)
	
	f.福特罗斯Fortrose = BFortrose福特罗斯
	f.福特罗斯Fortrose.SetParent(f)
	
	f.罗斯马基Rosemarkie = BRosemarkie罗斯马基
	f.罗斯马基Rosemarkie.SetParent(f)
	
	f.泰恩Tain = BTain泰恩
	f.泰恩Tain.SetParent(f)
	
}
