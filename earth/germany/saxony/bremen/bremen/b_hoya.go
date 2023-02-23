package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍亚HoyaBarony struct {
	feud.BaseBarony
}

var BHoya霍亚 feud.Barony = &霍亚HoyaBarony{}

func init() {
	f := BHoya霍亚.(*霍亚HoyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hoya",
		TitleName: "霍亚",
		TitleCode: "b_hoya",
	}
}
