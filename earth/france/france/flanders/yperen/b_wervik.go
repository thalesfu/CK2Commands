package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦尔菲克WervikBarony struct {
	feud.BaseBarony
}

var BWervik韦尔菲克 feud.Barony = &韦尔菲克WervikBarony{}

func init() {
    f := BWervik韦尔菲克.(*韦尔菲克WervikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wervik",
		TitleName: "韦尔菲克",
		TitleCode: "b_wervik",
	}
}
