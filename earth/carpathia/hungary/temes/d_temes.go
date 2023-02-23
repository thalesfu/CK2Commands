package temes

import (
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/temes/bacs"
	"github.com/thalesfu/CK2Commands/earth/carpathia/hungary/temes/temes"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TemesDuke interface {
	feud.Duke
	CBacs巴奇() bacs.BacsCounty
	CTemes泰梅什() temes.TemesCounty
}

type 泰梅什TemesDuke struct {
	feud.BaseDuke
	巴奇Bacs   bacs.BacsCounty
	泰梅什Temes temes.TemesCounty
}

func (d *泰梅什TemesDuke) CBacs巴奇() bacs.BacsCounty {
	return d.巴奇Bacs
}

func (d *泰梅什TemesDuke) CTemes泰梅什() temes.TemesCounty {
	return d.泰梅什Temes
}

var DTemes泰梅什 TemesDuke = &泰梅什TemesDuke{}

func init() {
	f := DTemes泰梅什.(*泰梅什TemesDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "temes",
		TitleName: "泰梅什",
		TitleCode: "d_temes",
		Counties:  map[string]feud.County{},
	}

	f.巴奇Bacs = bacs.CBacs巴奇
	f.巴奇Bacs.SetParent(f)

	f.泰梅什Temes = temes.CTemes泰梅什
	f.泰梅什Temes.SetParent(f)

}
