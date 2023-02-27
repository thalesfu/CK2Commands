package sinai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔AtturBarony struct {
	feud.BaseBarony
}

var BAttur图尔 feud.Barony = &图尔AtturBarony{}

func init() {
    f := BAttur图尔.(*图尔AtturBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "attur",
		TitleName: "图尔",
		TitleCode: "b_attur",
	}
}
