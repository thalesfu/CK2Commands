package sasaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽波陀槃阇KapadvanjBarony struct {
	feud.BaseBarony
}

var BKapadvanj伽波陀槃阇 feud.Barony = &伽波陀槃阇KapadvanjBarony{}

func init() {
    f := BKapadvanj伽波陀槃阇.(*伽波陀槃阇KapadvanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kapadvanj",
		TitleName: "伽波陀槃阇",
		TitleCode: "b_kapadvanj",
	}
}
