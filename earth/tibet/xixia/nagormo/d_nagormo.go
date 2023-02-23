package nagormo

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nagormo/lenghu"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nagormo/nagormo"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nagormo/qaidam"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nagormo/tanggula"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagormoDuke interface {
	feud.Duke
	CLenghu冷湖() lenghu.LenghuCounty
	CNagormo格尔木() nagormo.NagormoCounty
	CQaidam柴旦() qaidam.QaidamCounty
	CTanggula唐古拉() tanggula.TanggulaCounty
}

type 格尔木NagormoDuke struct {
	feud.BaseDuke
	冷湖Lenghu    lenghu.LenghuCounty
	格尔木Nagormo  nagormo.NagormoCounty
	柴旦Qaidam    qaidam.QaidamCounty
	唐古拉Tanggula tanggula.TanggulaCounty
}

func (d *格尔木NagormoDuke) CLenghu冷湖() lenghu.LenghuCounty {
	return d.冷湖Lenghu
}

func (d *格尔木NagormoDuke) CNagormo格尔木() nagormo.NagormoCounty {
	return d.格尔木Nagormo
}

func (d *格尔木NagormoDuke) CQaidam柴旦() qaidam.QaidamCounty {
	return d.柴旦Qaidam
}

func (d *格尔木NagormoDuke) CTanggula唐古拉() tanggula.TanggulaCounty {
	return d.唐古拉Tanggula
}

var DNagormo格尔木 NagormoDuke = &格尔木NagormoDuke{}

func init() {
	f := DNagormo格尔木.(*格尔木NagormoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nagormo",
		TitleName: "格尔木",
		TitleCode: "d_nagormo",
		Counties:  map[string]feud.County{},
	}

	f.冷湖Lenghu = lenghu.CLenghu冷湖
	f.冷湖Lenghu.SetParent(f)

	f.格尔木Nagormo = nagormo.CNagormo格尔木
	f.格尔木Nagormo.SetParent(f)

	f.柴旦Qaidam = qaidam.CQaidam柴旦
	f.柴旦Qaidam.SetParent(f)

	f.唐古拉Tanggula = tanggula.CTanggula唐古拉
	f.唐古拉Tanggula.SetParent(f)

}
