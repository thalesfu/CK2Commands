package minsk

import (
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/minsk/drutsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/minsk/minsk"
	"github.com/thalesfu/CK2Commands/earth/russia/ruthenia/minsk/novogrodek"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MinskDuke interface {
	feud.Duke
	CDrutsk德鲁茨克() drutsk.DrutskCounty
	CMinsk明斯克() minsk.MinskCounty
	CNovogrodek新格鲁多克() novogrodek.NovogrodekCounty
}

type 明斯克MinskDuke struct {
	feud.BaseDuke
	德鲁茨克Drutsk      drutsk.DrutskCounty
	明斯克Minsk        minsk.MinskCounty
	新格鲁多克Novogrodek novogrodek.NovogrodekCounty
}

func (d *明斯克MinskDuke) CDrutsk德鲁茨克() drutsk.DrutskCounty {
	return d.德鲁茨克Drutsk
}

func (d *明斯克MinskDuke) CMinsk明斯克() minsk.MinskCounty {
	return d.明斯克Minsk
}

func (d *明斯克MinskDuke) CNovogrodek新格鲁多克() novogrodek.NovogrodekCounty {
	return d.新格鲁多克Novogrodek
}

var DMinsk明斯克 MinskDuke = &明斯克MinskDuke{}

func init() {
	f := DMinsk明斯克.(*明斯克MinskDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "minsk",
		TitleName: "明斯克",
		TitleCode: "d_minsk",
		Counties:  map[string]feud.County{},
	}

	f.德鲁茨克Drutsk = drutsk.CDrutsk德鲁茨克
	f.德鲁茨克Drutsk.SetParent(f)

	f.明斯克Minsk = minsk.CMinsk明斯克
	f.明斯克Minsk.SetParent(f)

	f.新格鲁多克Novogrodek = novogrodek.CNovogrodek新格鲁多克
	f.新格鲁多克Novogrodek.SetParent(f)

}
