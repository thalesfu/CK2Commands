package sibir

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/sibir/osha"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/sibir/sibir"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/sibir/tura"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir/sibir/vagay"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SibirDuke interface {
	feud.Duke
	COsha奥沙() osha.OshaCounty
	CSibir失必儿() sibir.SibirCounty
	CTura图拉() tura.TuraCounty
	CVagay瓦盖() vagay.VagayCounty
}

type 失必儿SibirDuke struct {
	feud.BaseDuke
	奥沙Osha   osha.OshaCounty
	失必儿Sibir sibir.SibirCounty
	图拉Tura   tura.TuraCounty
	瓦盖Vagay  vagay.VagayCounty
}

func (d *失必儿SibirDuke) COsha奥沙() osha.OshaCounty {
	return d.奥沙Osha
}

func (d *失必儿SibirDuke) CSibir失必儿() sibir.SibirCounty {
	return d.失必儿Sibir
}

func (d *失必儿SibirDuke) CTura图拉() tura.TuraCounty {
	return d.图拉Tura
}

func (d *失必儿SibirDuke) CVagay瓦盖() vagay.VagayCounty {
	return d.瓦盖Vagay
}

var DSibir失必儿 SibirDuke = &失必儿SibirDuke{}

func init() {
	f := DSibir失必儿.(*失必儿SibirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sibir",
		TitleName: "失必儿",
		TitleCode: "d_sibir",
		Counties:  map[string]feud.County{},
	}

	f.奥沙Osha = osha.COsha奥沙
	f.奥沙Osha.SetParent(f)

	f.失必儿Sibir = sibir.CSibir失必儿
	f.失必儿Sibir.SetParent(f)

	f.图拉Tura = tura.CTura图拉
	f.图拉Tura.SetParent(f)

	f.瓦盖Vagay = vagay.CVagay瓦盖
	f.瓦盖Vagay.SetParent(f)

}
