package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博里霍尔姆BorgholmBarony struct {
	feud.BaseBarony
}

var BBorgholm博里霍尔姆 feud.Barony = &博里霍尔姆BorgholmBarony{}

func init() {
    f := BBorgholm博里霍尔姆.(*博里霍尔姆BorgholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borgholm",
		TitleName: "博里霍尔姆",
		TitleCode: "b_borgholm",
	}
}
