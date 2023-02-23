package connacht

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/connacht/breifne"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/connacht/connacht"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ConnachtDuke interface {
	feud.Duke
	CBreifne布雷夫讷() breifne.BreifneCounty
	CConnacht西康诺特() connacht.ConnachtCounty
}

type 康诺特ConnachtDuke struct {
	feud.BaseDuke
	布雷夫讷Breifne  breifne.BreifneCounty
	西康诺特Connacht connacht.ConnachtCounty
}

func (d *康诺特ConnachtDuke) CBreifne布雷夫讷() breifne.BreifneCounty {
	return d.布雷夫讷Breifne
}

func (d *康诺特ConnachtDuke) CConnacht西康诺特() connacht.ConnachtCounty {
	return d.西康诺特Connacht
}

var DConnacht康诺特 ConnachtDuke = &康诺特ConnachtDuke{}

func init() {
	f := DConnacht康诺特.(*康诺特ConnachtDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "connacht",
		TitleName: "康诺特",
		TitleCode: "d_connacht",
		Counties:  map[string]feud.County{},
	}

	f.布雷夫讷Breifne = breifne.CBreifne布雷夫讷
	f.布雷夫讷Breifne.SetParent(f)

	f.西康诺特Connacht = connacht.CConnacht西康诺特
	f.西康诺特Connacht.SetParent(f)

}
