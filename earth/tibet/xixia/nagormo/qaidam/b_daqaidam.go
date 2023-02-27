package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大柴旦DaqaidamBarony struct {
	feud.BaseBarony
}

var BDaqaidam大柴旦 feud.Barony = &大柴旦DaqaidamBarony{}

func init() {
    f := BDaqaidam大柴旦.(*大柴旦DaqaidamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daqaidam",
		TitleName: "大柴旦",
		TitleCode: "b_daqaidam",
	}
}
