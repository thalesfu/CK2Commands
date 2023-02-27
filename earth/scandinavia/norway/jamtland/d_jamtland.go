package jamtland

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/jamtland/herjedalen"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/norway/jamtland/jamtland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JamtlandDuke interface {
    feud.Duke
    CHerjedalen海里耶达伦() 	herjedalen.HerjedalenCounty
    CJamtland耶姆特兰() 	jamtland.JamtlandCounty
}

type 耶姆特兰JamtlandDuke struct {
	feud.BaseDuke
	海里耶达伦Herjedalen 	herjedalen.HerjedalenCounty
	耶姆特兰Jamtland 	jamtland.JamtlandCounty
}

func (d *耶姆特兰JamtlandDuke) CHerjedalen海里耶达伦() herjedalen.HerjedalenCounty {
	return d.海里耶达伦Herjedalen
}
    
func (d *耶姆特兰JamtlandDuke) CJamtland耶姆特兰() jamtland.JamtlandCounty {
	return d.耶姆特兰Jamtland
}
    
var DJamtland耶姆特兰 JamtlandDuke = &耶姆特兰JamtlandDuke{}

func init() {
	f := DJamtland耶姆特兰.(*耶姆特兰JamtlandDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jamtland",
		TitleName: "耶姆特兰",
		TitleCode: "d_jamtland",
		Counties:  map[string]feud.County{},
	}

	f.海里耶达伦Herjedalen = herjedalen.CHerjedalen海里耶达伦
	f.海里耶达伦Herjedalen.SetParent(f)
	
	f.耶姆特兰Jamtland = jamtland.CJamtland耶姆特兰
	f.耶姆特兰Jamtland.SetParent(f)
	
}
