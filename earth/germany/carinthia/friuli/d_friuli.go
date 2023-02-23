package friuli

import (
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/friuli/aquileia"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/friuli/treviso"
	"github.com/thalesfu/CK2Commands/earth/germany/carinthia/friuli/udine"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FriuliDuke interface {
	feud.Duke
	CAquileia阿奎莱亚() aquileia.AquileiaCounty
	CTreviso特雷维索() treviso.TrevisoCounty
	CUdine乌迪内() udine.UdineCounty
}

type 弗留利FriuliDuke struct {
	feud.BaseDuke
	阿奎莱亚Aquileia aquileia.AquileiaCounty
	特雷维索Treviso  treviso.TrevisoCounty
	乌迪内Udine     udine.UdineCounty
}

func (d *弗留利FriuliDuke) CAquileia阿奎莱亚() aquileia.AquileiaCounty {
	return d.阿奎莱亚Aquileia
}

func (d *弗留利FriuliDuke) CTreviso特雷维索() treviso.TrevisoCounty {
	return d.特雷维索Treviso
}

func (d *弗留利FriuliDuke) CUdine乌迪内() udine.UdineCounty {
	return d.乌迪内Udine
}

var DFriuli弗留利 FriuliDuke = &弗留利FriuliDuke{}

func init() {
	f := DFriuli弗留利.(*弗留利FriuliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "friuli",
		TitleName: "弗留利",
		TitleCode: "d_friuli",
		Counties:  map[string]feud.County{},
	}

	f.阿奎莱亚Aquileia = aquileia.CAquileia阿奎莱亚
	f.阿奎莱亚Aquileia.SetParent(f)

	f.特雷维索Treviso = treviso.CTreviso特雷维索
	f.特雷维索Treviso.SetParent(f)

	f.乌迪内Udine = udine.CUdine乌迪内
	f.乌迪内Udine.SetParent(f)

}
