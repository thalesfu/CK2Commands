package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 不动城AcalapuraBarony struct {
	feud.BaseBarony
}

var BAcalapura不动城 feud.Barony = &不动城AcalapuraBarony{}

func init() {
	f := BAcalapura不动城.(*不动城AcalapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "acalapura",
		TitleName: "不动城",
		TitleCode: "b_acalapura",
	}
}
