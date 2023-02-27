package selenge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔戈伊BorgoyBarony struct {
	feud.BaseBarony
}

var BBorgoy博尔戈伊 feud.Barony = &博尔戈伊BorgoyBarony{}

func init() {
    f := BBorgoy博尔戈伊.(*博尔戈伊BorgoyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgoy",
		TitleName: "博尔戈伊",
		TitleCode: "b_borgoy",
	}
}
