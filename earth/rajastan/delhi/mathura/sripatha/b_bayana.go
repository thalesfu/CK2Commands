package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆耶那BayanaBarony struct {
	feud.BaseBarony
}

var BBayana婆耶那 feud.Barony = &婆耶那BayanaBarony{}

func init() {
    f := BBayana婆耶那.(*婆耶那BayanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bayana",
		TitleName: "婆耶那",
		TitleCode: "b_bayana",
	}
}
