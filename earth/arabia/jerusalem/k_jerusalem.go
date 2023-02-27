package jerusalem

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/ascalon"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/galilee"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/jerusalem"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/oultrejourdain"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JerusalemKingdom interface {
    feud.Kingdom
    DAscalon亚实基伦() 	ascalon.AscalonDuke
    DGalilee加利利() 	galilee.GalileeDuke
    DJerusalem耶路撒冷() 	jerusalem.JerusalemDuke
    DOultrejourdain外约旦() 	oultrejourdain.OultrejourdainDuke
}

type 耶路撒冷JerusalemKingdom struct {
	feud.BaseKingdom
	亚实基伦Ascalon 	ascalon.AscalonDuke
	加利利Galilee 	galilee.GalileeDuke
	耶路撒冷Jerusalem 	jerusalem.JerusalemDuke
	外约旦Oultrejourdain 	oultrejourdain.OultrejourdainDuke
}

func (k *耶路撒冷JerusalemKingdom) DAscalon亚实基伦() ascalon.AscalonDuke {
	return k.亚实基伦Ascalon
}
    
func (k *耶路撒冷JerusalemKingdom) DGalilee加利利() galilee.GalileeDuke {
	return k.加利利Galilee
}
    
func (k *耶路撒冷JerusalemKingdom) DJerusalem耶路撒冷() jerusalem.JerusalemDuke {
	return k.耶路撒冷Jerusalem
}
    
func (k *耶路撒冷JerusalemKingdom) DOultrejourdain外约旦() oultrejourdain.OultrejourdainDuke {
	return k.外约旦Oultrejourdain
}
    
var KJerusalem耶路撒冷 JerusalemKingdom = &耶路撒冷JerusalemKingdom{}

func init() {
	f := KJerusalem耶路撒冷.(*耶路撒冷JerusalemKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "jerusalem",
		TitleName: "耶路撒冷",
		TitleCode: "k_jerusalem",
		Dukes:  map[string]feud.Duke{},
	}

	f.亚实基伦Ascalon = ascalon.DAscalon亚实基伦
	f.亚实基伦Ascalon.SetParent(f)
	
	f.加利利Galilee = galilee.DGalilee加利利
	f.加利利Galilee.SetParent(f)
	
	f.耶路撒冷Jerusalem = jerusalem.DJerusalem耶路撒冷
	f.耶路撒冷Jerusalem.SetParent(f)
	
	f.外约旦Oultrejourdain = oultrejourdain.DOultrejourdain外约旦
	f.外约旦Oultrejourdain.SetParent(f)
	
}
