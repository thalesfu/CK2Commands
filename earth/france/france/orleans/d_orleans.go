package orleans

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/orleans/blois"
	"github.com/thalesfu/CK2Commands/earth/france/france/orleans/chartres"
	"github.com/thalesfu/CK2Commands/earth/france/france/orleans/dunois"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrleansDuke interface {
    feud.Duke
    CBlois布卢瓦() 	blois.BloisCounty
    CChartres沙特尔() 	chartres.ChartresCounty
    CDunois迪努瓦() 	dunois.DunoisCounty
}

type 布卢瓦OrleansDuke struct {
	feud.BaseDuke
	布卢瓦Blois 	blois.BloisCounty
	沙特尔Chartres 	chartres.ChartresCounty
	迪努瓦Dunois 	dunois.DunoisCounty
}

func (d *布卢瓦OrleansDuke) CBlois布卢瓦() blois.BloisCounty {
	return d.布卢瓦Blois
}
    
func (d *布卢瓦OrleansDuke) CChartres沙特尔() chartres.ChartresCounty {
	return d.沙特尔Chartres
}
    
func (d *布卢瓦OrleansDuke) CDunois迪努瓦() dunois.DunoisCounty {
	return d.迪努瓦Dunois
}
    
var DOrleans布卢瓦 OrleansDuke = &布卢瓦OrleansDuke{}

func init() {
	f := DOrleans布卢瓦.(*布卢瓦OrleansDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "orleans",
		TitleName: "布卢瓦",
		TitleCode: "d_orleans",
		Counties:  map[string]feud.County{},
	}

	f.布卢瓦Blois = blois.CBlois布卢瓦
	f.布卢瓦Blois.SetParent(f)
	
	f.沙特尔Chartres = chartres.CChartres沙特尔
	f.沙特尔Chartres.SetParent(f)
	
	f.迪努瓦Dunois = dunois.CDunois迪努瓦
	f.迪努瓦Dunois.SetParent(f)
	
}
