package berry

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/berry/bourges"
	"github.com/thalesfu/CK2Commands/earth/france/france/berry/sancerre"
	"github.com/thalesfu/CK2Commands/earth/france/france/berry/tourraine"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BerryDuke interface {
    feud.Duke
    CBourges布尔日() 	bourges.BourgesCounty
    CSancerre桑塞尔() 	sancerre.SancerreCounty
    CTourraine图尔() 	tourraine.TourraineCounty
}

type 贝里BerryDuke struct {
	feud.BaseDuke
	布尔日Bourges 	bourges.BourgesCounty
	桑塞尔Sancerre 	sancerre.SancerreCounty
	图尔Tourraine 	tourraine.TourraineCounty
}

func (d *贝里BerryDuke) CBourges布尔日() bourges.BourgesCounty {
	return d.布尔日Bourges
}
    
func (d *贝里BerryDuke) CSancerre桑塞尔() sancerre.SancerreCounty {
	return d.桑塞尔Sancerre
}
    
func (d *贝里BerryDuke) CTourraine图尔() tourraine.TourraineCounty {
	return d.图尔Tourraine
}
    
var DBerry贝里 BerryDuke = &贝里BerryDuke{}

func init() {
	f := DBerry贝里.(*贝里BerryDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "berry",
		TitleName: "贝里",
		TitleCode: "d_berry",
		Counties:  map[string]feud.County{},
	}

	f.布尔日Bourges = bourges.CBourges布尔日
	f.布尔日Bourges.SetParent(f)
	
	f.桑塞尔Sancerre = sancerre.CSancerre桑塞尔
	f.桑塞尔Sancerre.SetParent(f)
	
	f.图尔Tourraine = tourraine.CTourraine图尔
	f.图尔Tourraine.SetParent(f)
	
}
