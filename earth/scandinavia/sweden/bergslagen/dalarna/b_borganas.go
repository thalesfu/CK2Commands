package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔加内斯BorganasBarony struct {
	feud.BaseBarony
}

var BBorganas博尔加内斯 feud.Barony = &博尔加内斯BorganasBarony{}

func init() {
    f := BBorganas博尔加内斯.(*博尔加内斯BorganasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borganas",
		TitleName: "博尔加内斯",
		TitleCode: "b_borganas",
	}
}
