package vengi

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/vengi/pithapuram"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/vengi/rajamahendravaram"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/vengi/vengipura"
	"github.com/thalesfu/CK2Commands/earth/deccan/andhra/vengi/vijayawada"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VengiDuke interface {
	feud.Duke
	CPithapuram毗多补罗() pithapuram.PithapuramCounty
	CRajamahendravaram罗阇摩醯因陀罗伐蓝() rajamahendravaram.RajamahendravaramCounty
	CVengipura瓶耆罗() vengipura.VengipuraCounty
	CVijayawada毗阇耶婆陀() vijayawada.VijayawadaCounty
}

type 瓶耆VengiDuke struct {
	feud.BaseDuke
	毗多补罗Pithapuram             pithapuram.PithapuramCounty
	罗阇摩醯因陀罗伐蓝Rajamahendravaram rajamahendravaram.RajamahendravaramCounty
	瓶耆罗Vengipura               vengipura.VengipuraCounty
	毗阇耶婆陀Vijayawada            vijayawada.VijayawadaCounty
}

func (d *瓶耆VengiDuke) CPithapuram毗多补罗() pithapuram.PithapuramCounty {
	return d.毗多补罗Pithapuram
}

func (d *瓶耆VengiDuke) CRajamahendravaram罗阇摩醯因陀罗伐蓝() rajamahendravaram.RajamahendravaramCounty {
	return d.罗阇摩醯因陀罗伐蓝Rajamahendravaram
}

func (d *瓶耆VengiDuke) CVengipura瓶耆罗() vengipura.VengipuraCounty {
	return d.瓶耆罗Vengipura
}

func (d *瓶耆VengiDuke) CVijayawada毗阇耶婆陀() vijayawada.VijayawadaCounty {
	return d.毗阇耶婆陀Vijayawada
}

var DVengi瓶耆 VengiDuke = &瓶耆VengiDuke{}

func init() {
	f := DVengi瓶耆.(*瓶耆VengiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "vengi",
		TitleName: "瓶耆",
		TitleCode: "d_vengi",
		Counties:  map[string]feud.County{},
	}

	f.毗多补罗Pithapuram = pithapuram.CPithapuram毗多补罗
	f.毗多补罗Pithapuram.SetParent(f)

	f.罗阇摩醯因陀罗伐蓝Rajamahendravaram = rajamahendravaram.CRajamahendravaram罗阇摩醯因陀罗伐蓝
	f.罗阇摩醯因陀罗伐蓝Rajamahendravaram.SetParent(f)

	f.瓶耆罗Vengipura = vengipura.CVengipura瓶耆罗
	f.瓶耆罗Vengipura.SetParent(f)

	f.毗阇耶婆陀Vijayawada = vijayawada.CVijayawada毗阇耶婆陀
	f.毗阇耶婆陀Vijayawada.SetParent(f)

}
