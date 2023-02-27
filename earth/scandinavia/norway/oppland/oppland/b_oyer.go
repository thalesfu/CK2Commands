package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 厄耶OyerBarony struct {
	feud.BaseBarony
}

var BOyer厄耶 feud.Barony = &厄耶OyerBarony{}

func init() {
    f := BOyer厄耶.(*厄耶OyerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oyer",
		TitleName: "厄耶",
		TitleCode: "b_oyer",
	}
}
