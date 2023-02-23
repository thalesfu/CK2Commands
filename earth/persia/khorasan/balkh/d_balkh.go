package balkh

import (
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/balkh/balkh"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/balkh/guzgan"
	"github.com/thalesfu/CK2Commands/earth/persia/khorasan/balkh/maymana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BalkhDuke interface {
	feud.Duke
	CBalkh缚喝() balkh.BalkhCounty
	CGuzgan胡实健() guzgan.GuzganCounty
	CMaymana梅马内() maymana.MaymanaCounty
}

type 缚喝BalkhDuke struct {
	feud.BaseDuke
	缚喝Balkh    balkh.BalkhCounty
	胡实健Guzgan  guzgan.GuzganCounty
	梅马内Maymana maymana.MaymanaCounty
}

func (d *缚喝BalkhDuke) CBalkh缚喝() balkh.BalkhCounty {
	return d.缚喝Balkh
}

func (d *缚喝BalkhDuke) CGuzgan胡实健() guzgan.GuzganCounty {
	return d.胡实健Guzgan
}

func (d *缚喝BalkhDuke) CMaymana梅马内() maymana.MaymanaCounty {
	return d.梅马内Maymana
}

var DBalkh缚喝 BalkhDuke = &缚喝BalkhDuke{}

func init() {
	f := DBalkh缚喝.(*缚喝BalkhDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "balkh",
		TitleName: "缚喝",
		TitleCode: "d_balkh",
		Counties:  map[string]feud.County{},
	}

	f.缚喝Balkh = balkh.CBalkh缚喝
	f.缚喝Balkh.SetParent(f)

	f.胡实健Guzgan = guzgan.CGuzgan胡实健
	f.胡实健Guzgan.SetParent(f)

	f.梅马内Maymana = maymana.CMaymana梅马内
	f.梅马内Maymana.SetParent(f)

}
