package benevento

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/benevento/benevento"
	"github.com/thalesfu/CK2Commands/earth/byzantium/sicily/benevento/foggia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeneventoDuke interface {
	feud.Duke
	CBenevento贝内文托() benevento.BeneventoCounty
	CFoggia福贾() foggia.FoggiaCounty
}

type 贝内文托BeneventoDuke struct {
	feud.BaseDuke
	贝内文托Benevento benevento.BeneventoCounty
	福贾Foggia      foggia.FoggiaCounty
}

func (d *贝内文托BeneventoDuke) CBenevento贝内文托() benevento.BeneventoCounty {
	return d.贝内文托Benevento
}

func (d *贝内文托BeneventoDuke) CFoggia福贾() foggia.FoggiaCounty {
	return d.福贾Foggia
}

var DBenevento贝内文托 BeneventoDuke = &贝内文托BeneventoDuke{}

func init() {
	f := DBenevento贝内文托.(*贝内文托BeneventoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "benevento",
		TitleName: "贝内文托",
		TitleCode: "d_benevento",
		Counties:  map[string]feud.County{},
	}

	f.贝内文托Benevento = benevento.CBenevento贝内文托
	f.贝内文托Benevento.SetParent(f)

	f.福贾Foggia = foggia.CFoggia福贾
	f.福贾Foggia.SetParent(f)

}
