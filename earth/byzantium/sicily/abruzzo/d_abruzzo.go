package abruzzo

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/abruzzo/aprutium"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/abruzzo/chieti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbruzzoDuke interface {
	feud.Duke
	CAprutium拉奎拉() aprutium.AprutiumCounty
	CChieti基耶蒂() chieti.ChietiCounty
}

type 阿布鲁佐AbruzzoDuke struct {
	feud.BaseDuke
	拉奎拉Aprutium aprutium.AprutiumCounty
	基耶蒂Chieti   chieti.ChietiCounty
}

func (d *阿布鲁佐AbruzzoDuke) CAprutium拉奎拉() aprutium.AprutiumCounty {
	return d.拉奎拉Aprutium
}

func (d *阿布鲁佐AbruzzoDuke) CChieti基耶蒂() chieti.ChietiCounty {
	return d.基耶蒂Chieti
}

var DAbruzzo阿布鲁佐 AbruzzoDuke = &阿布鲁佐AbruzzoDuke{}

func init() {
	f := DAbruzzo阿布鲁佐.(*阿布鲁佐AbruzzoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "abruzzo",
		TitleName: "阿布鲁佐",
		TitleCode: "d_abruzzo",
		Counties:  map[string]feud.County{},
	}

	f.拉奎拉Aprutium = aprutium.CAprutium拉奎拉
	f.拉奎拉Aprutium.SetParent(f)

	f.基耶蒂Chieti = chieti.CChieti基耶蒂
	f.基耶蒂Chieti.SetParent(f)

}
