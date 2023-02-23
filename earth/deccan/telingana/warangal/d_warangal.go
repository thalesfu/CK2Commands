package warangal

import (
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/warangal/balkonda"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/warangal/kambampet"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/warangal/orangallu"
	"github.com/thalesfu/CK2Commands/earth/deccan/telingana/warangal/vemulavada"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WarangalDuke interface {
	feud.Duke
	CBalkonda婆罗军荼() balkonda.BalkondaCounty
	CKambampet剑婆摩波底() kambampet.KambampetCounty
	COrangallu乌兰伽楼() orangallu.OrangalluCounty
	CVemulavada吠牟罗婆荼() vemulavada.VemulavadaCounty
}

type 婆浪伽罗WarangalDuke struct {
	feud.BaseDuke
	婆罗军荼Balkonda    balkonda.BalkondaCounty
	剑婆摩波底Kambampet  kambampet.KambampetCounty
	乌兰伽楼Orangallu   orangallu.OrangalluCounty
	吠牟罗婆荼Vemulavada vemulavada.VemulavadaCounty
}

func (d *婆浪伽罗WarangalDuke) CBalkonda婆罗军荼() balkonda.BalkondaCounty {
	return d.婆罗军荼Balkonda
}

func (d *婆浪伽罗WarangalDuke) CKambampet剑婆摩波底() kambampet.KambampetCounty {
	return d.剑婆摩波底Kambampet
}

func (d *婆浪伽罗WarangalDuke) COrangallu乌兰伽楼() orangallu.OrangalluCounty {
	return d.乌兰伽楼Orangallu
}

func (d *婆浪伽罗WarangalDuke) CVemulavada吠牟罗婆荼() vemulavada.VemulavadaCounty {
	return d.吠牟罗婆荼Vemulavada
}

var DWarangal婆浪伽罗 WarangalDuke = &婆浪伽罗WarangalDuke{}

func init() {
	f := DWarangal婆浪伽罗.(*婆浪伽罗WarangalDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "warangal",
		TitleName: "婆浪伽罗",
		TitleCode: "d_warangal",
		Counties:  map[string]feud.County{},
	}

	f.婆罗军荼Balkonda = balkonda.CBalkonda婆罗军荼
	f.婆罗军荼Balkonda.SetParent(f)

	f.剑婆摩波底Kambampet = kambampet.CKambampet剑婆摩波底
	f.剑婆摩波底Kambampet.SetParent(f)

	f.乌兰伽楼Orangallu = orangallu.COrangallu乌兰伽楼
	f.乌兰伽楼Orangallu.SetParent(f)

	f.吠牟罗婆荼Vemulavada = vemulavada.CVemulavada吠牟罗婆荼
	f.吠牟罗婆荼Vemulavada.SetParent(f)

}
