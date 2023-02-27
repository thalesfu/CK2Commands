package sarasvati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富伽尔BukarBarony struct {
	feud.BaseBarony
}

var BBukar富伽尔 feud.Barony = &富伽尔BukarBarony{}

func init() {
    f := BBukar富伽尔.(*富伽尔BukarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bukar",
		TitleName: "富伽尔",
		TitleCode: "b_bukar",
	}
}
