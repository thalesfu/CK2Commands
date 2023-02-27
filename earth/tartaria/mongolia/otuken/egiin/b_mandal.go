package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼达勒MandalBarony struct {
	feud.BaseBarony
}

var BMandal曼达勒 feud.Barony = &曼达勒MandalBarony{}

func init() {
    f := BMandal曼达勒.(*曼达勒MandalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandal",
		TitleName: "曼达勒",
		TitleCode: "b_mandal",
	}
}
