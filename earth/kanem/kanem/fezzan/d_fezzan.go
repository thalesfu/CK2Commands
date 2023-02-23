package fezzan

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/fezzan/djado"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/fezzan/murzuk"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/fezzan/zawila"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FezzanDuke interface {
	feud.Duke
	CDjado贾多() djado.DjadoCounty
	CMurzuk迈尔祖格() murzuk.MurzukCounty
	CZawila宰维莱() zawila.ZawilaCounty
}

type 费赞FezzanDuke struct {
	feud.BaseDuke
	贾多Djado    djado.DjadoCounty
	迈尔祖格Murzuk murzuk.MurzukCounty
	宰维莱Zawila  zawila.ZawilaCounty
}

func (d *费赞FezzanDuke) CDjado贾多() djado.DjadoCounty {
	return d.贾多Djado
}

func (d *费赞FezzanDuke) CMurzuk迈尔祖格() murzuk.MurzukCounty {
	return d.迈尔祖格Murzuk
}

func (d *费赞FezzanDuke) CZawila宰维莱() zawila.ZawilaCounty {
	return d.宰维莱Zawila
}

var DFezzan费赞 FezzanDuke = &费赞FezzanDuke{}

func init() {
	f := DFezzan费赞.(*费赞FezzanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "fezzan",
		TitleName: "费赞",
		TitleCode: "d_fezzan",
		Counties:  map[string]feud.County{},
	}

	f.贾多Djado = djado.CDjado贾多
	f.贾多Djado.SetParent(f)

	f.迈尔祖格Murzuk = murzuk.CMurzuk迈尔祖格
	f.迈尔祖格Murzuk.SetParent(f)

	f.宰维莱Zawila = zawila.CZawila宰维莱
	f.宰维莱Zawila.SetParent(f)

}
