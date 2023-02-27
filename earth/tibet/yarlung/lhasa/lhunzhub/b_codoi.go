package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 春堆CodoiBarony struct {
	feud.BaseBarony
}

var BCodoi春堆 feud.Barony = &春堆CodoiBarony{}

func init() {
    f := BCodoi春堆.(*春堆CodoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "codoi",
		TitleName: "春堆",
		TitleCode: "b_codoi",
	}
}
