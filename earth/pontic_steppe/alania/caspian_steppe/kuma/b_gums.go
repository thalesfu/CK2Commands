package kuma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 居姆谢GumsBarony struct {
	feud.BaseBarony
}

var BGums居姆谢 feud.Barony = &居姆谢GumsBarony{}

func init() {
    f := BGums居姆谢.(*居姆谢GumsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gums",
		TitleName: "居姆谢",
		TitleCode: "b_gums",
	}
}
