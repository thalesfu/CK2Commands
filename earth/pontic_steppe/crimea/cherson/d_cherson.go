package cherson

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/cherson/cherson"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/cherson/korchev"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/cherson/theodosia"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/cherson/tmutarakan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChersonDuke interface {
    feud.Duke
    CCherson赫尔松() 	cherson.ChersonCounty
    CKorchev刻赤() 	korchev.KorchevCounty
    CTheodosia狄奥多西亚() 	theodosia.TheodosiaCounty
    CTmutarakan特穆塔拉坎() 	tmutarakan.TmutarakanCounty
}

type 赫尔松ChersonDuke struct {
	feud.BaseDuke
	赫尔松Cherson 	cherson.ChersonCounty
	刻赤Korchev 	korchev.KorchevCounty
	狄奥多西亚Theodosia 	theodosia.TheodosiaCounty
	特穆塔拉坎Tmutarakan 	tmutarakan.TmutarakanCounty
}

func (d *赫尔松ChersonDuke) CCherson赫尔松() cherson.ChersonCounty {
	return d.赫尔松Cherson
}
    
func (d *赫尔松ChersonDuke) CKorchev刻赤() korchev.KorchevCounty {
	return d.刻赤Korchev
}
    
func (d *赫尔松ChersonDuke) CTheodosia狄奥多西亚() theodosia.TheodosiaCounty {
	return d.狄奥多西亚Theodosia
}
    
func (d *赫尔松ChersonDuke) CTmutarakan特穆塔拉坎() tmutarakan.TmutarakanCounty {
	return d.特穆塔拉坎Tmutarakan
}
    
var DCherson赫尔松 ChersonDuke = &赫尔松ChersonDuke{}

func init() {
	f := DCherson赫尔松.(*赫尔松ChersonDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cherson",
		TitleName: "赫尔松",
		TitleCode: "d_cherson",
		Counties:  map[string]feud.County{},
	}

	f.赫尔松Cherson = cherson.CCherson赫尔松
	f.赫尔松Cherson.SetParent(f)
	
	f.刻赤Korchev = korchev.CKorchev刻赤
	f.刻赤Korchev.SetParent(f)
	
	f.狄奥多西亚Theodosia = theodosia.CTheodosia狄奥多西亚
	f.狄奥多西亚Theodosia.SetParent(f)
	
	f.特穆塔拉坎Tmutarakan = tmutarakan.CTmutarakan特穆塔拉坎
	f.特穆塔拉坎Tmutarakan.SetParent(f)
	
}
