package trebizond

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/trebizond/chaldea"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/trebizond/theodosiopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/trebizond/trebizond/trapezous"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TrebizondDuke interface {
    feud.Duke
    CChaldea迦勒底() 	chaldea.ChaldeaCounty
    CTheodosiopolis狄奥多西波利斯() 	theodosiopolis.TheodosiopolisCounty
    CTrapezous特拉佩祖斯() 	trapezous.TrapezousCounty
}

type 特拉比松TrebizondDuke struct {
	feud.BaseDuke
	迦勒底Chaldea 	chaldea.ChaldeaCounty
	狄奥多西波利斯Theodosiopolis 	theodosiopolis.TheodosiopolisCounty
	特拉佩祖斯Trapezous 	trapezous.TrapezousCounty
}

func (d *特拉比松TrebizondDuke) CChaldea迦勒底() chaldea.ChaldeaCounty {
	return d.迦勒底Chaldea
}
    
func (d *特拉比松TrebizondDuke) CTheodosiopolis狄奥多西波利斯() theodosiopolis.TheodosiopolisCounty {
	return d.狄奥多西波利斯Theodosiopolis
}
    
func (d *特拉比松TrebizondDuke) CTrapezous特拉佩祖斯() trapezous.TrapezousCounty {
	return d.特拉佩祖斯Trapezous
}
    
var DTrebizond特拉比松 TrebizondDuke = &特拉比松TrebizondDuke{}

func init() {
	f := DTrebizond特拉比松.(*特拉比松TrebizondDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "trebizond",
		TitleName: "特拉比松",
		TitleCode: "d_trebizond",
		Counties:  map[string]feud.County{},
	}

	f.迦勒底Chaldea = chaldea.CChaldea迦勒底
	f.迦勒底Chaldea.SetParent(f)
	
	f.狄奥多西波利斯Theodosiopolis = theodosiopolis.CTheodosiopolis狄奥多西波利斯
	f.狄奥多西波利斯Theodosiopolis.SetParent(f)
	
	f.特拉佩祖斯Trapezous = trapezous.CTrapezous特拉佩祖斯
	f.特拉佩祖斯Trapezous.SetParent(f)
	
}
