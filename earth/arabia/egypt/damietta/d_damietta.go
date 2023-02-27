package damietta

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/damietta/delta"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/damietta/manupura"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/damietta/pelusia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamiettaDuke interface {
    feud.Duke
    CDelta三角洲() 	delta.DeltaCounty
    CManupura迈哈莱() 	manupura.ManupuraCounty
    CPelusia廷尼斯() 	pelusia.PelusiaCounty
}

type 达米埃塔DamiettaDuke struct {
	feud.BaseDuke
	三角洲Delta 	delta.DeltaCounty
	迈哈莱Manupura 	manupura.ManupuraCounty
	廷尼斯Pelusia 	pelusia.PelusiaCounty
}

func (d *达米埃塔DamiettaDuke) CDelta三角洲() delta.DeltaCounty {
	return d.三角洲Delta
}
    
func (d *达米埃塔DamiettaDuke) CManupura迈哈莱() manupura.ManupuraCounty {
	return d.迈哈莱Manupura
}
    
func (d *达米埃塔DamiettaDuke) CPelusia廷尼斯() pelusia.PelusiaCounty {
	return d.廷尼斯Pelusia
}
    
var DDamietta达米埃塔 DamiettaDuke = &达米埃塔DamiettaDuke{}

func init() {
	f := DDamietta达米埃塔.(*达米埃塔DamiettaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "damietta",
		TitleName: "达米埃塔",
		TitleCode: "d_damietta",
		Counties:  map[string]feud.County{},
	}

	f.三角洲Delta = delta.CDelta三角洲
	f.三角洲Delta.SetParent(f)
	
	f.迈哈莱Manupura = manupura.CManupura迈哈莱
	f.迈哈莱Manupura.SetParent(f)
	
	f.廷尼斯Pelusia = pelusia.CPelusia廷尼斯
	f.廷尼斯Pelusia.SetParent(f)
	
}
