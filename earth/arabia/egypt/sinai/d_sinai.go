package sinai

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/sinai/eilat"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/sinai/farama"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/sinai/sinai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SinaiDuke interface {
	feud.Duke
	CEilat埃拉特() eilat.EilatCounty
	CFarama珀卢西亚() farama.FaramaCounty
	CSinai西奈() sinai.SinaiCounty
}

type 西奈SinaiDuke struct {
	feud.BaseDuke
	埃拉特Eilat   eilat.EilatCounty
	珀卢西亚Farama farama.FaramaCounty
	西奈Sinai    sinai.SinaiCounty
}

func (d *西奈SinaiDuke) CEilat埃拉特() eilat.EilatCounty {
	return d.埃拉特Eilat
}

func (d *西奈SinaiDuke) CFarama珀卢西亚() farama.FaramaCounty {
	return d.珀卢西亚Farama
}

func (d *西奈SinaiDuke) CSinai西奈() sinai.SinaiCounty {
	return d.西奈Sinai
}

var DSinai西奈 SinaiDuke = &西奈SinaiDuke{}

func init() {
	f := DSinai西奈.(*西奈SinaiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sinai",
		TitleName: "西奈",
		TitleCode: "d_sinai",
		Counties:  map[string]feud.County{},
	}

	f.埃拉特Eilat = eilat.CEilat埃拉特
	f.埃拉特Eilat.SetParent(f)

	f.珀卢西亚Farama = farama.CFarama珀卢西亚
	f.珀卢西亚Farama.SetParent(f)

	f.西奈Sinai = sinai.CSinai西奈
	f.西奈Sinai.SetParent(f)

}
