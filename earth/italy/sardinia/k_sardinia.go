package sardinia

import (
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/corsica"
	"github.com/thalesfu/CK2Commands/earth/italy/sardinia/sardinia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SardiniaKingdom interface {
    feud.Kingdom
    DCorsica科西嘉() 	corsica.CorsicaDuke
    DSardinia撒丁() 	sardinia.SardiniaDuke
}

type 撒丁与科西嘉SardiniaKingdom struct {
	feud.BaseKingdom
	科西嘉Corsica 	corsica.CorsicaDuke
	撒丁Sardinia 	sardinia.SardiniaDuke
}

func (k *撒丁与科西嘉SardiniaKingdom) DCorsica科西嘉() corsica.CorsicaDuke {
	return k.科西嘉Corsica
}
    
func (k *撒丁与科西嘉SardiniaKingdom) DSardinia撒丁() sardinia.SardiniaDuke {
	return k.撒丁Sardinia
}
    
var KSardinia撒丁与科西嘉 SardiniaKingdom = &撒丁与科西嘉SardiniaKingdom{}

func init() {
	f := KSardinia撒丁与科西嘉.(*撒丁与科西嘉SardiniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sardinia",
		TitleName: "撒丁与科西嘉",
		TitleCode: "k_sardinia",
		Dukes:  map[string]feud.Duke{},
	}

	f.科西嘉Corsica = corsica.DCorsica科西嘉
	f.科西嘉Corsica.SetParent(f)
	
	f.撒丁Sardinia = sardinia.DSardinia撒丁
	f.撒丁Sardinia.SetParent(f)
	
}
