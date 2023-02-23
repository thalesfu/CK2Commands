package jerid

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/jerid/gabes"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/jerid/gafsa"
	"github.com/thalesfu/CK2Commands/earth/maghreb/africa/jerid/suf"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JeridDuke interface {
	feud.Duke
	CGabes加贝斯() gabes.GabesCounty
	CGafsa加夫萨() gafsa.GafsaCounty
	CSuf苏夫() suf.SufCounty
}

type 杰里德JeridDuke struct {
	feud.BaseDuke
	加贝斯Gabes gabes.GabesCounty
	加夫萨Gafsa gafsa.GafsaCounty
	苏夫Suf    suf.SufCounty
}

func (d *杰里德JeridDuke) CGabes加贝斯() gabes.GabesCounty {
	return d.加贝斯Gabes
}

func (d *杰里德JeridDuke) CGafsa加夫萨() gafsa.GafsaCounty {
	return d.加夫萨Gafsa
}

func (d *杰里德JeridDuke) CSuf苏夫() suf.SufCounty {
	return d.苏夫Suf
}

var DJerid杰里德 JeridDuke = &杰里德JeridDuke{}

func init() {
	f := DJerid杰里德.(*杰里德JeridDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "jerid",
		TitleName: "杰里德",
		TitleCode: "d_jerid",
		Counties:  map[string]feud.County{},
	}

	f.加贝斯Gabes = gabes.CGabes加贝斯
	f.加贝斯Gabes.SetParent(f)

	f.加夫萨Gafsa = gafsa.CGafsa加夫萨
	f.加夫萨Gafsa.SetParent(f)

	f.苏夫Suf = suf.CSuf苏夫
	f.苏夫Suf.SetParent(f)

}
