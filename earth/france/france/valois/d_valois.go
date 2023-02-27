package valois

import (
	"github.com/thalesfu/CK2Commands/earth/france/france/valois/ile_de_france"
	"github.com/thalesfu/CK2Commands/earth/france/france/valois/orleans"
	"github.com/thalesfu/CK2Commands/earth/france/france/valois/senlis"
	"github.com/thalesfu/CK2Commands/earth/france/france/valois/sens"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ValoisDuke interface {
    feud.Duke
    CIle_de_france法兰西岛() 	ile_de_france.Ile_de_franceCounty
    COrleans奥尔良() 	orleans.OrleansCounty
    CSenlis桑利斯() 	senlis.SenlisCounty
    CSens桑斯() 	sens.SensCounty
}

type 巴黎ValoisDuke struct {
	feud.BaseDuke
	法兰西岛Ile_de_france 	ile_de_france.Ile_de_franceCounty
	奥尔良Orleans 	orleans.OrleansCounty
	桑利斯Senlis 	senlis.SenlisCounty
	桑斯Sens 	sens.SensCounty
}

func (d *巴黎ValoisDuke) CIle_de_france法兰西岛() ile_de_france.Ile_de_franceCounty {
	return d.法兰西岛Ile_de_france
}
    
func (d *巴黎ValoisDuke) COrleans奥尔良() orleans.OrleansCounty {
	return d.奥尔良Orleans
}
    
func (d *巴黎ValoisDuke) CSenlis桑利斯() senlis.SenlisCounty {
	return d.桑利斯Senlis
}
    
func (d *巴黎ValoisDuke) CSens桑斯() sens.SensCounty {
	return d.桑斯Sens
}
    
var DValois巴黎 ValoisDuke = &巴黎ValoisDuke{}

func init() {
	f := DValois巴黎.(*巴黎ValoisDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "valois",
		TitleName: "巴黎",
		TitleCode: "d_valois",
		Counties:  map[string]feud.County{},
	}

	f.法兰西岛Ile_de_france = ile_de_france.CIle_de_france法兰西岛
	f.法兰西岛Ile_de_france.SetParent(f)
	
	f.奥尔良Orleans = orleans.COrleans奥尔良
	f.奥尔良Orleans.SetParent(f)
	
	f.桑利斯Senlis = senlis.CSenlis桑利斯
	f.桑利斯Senlis.SetParent(f)
	
	f.桑斯Sens = sens.CSens桑斯
	f.桑斯Sens.SetParent(f)
	
}
