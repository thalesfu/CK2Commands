package meath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉Tara_meathBarony struct {
	feud.BaseBarony
}

var BTara_meath塔拉 feud.Barony = &塔拉Tara_meathBarony{}

func init() {
    f := BTara_meath塔拉.(*塔拉Tara_meathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tara_meath",
		TitleName: "塔拉",
		TitleCode: "b_tara_meath",
	}
}
