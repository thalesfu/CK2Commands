package burgundy

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/auxerre"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/chalons"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/charolais"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/dijon"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/macon"
	"github.com/thalesfu/CK2Commands/earth/france/france/burgundy/nevers"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurgundyDuke interface {
    feud.Duke
    CAuxerre欧塞尔() 	auxerre.AuxerreCounty
    CChalons沙隆() 	chalons.ChalonsCounty
    CCharolais沙罗莱() 	charolais.CharolaisCounty
    CDijon第戎() 	dijon.DijonCounty
    CMacon马孔() 	macon.MaconCounty
    CNevers讷韦尔() 	nevers.NeversCounty
}

type 勃艮第BurgundyDuke struct {
	feud.BaseDuke
	欧塞尔Auxerre 	auxerre.AuxerreCounty
	沙隆Chalons 	chalons.ChalonsCounty
	沙罗莱Charolais 	charolais.CharolaisCounty
	第戎Dijon 	dijon.DijonCounty
	马孔Macon 	macon.MaconCounty
	讷韦尔Nevers 	nevers.NeversCounty
}

func (d *勃艮第BurgundyDuke) CAuxerre欧塞尔() auxerre.AuxerreCounty {
	return d.欧塞尔Auxerre
}
    
func (d *勃艮第BurgundyDuke) CChalons沙隆() chalons.ChalonsCounty {
	return d.沙隆Chalons
}
    
func (d *勃艮第BurgundyDuke) CCharolais沙罗莱() charolais.CharolaisCounty {
	return d.沙罗莱Charolais
}
    
func (d *勃艮第BurgundyDuke) CDijon第戎() dijon.DijonCounty {
	return d.第戎Dijon
}
    
func (d *勃艮第BurgundyDuke) CMacon马孔() macon.MaconCounty {
	return d.马孔Macon
}
    
func (d *勃艮第BurgundyDuke) CNevers讷韦尔() nevers.NeversCounty {
	return d.讷韦尔Nevers
}
    
var DBurgundy勃艮第 BurgundyDuke = &勃艮第BurgundyDuke{}

func init() {
	f := DBurgundy勃艮第.(*勃艮第BurgundyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "burgundy",
		TitleName: "勃艮第",
		TitleCode: "d_burgundy",
		Counties:  map[string]feud.County{},
	}

	f.欧塞尔Auxerre = auxerre.CAuxerre欧塞尔
	f.欧塞尔Auxerre.SetParent(f)
	
	f.沙隆Chalons = chalons.CChalons沙隆
	f.沙隆Chalons.SetParent(f)
	
	f.沙罗莱Charolais = charolais.CCharolais沙罗莱
	f.沙罗莱Charolais.SetParent(f)
	
	f.第戎Dijon = dijon.CDijon第戎
	f.第戎Dijon.SetParent(f)
	
	f.马孔Macon = macon.CMacon马孔
	f.马孔Macon.SetParent(f)
	
	f.讷韦尔Nevers = nevers.CNevers讷韦尔
	f.讷韦尔Nevers.SetParent(f)
	
}
