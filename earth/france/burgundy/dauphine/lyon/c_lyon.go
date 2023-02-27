package lyon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LyonCounty interface {
    feud.County
    BAnse昂斯() 	feud.Barony
    BBrindas布兰达() 	feud.Barony
    BChessy谢西() 	feud.Barony
    BIrigny伊里尼() 	feud.Barony
    BLacenas拉塞纳() 	feud.Barony
    BLyon里昂() 	feud.Barony
    BPusignan普西尼昂() 	feud.Barony
    BStjeanbaptiste圣让巴蒂斯特() 	feud.Barony
}

type 里昂LyonCounty struct {
	feud.BaseCounty
	昂斯Anse 	feud.Barony
	布兰达Brindas 	feud.Barony
	谢西Chessy 	feud.Barony
	伊里尼Irigny 	feud.Barony
	拉塞纳Lacenas 	feud.Barony
	里昂Lyon 	feud.Barony
	普西尼昂Pusignan 	feud.Barony
	圣让巴蒂斯特Stjeanbaptiste 	feud.Barony
}

func (c *里昂LyonCounty) BAnse昂斯() feud.Barony {
	return c.昂斯Anse
}
    
func (c *里昂LyonCounty) BBrindas布兰达() feud.Barony {
	return c.布兰达Brindas
}
    
func (c *里昂LyonCounty) BChessy谢西() feud.Barony {
	return c.谢西Chessy
}
    
func (c *里昂LyonCounty) BIrigny伊里尼() feud.Barony {
	return c.伊里尼Irigny
}
    
func (c *里昂LyonCounty) BLacenas拉塞纳() feud.Barony {
	return c.拉塞纳Lacenas
}
    
func (c *里昂LyonCounty) BLyon里昂() feud.Barony {
	return c.里昂Lyon
}
    
func (c *里昂LyonCounty) BPusignan普西尼昂() feud.Barony {
	return c.普西尼昂Pusignan
}
    
func (c *里昂LyonCounty) BStjeanbaptiste圣让巴蒂斯特() feud.Barony {
	return c.圣让巴蒂斯特Stjeanbaptiste
}
    
var CLyon里昂 LyonCounty = &里昂LyonCounty{}

func init() {
	f := CLyon里昂.(*里昂LyonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "227",
		Title:     "lyon",
		TitleName: "里昂",
		TitleCode: "c_lyon",
		Baronies:  map[string]feud.Barony{},
	}

	f.昂斯Anse = BAnse昂斯
	f.昂斯Anse.SetParent(f)
	
	f.布兰达Brindas = BBrindas布兰达
	f.布兰达Brindas.SetParent(f)
	
	f.谢西Chessy = BChessy谢西
	f.谢西Chessy.SetParent(f)
	
	f.伊里尼Irigny = BIrigny伊里尼
	f.伊里尼Irigny.SetParent(f)
	
	f.拉塞纳Lacenas = BLacenas拉塞纳
	f.拉塞纳Lacenas.SetParent(f)
	
	f.里昂Lyon = BLyon里昂
	f.里昂Lyon.SetParent(f)
	
	f.普西尼昂Pusignan = BPusignan普西尼昂
	f.普西尼昂Pusignan.SetParent(f)
	
	f.圣让巴蒂斯特Stjeanbaptiste = BStjeanbaptiste圣让巴蒂斯特
	f.圣让巴蒂斯特Stjeanbaptiste.SetParent(f)
	
}
