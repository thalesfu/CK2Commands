package axum

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/axum/akordat"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/axum/aksum"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/axum/massawa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AxumDuke interface {
    feud.Duke
    CAkordat阿科达特() 	akordat.AkordatCounty
    CAksum阿克苏姆() 	aksum.AksumCounty
    CMassawa马萨瓦() 	massawa.MassawaCounty
}

type 阿克苏姆AxumDuke struct {
	feud.BaseDuke
	阿科达特Akordat 	akordat.AkordatCounty
	阿克苏姆Aksum 	aksum.AksumCounty
	马萨瓦Massawa 	massawa.MassawaCounty
}

func (d *阿克苏姆AxumDuke) CAkordat阿科达特() akordat.AkordatCounty {
	return d.阿科达特Akordat
}
    
func (d *阿克苏姆AxumDuke) CAksum阿克苏姆() aksum.AksumCounty {
	return d.阿克苏姆Aksum
}
    
func (d *阿克苏姆AxumDuke) CMassawa马萨瓦() massawa.MassawaCounty {
	return d.马萨瓦Massawa
}
    
var DAxum阿克苏姆 AxumDuke = &阿克苏姆AxumDuke{}

func init() {
	f := DAxum阿克苏姆.(*阿克苏姆AxumDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "axum",
		TitleName: "阿克苏姆",
		TitleCode: "d_axum",
		Counties:  map[string]feud.County{},
	}

	f.阿科达特Akordat = akordat.CAkordat阿科达特
	f.阿科达特Akordat.SetParent(f)
	
	f.阿克苏姆Aksum = aksum.CAksum阿克苏姆
	f.阿克苏姆Aksum.SetParent(f)
	
	f.马萨瓦Massawa = massawa.CMassawa马萨瓦
	f.马萨瓦Massawa.SetParent(f)
	
}
