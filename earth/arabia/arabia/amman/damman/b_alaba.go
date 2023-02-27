package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾巴AlabaBarony struct {
	feud.BaseBarony
}

var BAlaba艾巴 feud.Barony = &艾巴AlabaBarony{}

func init() {
    f := BAlaba艾巴.(*艾巴AlabaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alaba",
		TitleName: "艾巴",
		TitleCode: "b_alaba",
	}
}
