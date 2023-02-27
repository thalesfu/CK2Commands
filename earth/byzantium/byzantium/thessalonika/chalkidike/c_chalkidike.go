package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChalkidikeCounty interface {
    feud.County
    BChrysiopolis赫律索波利斯() 	feud.Barony
    BCrusis克鲁西斯() 	feud.Barony
    BMntathos阿索斯山() 	feud.Barony
    BPallene帕勒涅() 	feud.Barony
    BPolygyros波利伊罗斯() 	feud.Barony
    BSithonia锡索尼亚() 	feud.Barony
    BZicna基克纳() 	feud.Barony
}

type 哈尔基季基ChalkidikeCounty struct {
	feud.BaseCounty
	赫律索波利斯Chrysiopolis 	feud.Barony
	克鲁西斯Crusis 	feud.Barony
	阿索斯山Mntathos 	feud.Barony
	帕勒涅Pallene 	feud.Barony
	波利伊罗斯Polygyros 	feud.Barony
	锡索尼亚Sithonia 	feud.Barony
	基克纳Zicna 	feud.Barony
}

func (c *哈尔基季基ChalkidikeCounty) BChrysiopolis赫律索波利斯() feud.Barony {
	return c.赫律索波利斯Chrysiopolis
}
    
func (c *哈尔基季基ChalkidikeCounty) BCrusis克鲁西斯() feud.Barony {
	return c.克鲁西斯Crusis
}
    
func (c *哈尔基季基ChalkidikeCounty) BMntathos阿索斯山() feud.Barony {
	return c.阿索斯山Mntathos
}
    
func (c *哈尔基季基ChalkidikeCounty) BPallene帕勒涅() feud.Barony {
	return c.帕勒涅Pallene
}
    
func (c *哈尔基季基ChalkidikeCounty) BPolygyros波利伊罗斯() feud.Barony {
	return c.波利伊罗斯Polygyros
}
    
func (c *哈尔基季基ChalkidikeCounty) BSithonia锡索尼亚() feud.Barony {
	return c.锡索尼亚Sithonia
}
    
func (c *哈尔基季基ChalkidikeCounty) BZicna基克纳() feud.Barony {
	return c.基克纳Zicna
}
    
var CChalkidike哈尔基季基 ChalkidikeCounty = &哈尔基季基ChalkidikeCounty{}

func init() {
	f := CChalkidike哈尔基季基.(*哈尔基季基ChalkidikeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "491",
		Title:     "chalkidike",
		TitleName: "哈尔基季基",
		TitleCode: "c_chalkidike",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫律索波利斯Chrysiopolis = BChrysiopolis赫律索波利斯
	f.赫律索波利斯Chrysiopolis.SetParent(f)
	
	f.克鲁西斯Crusis = BCrusis克鲁西斯
	f.克鲁西斯Crusis.SetParent(f)
	
	f.阿索斯山Mntathos = BMntathos阿索斯山
	f.阿索斯山Mntathos.SetParent(f)
	
	f.帕勒涅Pallene = BPallene帕勒涅
	f.帕勒涅Pallene.SetParent(f)
	
	f.波利伊罗斯Polygyros = BPolygyros波利伊罗斯
	f.波利伊罗斯Polygyros.SetParent(f)
	
	f.锡索尼亚Sithonia = BSithonia锡索尼亚
	f.锡索尼亚Sithonia.SetParent(f)
	
	f.基克纳Zicna = BZicna基克纳
	f.基克纳Zicna.SetParent(f)
	
}
