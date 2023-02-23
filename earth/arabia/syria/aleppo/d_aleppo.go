package aleppo

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/aleppo/aleppo"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/aleppo/asas"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/aleppo/hama"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/aleppo/homs"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AleppoDuke interface {
	feud.Duke
	CAleppo阿勒颇() aleppo.AleppoCounty
	CAsas阿萨斯() asas.AsasCounty
	CHama哈马() hama.HamaCounty
	CHoms霍姆斯() homs.HomsCounty
}

type 阿勒颇AleppoDuke struct {
	feud.BaseDuke
	阿勒颇Aleppo aleppo.AleppoCounty
	阿萨斯Asas   asas.AsasCounty
	哈马Hama    hama.HamaCounty
	霍姆斯Homs   homs.HomsCounty
}

func (d *阿勒颇AleppoDuke) CAleppo阿勒颇() aleppo.AleppoCounty {
	return d.阿勒颇Aleppo
}

func (d *阿勒颇AleppoDuke) CAsas阿萨斯() asas.AsasCounty {
	return d.阿萨斯Asas
}

func (d *阿勒颇AleppoDuke) CHama哈马() hama.HamaCounty {
	return d.哈马Hama
}

func (d *阿勒颇AleppoDuke) CHoms霍姆斯() homs.HomsCounty {
	return d.霍姆斯Homs
}

var DAleppo阿勒颇 AleppoDuke = &阿勒颇AleppoDuke{}

func init() {
	f := DAleppo阿勒颇.(*阿勒颇AleppoDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aleppo",
		TitleName: "阿勒颇",
		TitleCode: "d_aleppo",
		Counties:  map[string]feud.County{},
	}

	f.阿勒颇Aleppo = aleppo.CAleppo阿勒颇
	f.阿勒颇Aleppo.SetParent(f)

	f.阿萨斯Asas = asas.CAsas阿萨斯
	f.阿萨斯Asas.SetParent(f)

	f.哈马Hama = hama.CHama哈马
	f.哈马Hama.SetParent(f)

	f.霍姆斯Homs = homs.CHoms霍姆斯
	f.霍姆斯Homs.SetParent(f)

}
