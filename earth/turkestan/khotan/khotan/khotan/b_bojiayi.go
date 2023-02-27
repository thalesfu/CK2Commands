package khotan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勃伽夷BojiayiBarony struct {
	feud.BaseBarony
}

var BBojiayi勃伽夷 feud.Barony = &勃伽夷BojiayiBarony{}

func init() {
    f := BBojiayi勃伽夷.(*勃伽夷BojiayiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bojiayi",
		TitleName: "勃伽夷",
		TitleCode: "b_bojiayi",
	}
}
