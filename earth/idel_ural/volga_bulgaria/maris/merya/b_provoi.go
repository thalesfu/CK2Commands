package merya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗沃伊ProvoiBarony struct {
	feud.BaseBarony
}

var BProvoi普罗沃伊 feud.Barony = &普罗沃伊ProvoiBarony{}

func init() {
    f := BProvoi普罗沃伊.(*普罗沃伊ProvoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "provoi",
		TitleName: "普罗沃伊",
		TitleCode: "b_provoi",
	}
}
