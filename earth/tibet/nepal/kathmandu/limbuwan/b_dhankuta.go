package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 驮那矩吒DhankutaBarony struct {
	feud.BaseBarony
}

var BDhankuta驮那矩吒 feud.Barony = &驮那矩吒DhankutaBarony{}

func init() {
	f := BDhankuta驮那矩吒.(*驮那矩吒DhankutaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhankuta",
		TitleName: "驮那矩吒",
		TitleCode: "b_dhankuta",
	}
}
