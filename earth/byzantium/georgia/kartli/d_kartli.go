package kartli

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kartli/imeretia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kartli/kartli"
	"github.com/thalesfu/CK2Commands/earth/byzantium/georgia/kartli/mtskheta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KartliDuke interface {
	feud.Duke
	CImeretia伊梅列季亚() imeretia.ImeretiaCounty
	CKartli第比利斯() kartli.KartliCounty
	CMtskheta姆茨赫塔() mtskheta.MtskhetaCounty
}

type 卡尔特利KartliDuke struct {
	feud.BaseDuke
	伊梅列季亚Imeretia imeretia.ImeretiaCounty
	第比利斯Kartli    kartli.KartliCounty
	姆茨赫塔Mtskheta  mtskheta.MtskhetaCounty
}

func (d *卡尔特利KartliDuke) CImeretia伊梅列季亚() imeretia.ImeretiaCounty {
	return d.伊梅列季亚Imeretia
}

func (d *卡尔特利KartliDuke) CKartli第比利斯() kartli.KartliCounty {
	return d.第比利斯Kartli
}

func (d *卡尔特利KartliDuke) CMtskheta姆茨赫塔() mtskheta.MtskhetaCounty {
	return d.姆茨赫塔Mtskheta
}

var DKartli卡尔特利 KartliDuke = &卡尔特利KartliDuke{}

func init() {
	f := DKartli卡尔特利.(*卡尔特利KartliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kartli",
		TitleName: "卡尔特利",
		TitleCode: "d_kartli",
		Counties:  map[string]feud.County{},
	}

	f.伊梅列季亚Imeretia = imeretia.CImeretia伊梅列季亚
	f.伊梅列季亚Imeretia.SetParent(f)

	f.第比利斯Kartli = kartli.CKartli第比利斯
	f.第比利斯Kartli.SetParent(f)

	f.姆茨赫塔Mtskheta = mtskheta.CMtskheta姆茨赫塔
	f.姆茨赫塔Mtskheta.SetParent(f)

}
