package syr_darya

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/syr_darya/chach"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/syr_darya/otrar"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/syr_darya/syr_darya"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/syr_darya/yangikent"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Syr_daryaDuke interface {
    feud.Duke
    CChach赭石() 	chach.ChachCounty
    COtrar讹答剌() 	otrar.OtrarCounty
    CSyr_darya昔格纳黑() 	syr_darya.Syr_daryaCounty
    CYangikent养吉干() 	yangikent.YangikentCounty
}

type 锡尔河Syr_daryaDuke struct {
	feud.BaseDuke
	赭石Chach 	chach.ChachCounty
	讹答剌Otrar 	otrar.OtrarCounty
	昔格纳黑Syr_darya 	syr_darya.Syr_daryaCounty
	养吉干Yangikent 	yangikent.YangikentCounty
}

func (d *锡尔河Syr_daryaDuke) CChach赭石() chach.ChachCounty {
	return d.赭石Chach
}
    
func (d *锡尔河Syr_daryaDuke) COtrar讹答剌() otrar.OtrarCounty {
	return d.讹答剌Otrar
}
    
func (d *锡尔河Syr_daryaDuke) CSyr_darya昔格纳黑() syr_darya.Syr_daryaCounty {
	return d.昔格纳黑Syr_darya
}
    
func (d *锡尔河Syr_daryaDuke) CYangikent养吉干() yangikent.YangikentCounty {
	return d.养吉干Yangikent
}
    
var DSyr_darya锡尔河 Syr_daryaDuke = &锡尔河Syr_daryaDuke{}

func init() {
	f := DSyr_darya锡尔河.(*锡尔河Syr_daryaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "syr_darya",
		TitleName: "锡尔河",
		TitleCode: "d_syr_darya",
		Counties:  map[string]feud.County{},
	}

	f.赭石Chach = chach.CChach赭石
	f.赭石Chach.SetParent(f)
	
	f.讹答剌Otrar = otrar.COtrar讹答剌
	f.讹答剌Otrar.SetParent(f)
	
	f.昔格纳黑Syr_darya = syr_darya.CSyr_darya昔格纳黑
	f.昔格纳黑Syr_darya.SetParent(f)
	
	f.养吉干Yangikent = yangikent.CYangikent养吉干
	f.养吉干Yangikent.SetParent(f)
	
}
