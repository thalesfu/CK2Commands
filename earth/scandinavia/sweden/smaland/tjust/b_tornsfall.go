package tjust

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特恩斯法尔TornsfallBarony struct {
	feud.BaseBarony
}

var BTornsfall特恩斯法尔 feud.Barony = &特恩斯法尔TornsfallBarony{}

func init() {
    f := BTornsfall特恩斯法尔.(*特恩斯法尔TornsfallBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tornsfall",
		TitleName: "特恩斯法尔",
		TitleCode: "b_tornsfall",
	}
}
