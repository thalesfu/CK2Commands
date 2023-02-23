package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔塔TataBarony struct {
	feud.BaseBarony
}

var BTata塔塔 feud.Barony = &塔塔TataBarony{}

func init() {
	f := BTata塔塔.(*塔塔TataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tata",
		TitleName: "塔塔",
		TitleCode: "b_tata",
	}
}
