package jazira

import (
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/jazira/amida"
	"github.com/thalesfu/CK2Commands/earth/persia/mesopotamia/jazira/bira"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JaziraDuke interface {
    feud.Duke
    CAmida阿米达() 	amida.AmidaCounty
    CBira比雷() 	bira.BiraCounty
}

type 贾兹拉JaziraDuke struct {
	feud.BaseDuke
	阿米达Amida 	amida.AmidaCounty
	比雷Bira 	bira.BiraCounty
}

func (d *贾兹拉JaziraDuke) CAmida阿米达() amida.AmidaCounty {
	return d.阿米达Amida
}
    
func (d *贾兹拉JaziraDuke) CBira比雷() bira.BiraCounty {
	return d.比雷Bira
}
    
var DJazira贾兹拉 JaziraDuke = &贾兹拉JaziraDuke{}

func init() {
	f := DJazira贾兹拉.(*贾兹拉JaziraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jazira",
		TitleName: "贾兹拉",
		TitleCode: "d_jazira",
		Counties:  map[string]feud.County{},
	}

	f.阿米达Amida = amida.CAmida阿米达
	f.阿米达Amida.SetParent(f)
	
	f.比雷Bira = bira.CBira比雷
	f.比雷Bira.SetParent(f)
	
}
