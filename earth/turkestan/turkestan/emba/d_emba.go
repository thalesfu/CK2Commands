package emba

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/emba/emba"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/emba/kangly"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EmbaDuke interface {
    feud.Duke
    CEmba恩巴() 	emba.EmbaCounty
    CKangly康里() 	kangly.KanglyCounty
}

type 恩巴EmbaDuke struct {
	feud.BaseDuke
	恩巴Emba 	emba.EmbaCounty
	康里Kangly 	kangly.KanglyCounty
}

func (d *恩巴EmbaDuke) CEmba恩巴() emba.EmbaCounty {
	return d.恩巴Emba
}
    
func (d *恩巴EmbaDuke) CKangly康里() kangly.KanglyCounty {
	return d.康里Kangly
}
    
var DEmba恩巴 EmbaDuke = &恩巴EmbaDuke{}

func init() {
	f := DEmba恩巴.(*恩巴EmbaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "emba",
		TitleName: "恩巴",
		TitleCode: "d_emba",
		Counties:  map[string]feud.County{},
	}

	f.恩巴Emba = emba.CEmba恩巴
	f.恩巴Emba.SetParent(f)
	
	f.康里Kangly = kangly.CKangly康里
	f.康里Kangly.SetParent(f)
	
}
