package gafsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托泽尔TozeurBarony struct {
	feud.BaseBarony
}

var BTozeur托泽尔 feud.Barony = &托泽尔TozeurBarony{}

func init() {
    f := BTozeur托泽尔.(*托泽尔TozeurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tozeur",
		TitleName: "托泽尔",
		TitleCode: "b_tozeur",
	}
}
