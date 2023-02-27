package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格温GerwinBarony struct {
	feud.BaseBarony
}

var BGerwin格温 feud.Barony = &格温GerwinBarony{}

func init() {
    f := BGerwin格温.(*格温GerwinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerwin",
		TitleName: "格温",
		TitleCode: "b_gerwin",
	}
}
