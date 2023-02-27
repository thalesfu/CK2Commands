package spoleto

import (
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/spoleto/perugia"
	"github.com/thalesfu/CK2Commands/earth/italy/papacy/spoleto/spoleto"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SpoletoDuke interface {
    feud.Duke
    CPerugia佩鲁贾() 	perugia.PerugiaCounty
    CSpoleto斯波莱托() 	spoleto.SpoletoCounty
}

type 斯波莱托SpoletoDuke struct {
	feud.BaseDuke
	佩鲁贾Perugia 	perugia.PerugiaCounty
	斯波莱托Spoleto 	spoleto.SpoletoCounty
}

func (d *斯波莱托SpoletoDuke) CPerugia佩鲁贾() perugia.PerugiaCounty {
	return d.佩鲁贾Perugia
}
    
func (d *斯波莱托SpoletoDuke) CSpoleto斯波莱托() spoleto.SpoletoCounty {
	return d.斯波莱托Spoleto
}
    
var DSpoleto斯波莱托 SpoletoDuke = &斯波莱托SpoletoDuke{}

func init() {
	f := DSpoleto斯波莱托.(*斯波莱托SpoletoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "spoleto",
		TitleName: "斯波莱托",
		TitleCode: "d_spoleto",
		Counties:  map[string]feud.County{},
	}

	f.佩鲁贾Perugia = perugia.CPerugia佩鲁贾
	f.佩鲁贾Perugia.SetParent(f)
	
	f.斯波莱托Spoleto = spoleto.CSpoleto斯波莱托
	f.斯波莱托Spoleto.SetParent(f)
	
}
