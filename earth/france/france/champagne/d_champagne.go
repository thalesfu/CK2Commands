package champagne

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/champagne/brie"
	"github.com/thalesfu/CK2Commands/earth/france/france/champagne/reims"
	"github.com/thalesfu/CK2Commands/earth/france/france/champagne/troyes"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChampagneDuke interface {
    feud.Duke
    CBrie布里() 	brie.BrieCounty
    CReims兰斯() 	reims.ReimsCounty
    CTroyes特鲁瓦() 	troyes.TroyesCounty
}

type 香槟ChampagneDuke struct {
	feud.BaseDuke
	布里Brie 	brie.BrieCounty
	兰斯Reims 	reims.ReimsCounty
	特鲁瓦Troyes 	troyes.TroyesCounty
}

func (d *香槟ChampagneDuke) CBrie布里() brie.BrieCounty {
	return d.布里Brie
}
    
func (d *香槟ChampagneDuke) CReims兰斯() reims.ReimsCounty {
	return d.兰斯Reims
}
    
func (d *香槟ChampagneDuke) CTroyes特鲁瓦() troyes.TroyesCounty {
	return d.特鲁瓦Troyes
}
    
var DChampagne香槟 ChampagneDuke = &香槟ChampagneDuke{}

func init() {
	f := DChampagne香槟.(*香槟ChampagneDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "champagne",
		TitleName: "香槟",
		TitleCode: "d_champagne",
		Counties:  map[string]feud.County{},
	}

	f.布里Brie = brie.CBrie布里
	f.布里Brie.SetParent(f)
	
	f.兰斯Reims = reims.CReims兰斯
	f.兰斯Reims.SetParent(f)
	
	f.特鲁瓦Troyes = troyes.CTroyes特鲁瓦
	f.特鲁瓦Troyes.SetParent(f)
	
}
