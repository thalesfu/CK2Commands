package the_isles

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/the_isles/isle_of_man"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type The_islesDuke interface {
    feud.Duke
    CIsle_of_man马恩岛() 	isle_of_man.Isle_of_manCounty
}

type 马恩The_islesDuke struct {
	feud.BaseDuke
	马恩岛Isle_of_man 	isle_of_man.Isle_of_manCounty
}

func (d *马恩The_islesDuke) CIsle_of_man马恩岛() isle_of_man.Isle_of_manCounty {
	return d.马恩岛Isle_of_man
}
    
var DThe_isles马恩 The_islesDuke = &马恩The_islesDuke{}

func init() {
	f := DThe_isles马恩.(*马恩The_islesDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "the_isles",
		TitleName: "马恩",
		TitleCode: "d_the_isles",
		Counties:  map[string]feud.County{},
	}

	f.马恩岛Isle_of_man = isle_of_man.CIsle_of_man马恩岛
	f.马恩岛Isle_of_man.SetParent(f)
	
}
