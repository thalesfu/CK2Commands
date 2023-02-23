package kiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迪昂吉纳雷DianguinareBarony struct {
	feud.BaseBarony
}

var BDianguinare迪昂吉纳雷 feud.Barony = &迪昂吉纳雷DianguinareBarony{}

func init() {
	f := BDianguinare迪昂吉纳雷.(*迪昂吉纳雷DianguinareBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dianguinare",
		TitleName: "迪昂吉纳雷",
		TitleCode: "b_dianguinare",
	}
}
