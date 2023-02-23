package antioch

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/antioch/alexandretta"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/antioch/antiocheia"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/antioch/archa"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AntiochDuke interface {
	feud.Duke
	CAlexandretta亚历山大勒塔() alexandretta.AlexandrettaCounty
	CAntiocheia安条克() antiocheia.AntiocheiaCounty
	CArcha阿尔卡() archa.ArchaCounty
}

type 安条克AntiochDuke struct {
	feud.BaseDuke
	亚历山大勒塔Alexandretta alexandretta.AlexandrettaCounty
	安条克Antiocheia      antiocheia.AntiocheiaCounty
	阿尔卡Archa           archa.ArchaCounty
}

func (d *安条克AntiochDuke) CAlexandretta亚历山大勒塔() alexandretta.AlexandrettaCounty {
	return d.亚历山大勒塔Alexandretta
}

func (d *安条克AntiochDuke) CAntiocheia安条克() antiocheia.AntiocheiaCounty {
	return d.安条克Antiocheia
}

func (d *安条克AntiochDuke) CArcha阿尔卡() archa.ArchaCounty {
	return d.阿尔卡Archa
}

var DAntioch安条克 AntiochDuke = &安条克AntiochDuke{}

func init() {
	f := DAntioch安条克.(*安条克AntiochDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "antioch",
		TitleName: "安条克",
		TitleCode: "d_antioch",
		Counties:  map[string]feud.County{},
	}

	f.亚历山大勒塔Alexandretta = alexandretta.CAlexandretta亚历山大勒塔
	f.亚历山大勒塔Alexandretta.SetParent(f)

	f.安条克Antiocheia = antiocheia.CAntiocheia安条克
	f.安条克Antiocheia.SetParent(f)

	f.阿尔卡Archa = archa.CArcha阿尔卡
	f.阿尔卡Archa.SetParent(f)

}
