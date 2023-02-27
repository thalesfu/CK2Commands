package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾尔亚奈ArianaBarony struct {
	feud.BaseBarony
}

var BAriana艾尔亚奈 feud.Barony = &艾尔亚奈ArianaBarony{}

func init() {
    f := BAriana艾尔亚奈.(*艾尔亚奈ArianaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ariana",
		TitleName: "艾尔亚奈",
		TitleCode: "b_ariana",
	}
}
