package harer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德雷达瓦DiridawaBarony struct {
	feud.BaseBarony
}

var BDiridawa德雷达瓦 feud.Barony = &德雷达瓦DiridawaBarony{}

func init() {
	f := BDiridawa德雷达瓦.(*德雷达瓦DiridawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diridawa",
		TitleName: "德雷达瓦",
		TitleCode: "b_diridawa",
	}
}
