package east_karelia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶尔姆YelmBarony struct {
	feud.BaseBarony
}

var BYelm耶尔姆 feud.Barony = &耶尔姆YelmBarony{}

func init() {
    f := BYelm耶尔姆.(*耶尔姆YelmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yelm",
		TitleName: "耶尔姆",
		TitleCode: "b_yelm",
	}
}
