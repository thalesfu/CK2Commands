package otuken

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/otuken/egiin"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/otuken/gorgol"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/otuken/khovsgol"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/otuken/otuken"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OtukenDuke interface {
	feud.Duke
	CEgiin额金() egiin.EgiinCounty
	CGorgol叶尼塞() gorgol.GorgolCounty
	CKhovsgol库苏古尔() khovsgol.KhovsgolCounty
	COtuken于都斤() otuken.OtukenCounty
}

type 于都斤OtukenDuke struct {
	feud.BaseDuke
	额金Egiin      egiin.EgiinCounty
	叶尼塞Gorgol    gorgol.GorgolCounty
	库苏古尔Khovsgol khovsgol.KhovsgolCounty
	于都斤Otuken    otuken.OtukenCounty
}

func (d *于都斤OtukenDuke) CEgiin额金() egiin.EgiinCounty {
	return d.额金Egiin
}

func (d *于都斤OtukenDuke) CGorgol叶尼塞() gorgol.GorgolCounty {
	return d.叶尼塞Gorgol
}

func (d *于都斤OtukenDuke) CKhovsgol库苏古尔() khovsgol.KhovsgolCounty {
	return d.库苏古尔Khovsgol
}

func (d *于都斤OtukenDuke) COtuken于都斤() otuken.OtukenCounty {
	return d.于都斤Otuken
}

var DOtuken于都斤 OtukenDuke = &于都斤OtukenDuke{}

func init() {
	f := DOtuken于都斤.(*于都斤OtukenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "otuken",
		TitleName: "于都斤",
		TitleCode: "d_otuken",
		Counties:  map[string]feud.County{},
	}

	f.额金Egiin = egiin.CEgiin额金
	f.额金Egiin.SetParent(f)

	f.叶尼塞Gorgol = gorgol.CGorgol叶尼塞
	f.叶尼塞Gorgol.SetParent(f)

	f.库苏古尔Khovsgol = khovsgol.CKhovsgol库苏古尔
	f.库苏古尔Khovsgol.SetParent(f)

	f.于都斤Otuken = otuken.COtuken于都斤
	f.于都斤Otuken.SetParent(f)

}
