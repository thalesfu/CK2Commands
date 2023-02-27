package eastern_sayan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叶尔热伊ErzheyBarony struct {
	feud.BaseBarony
}

var BErzhey叶尔热伊 feud.Barony = &叶尔热伊ErzheyBarony{}

func init() {
    f := BErzhey叶尔热伊.(*叶尔热伊ErzheyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "erzhey",
		TitleName: "叶尔热伊",
		TitleCode: "b_erzhey",
	}
}
