package benghazi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班加西BenghaziBarony struct {
	feud.BaseBarony
}

var BBenghazi班加西 feud.Barony = &班加西BenghaziBarony{}

func init() {
    f := BBenghazi班加西.(*班加西BenghaziBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "benghazi",
		TitleName: "班加西",
		TitleCode: "b_benghazi",
	}
}
