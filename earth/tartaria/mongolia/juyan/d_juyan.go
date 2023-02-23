package juyan

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/juyan/ejin"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/juyan/juyan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JuyanDuke interface {
	feud.Duke
	CEjin亦集乃() ejin.EjinCounty
	CJuyan居延() juyan.JuyanCounty
}

type 居延JuyanDuke struct {
	feud.BaseDuke
	亦集乃Ejin ejin.EjinCounty
	居延Juyan juyan.JuyanCounty
}

func (d *居延JuyanDuke) CEjin亦集乃() ejin.EjinCounty {
	return d.亦集乃Ejin
}

func (d *居延JuyanDuke) CJuyan居延() juyan.JuyanCounty {
	return d.居延Juyan
}

var DJuyan居延 JuyanDuke = &居延JuyanDuke{}

func init() {
	f := DJuyan居延.(*居延JuyanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "juyan",
		TitleName: "居延",
		TitleCode: "d_juyan",
		Counties:  map[string]feud.County{},
	}

	f.亦集乃Ejin = ejin.CEjin亦集乃
	f.亦集乃Ejin.SetParent(f)

	f.居延Juyan = juyan.CJuyan居延
	f.居延Juyan.SetParent(f)

}
