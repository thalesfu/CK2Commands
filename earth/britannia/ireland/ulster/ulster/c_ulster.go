package ulster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UlsterCounty interface {
    feud.County
    BBangor班戈() 	feud.Barony
    BCarrickfergus卡里克弗格斯() 	feud.Barony
    BConnor康诺尔() 	feud.Barony
    BDownpatrick唐帕特里克() 	feud.Barony
    BDromore德罗莫尔() 	feud.Barony
    BDunluce邓卢斯() 	feud.Barony
    BDunseverick顿塞沃里克() 	feud.Barony
    BLarne拉恩() 	feud.Barony
}

type 阿尔斯特UlsterCounty struct {
	feud.BaseCounty
	班戈Bangor 	feud.Barony
	卡里克弗格斯Carrickfergus 	feud.Barony
	康诺尔Connor 	feud.Barony
	唐帕特里克Downpatrick 	feud.Barony
	德罗莫尔Dromore 	feud.Barony
	邓卢斯Dunluce 	feud.Barony
	顿塞沃里克Dunseverick 	feud.Barony
	拉恩Larne 	feud.Barony
}

func (c *阿尔斯特UlsterCounty) BBangor班戈() feud.Barony {
	return c.班戈Bangor
}
    
func (c *阿尔斯特UlsterCounty) BCarrickfergus卡里克弗格斯() feud.Barony {
	return c.卡里克弗格斯Carrickfergus
}
    
func (c *阿尔斯特UlsterCounty) BConnor康诺尔() feud.Barony {
	return c.康诺尔Connor
}
    
func (c *阿尔斯特UlsterCounty) BDownpatrick唐帕特里克() feud.Barony {
	return c.唐帕特里克Downpatrick
}
    
func (c *阿尔斯特UlsterCounty) BDromore德罗莫尔() feud.Barony {
	return c.德罗莫尔Dromore
}
    
func (c *阿尔斯特UlsterCounty) BDunluce邓卢斯() feud.Barony {
	return c.邓卢斯Dunluce
}
    
func (c *阿尔斯特UlsterCounty) BDunseverick顿塞沃里克() feud.Barony {
	return c.顿塞沃里克Dunseverick
}
    
func (c *阿尔斯特UlsterCounty) BLarne拉恩() feud.Barony {
	return c.拉恩Larne
}
    
var CUlster阿尔斯特 UlsterCounty = &阿尔斯特UlsterCounty{}

func init() {
	f := CUlster阿尔斯特.(*阿尔斯特UlsterCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "5",
		Title:     "ulster",
		TitleName: "阿尔斯特",
		TitleCode: "c_ulster",
		Baronies:  map[string]feud.Barony{},
	}

	f.班戈Bangor = BBangor班戈
	f.班戈Bangor.SetParent(f)
	
	f.卡里克弗格斯Carrickfergus = BCarrickfergus卡里克弗格斯
	f.卡里克弗格斯Carrickfergus.SetParent(f)
	
	f.康诺尔Connor = BConnor康诺尔
	f.康诺尔Connor.SetParent(f)
	
	f.唐帕特里克Downpatrick = BDownpatrick唐帕特里克
	f.唐帕特里克Downpatrick.SetParent(f)
	
	f.德罗莫尔Dromore = BDromore德罗莫尔
	f.德罗莫尔Dromore.SetParent(f)
	
	f.邓卢斯Dunluce = BDunluce邓卢斯
	f.邓卢斯Dunluce.SetParent(f)
	
	f.顿塞沃里克Dunseverick = BDunseverick顿塞沃里克
	f.顿塞沃里克Dunseverick.SetParent(f)
	
	f.拉恩Larne = BLarne拉恩
	f.拉恩Larne.SetParent(f)
	
}
