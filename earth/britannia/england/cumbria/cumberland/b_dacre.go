package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戴克DacreBarony struct {
	feud.BaseBarony
}

var BDacre戴克 feud.Barony = &戴克DacreBarony{}

func init() {
	f := BDacre戴克.(*戴克DacreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dacre",
		TitleName: "戴克",
		TitleCode: "b_dacre",
	}
}
