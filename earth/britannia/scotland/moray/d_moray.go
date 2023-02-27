package moray

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/moray/buchan"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/moray/caithness"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/moray/moray"
	"github.com/thalesfu/CK2Commands/earth/britannia/scotland/moray/ross"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MorayDuke interface {
    feud.Duke
    CBuchan巴肯() 	buchan.BuchanCounty
    CCaithness凯斯内斯() 	caithness.CaithnessCounty
    CMoray马里() 	moray.MorayCounty
    CRoss罗斯() 	ross.RossCounty
}

type 马里MorayDuke struct {
	feud.BaseDuke
	巴肯Buchan 	buchan.BuchanCounty
	凯斯内斯Caithness 	caithness.CaithnessCounty
	马里Moray 	moray.MorayCounty
	罗斯Ross 	ross.RossCounty
}

func (d *马里MorayDuke) CBuchan巴肯() buchan.BuchanCounty {
	return d.巴肯Buchan
}
    
func (d *马里MorayDuke) CCaithness凯斯内斯() caithness.CaithnessCounty {
	return d.凯斯内斯Caithness
}
    
func (d *马里MorayDuke) CMoray马里() moray.MorayCounty {
	return d.马里Moray
}
    
func (d *马里MorayDuke) CRoss罗斯() ross.RossCounty {
	return d.罗斯Ross
}
    
var DMoray马里 MorayDuke = &马里MorayDuke{}

func init() {
	f := DMoray马里.(*马里MorayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "moray",
		TitleName: "马里",
		TitleCode: "d_moray",
		Counties:  map[string]feud.County{},
	}

	f.巴肯Buchan = buchan.CBuchan巴肯
	f.巴肯Buchan.SetParent(f)
	
	f.凯斯内斯Caithness = caithness.CCaithness凯斯内斯
	f.凯斯内斯Caithness.SetParent(f)
	
	f.马里Moray = moray.CMoray马里
	f.马里Moray.SetParent(f)
	
	f.罗斯Ross = ross.CRoss罗斯
	f.罗斯Ross.SetParent(f)
	
}
