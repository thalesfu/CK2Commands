package nefoud

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud/dariya"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud/hail"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud/hajr"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud/halaban"
	"github.com/thalesfu/CK2Commands/earth/arabia/arabia/nefoud/rafha"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NefoudDuke interface {
    feud.Duke
    CDariya德拉伊耶() 	dariya.DariyaCounty
    CHail哈伊勒() 	hail.HailCounty
    CHajr哈杰尔() 	hajr.HajrCounty
    CHalaban哈莱班() 	halaban.HalabanCounty
    CRafha拉夫哈() 	rafha.RafhaCounty
}

type 内夫得NefoudDuke struct {
	feud.BaseDuke
	德拉伊耶Dariya 	dariya.DariyaCounty
	哈伊勒Hail 	hail.HailCounty
	哈杰尔Hajr 	hajr.HajrCounty
	哈莱班Halaban 	halaban.HalabanCounty
	拉夫哈Rafha 	rafha.RafhaCounty
}

func (d *内夫得NefoudDuke) CDariya德拉伊耶() dariya.DariyaCounty {
	return d.德拉伊耶Dariya
}
    
func (d *内夫得NefoudDuke) CHail哈伊勒() hail.HailCounty {
	return d.哈伊勒Hail
}
    
func (d *内夫得NefoudDuke) CHajr哈杰尔() hajr.HajrCounty {
	return d.哈杰尔Hajr
}
    
func (d *内夫得NefoudDuke) CHalaban哈莱班() halaban.HalabanCounty {
	return d.哈莱班Halaban
}
    
func (d *内夫得NefoudDuke) CRafha拉夫哈() rafha.RafhaCounty {
	return d.拉夫哈Rafha
}
    
var DNefoud内夫得 NefoudDuke = &内夫得NefoudDuke{}

func init() {
	f := DNefoud内夫得.(*内夫得NefoudDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nefoud",
		TitleName: "内夫得",
		TitleCode: "d_nefoud",
		Counties:  map[string]feud.County{},
	}

	f.德拉伊耶Dariya = dariya.CDariya德拉伊耶
	f.德拉伊耶Dariya.SetParent(f)
	
	f.哈伊勒Hail = hail.CHail哈伊勒
	f.哈伊勒Hail.SetParent(f)
	
	f.哈杰尔Hajr = hajr.CHajr哈杰尔
	f.哈杰尔Hajr.SetParent(f)
	
	f.哈莱班Halaban = halaban.CHalaban哈莱班
	f.哈莱班Halaban.SetParent(f)
	
	f.拉夫哈Rafha = rafha.CRafha拉夫哈
	f.拉夫哈Rafha.SetParent(f)
	
}
