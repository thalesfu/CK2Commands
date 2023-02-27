package yeruslan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔西BorsiBarony struct {
	feud.BaseBarony
}

var BBorsi博尔西 feud.Barony = &博尔西BorsiBarony{}

func init() {
    f := BBorsi博尔西.(*博尔西BorsiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borsi",
		TitleName: "博尔西",
		TitleCode: "b_borsi",
	}
}
