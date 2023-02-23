package tlemcen

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tlemcen/figuig"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tlemcen/snassen"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/tlemcen/tlemcen"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TlemcenDuke interface {
	feud.Duke
	CFiguig坦德拉拉() figuig.FiguigCounty
	CSnassen瓦伊达() snassen.SnassenCounty
	CTlemcen特莱姆森() tlemcen.TlemcenCounty
}

type 特莱姆森TlemcenDuke struct {
	feud.BaseDuke
	坦德拉拉Figuig  figuig.FiguigCounty
	瓦伊达Snassen  snassen.SnassenCounty
	特莱姆森Tlemcen tlemcen.TlemcenCounty
}

func (d *特莱姆森TlemcenDuke) CFiguig坦德拉拉() figuig.FiguigCounty {
	return d.坦德拉拉Figuig
}

func (d *特莱姆森TlemcenDuke) CSnassen瓦伊达() snassen.SnassenCounty {
	return d.瓦伊达Snassen
}

func (d *特莱姆森TlemcenDuke) CTlemcen特莱姆森() tlemcen.TlemcenCounty {
	return d.特莱姆森Tlemcen
}

var DTlemcen特莱姆森 TlemcenDuke = &特莱姆森TlemcenDuke{}

func init() {
	f := DTlemcen特莱姆森.(*特莱姆森TlemcenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tlemcen",
		TitleName: "特莱姆森",
		TitleCode: "d_tlemcen",
		Counties:  map[string]feud.County{},
	}

	f.坦德拉拉Figuig = figuig.CFiguig坦德拉拉
	f.坦德拉拉Figuig.SetParent(f)

	f.瓦伊达Snassen = snassen.CSnassen瓦伊达
	f.瓦伊达Snassen.SetParent(f)

	f.特莱姆森Tlemcen = tlemcen.CTlemcen特莱姆森
	f.特莱姆森Tlemcen.SetParent(f)

}
