package nikomedeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马拉吉纳MalaginaBarony struct {
	feud.BaseBarony
}

var BMalagina马拉吉纳 feud.Barony = &马拉吉纳MalaginaBarony{}

func init() {
    f := BMalagina马拉吉纳.(*马拉吉纳MalaginaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malagina",
		TitleName: "马拉吉纳",
		TitleCode: "b_malagina",
	}
}
