package canterbury

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/canterbury/kent"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/canterbury/surrey"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/canterbury/sussex"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CanterburyDuke interface {
	feud.Duke
	CKent肯特() kent.KentCounty
	CSurrey萨里() surrey.SurreyCounty
	CSussex萨塞克斯() sussex.SussexCounty
}

type 肯特CanterburyDuke struct {
	feud.BaseDuke
	肯特Kent     kent.KentCounty
	萨里Surrey   surrey.SurreyCounty
	萨塞克斯Sussex sussex.SussexCounty
}

func (d *肯特CanterburyDuke) CKent肯特() kent.KentCounty {
	return d.肯特Kent
}

func (d *肯特CanterburyDuke) CSurrey萨里() surrey.SurreyCounty {
	return d.萨里Surrey
}

func (d *肯特CanterburyDuke) CSussex萨塞克斯() sussex.SussexCounty {
	return d.萨塞克斯Sussex
}

var DCanterbury肯特 CanterburyDuke = &肯特CanterburyDuke{}

func init() {
	f := DCanterbury肯特.(*肯特CanterburyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "canterbury",
		TitleName: "肯特",
		TitleCode: "d_canterbury",
		Counties:  map[string]feud.County{},
	}

	f.肯特Kent = kent.CKent肯特
	f.肯特Kent.SetParent(f)

	f.萨里Surrey = surrey.CSurrey萨里
	f.萨里Surrey.SetParent(f)

	f.萨塞克斯Sussex = sussex.CSussex萨塞克斯
	f.萨塞克斯Sussex.SetParent(f)

}
