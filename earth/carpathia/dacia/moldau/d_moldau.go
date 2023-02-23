package moldau

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/moldau/bacau"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/moldau/iasi"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/moldau/peresechen"
	"github.com/thalesfu/CK2Commands/earth/carpathia/dacia/moldau/torki"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MoldauDuke interface {
	feud.Duke
	CBacau巴克乌() bacau.BacauCounty
	CIasi雅西() iasi.IasiCounty
	CPeresechen佩列塞琴() peresechen.PeresechenCounty
	CTorki苏恰瓦() torki.TorkiCounty
}

type 摩尔达维亚MoldauDuke struct {
	feud.BaseDuke
	巴克乌Bacau       bacau.BacauCounty
	雅西Iasi         iasi.IasiCounty
	佩列塞琴Peresechen peresechen.PeresechenCounty
	苏恰瓦Torki       torki.TorkiCounty
}

func (d *摩尔达维亚MoldauDuke) CBacau巴克乌() bacau.BacauCounty {
	return d.巴克乌Bacau
}

func (d *摩尔达维亚MoldauDuke) CIasi雅西() iasi.IasiCounty {
	return d.雅西Iasi
}

func (d *摩尔达维亚MoldauDuke) CPeresechen佩列塞琴() peresechen.PeresechenCounty {
	return d.佩列塞琴Peresechen
}

func (d *摩尔达维亚MoldauDuke) CTorki苏恰瓦() torki.TorkiCounty {
	return d.苏恰瓦Torki
}

var DMoldau摩尔达维亚 MoldauDuke = &摩尔达维亚MoldauDuke{}

func init() {
	f := DMoldau摩尔达维亚.(*摩尔达维亚MoldauDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "moldau",
		TitleName: "摩尔达维亚",
		TitleCode: "d_moldau",
		Counties:  map[string]feud.County{},
	}

	f.巴克乌Bacau = bacau.CBacau巴克乌
	f.巴克乌Bacau.SetParent(f)

	f.雅西Iasi = iasi.CIasi雅西
	f.雅西Iasi.SetParent(f)

	f.佩列塞琴Peresechen = peresechen.CPeresechen佩列塞琴
	f.佩列塞琴Peresechen.SetParent(f)

	f.苏恰瓦Torki = torki.CTorki苏恰瓦
	f.苏恰瓦Torki.SetParent(f)

}
