package abakan

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/abakan/abakan"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/abakan/erchis"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/abakan/kuznetsk"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/abakan/tom"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AbakanDuke interface {
	feud.Duke
	CAbakan阿巴坎() abakan.AbakanCounty
	CErchis克季() erchis.ErchisCounty
	CKuznetsk库兹涅茨克() kuznetsk.KuznetskCounty
	CTom托木() tom.TomCounty
}

type 阿巴坎AbakanDuke struct {
	feud.BaseDuke
	阿巴坎Abakan     abakan.AbakanCounty
	克季Erchis      erchis.ErchisCounty
	库兹涅茨克Kuznetsk kuznetsk.KuznetskCounty
	托木Tom         tom.TomCounty
}

func (d *阿巴坎AbakanDuke) CAbakan阿巴坎() abakan.AbakanCounty {
	return d.阿巴坎Abakan
}

func (d *阿巴坎AbakanDuke) CErchis克季() erchis.ErchisCounty {
	return d.克季Erchis
}

func (d *阿巴坎AbakanDuke) CKuznetsk库兹涅茨克() kuznetsk.KuznetskCounty {
	return d.库兹涅茨克Kuznetsk
}

func (d *阿巴坎AbakanDuke) CTom托木() tom.TomCounty {
	return d.托木Tom
}

var DAbakan阿巴坎 AbakanDuke = &阿巴坎AbakanDuke{}

func init() {
	f := DAbakan阿巴坎.(*阿巴坎AbakanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "abakan",
		TitleName: "阿巴坎",
		TitleCode: "d_abakan",
		Counties:  map[string]feud.County{},
	}

	f.阿巴坎Abakan = abakan.CAbakan阿巴坎
	f.阿巴坎Abakan.SetParent(f)

	f.克季Erchis = erchis.CErchis克季
	f.克季Erchis.SetParent(f)

	f.库兹涅茨克Kuznetsk = kuznetsk.CKuznetsk库兹涅茨克
	f.库兹涅茨克Kuznetsk.SetParent(f)

	f.托木Tom = tom.CTom托木
	f.托木Tom.SetParent(f)

}
