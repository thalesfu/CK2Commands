package tus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔万ShervanBarony struct {
	feud.BaseBarony
}

var BShervan希尔万 feud.Barony = &希尔万ShervanBarony{}

func init() {
    f := BShervan希尔万.(*希尔万ShervanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shervan",
		TitleName: "希尔万",
		TitleCode: "b_shervan",
	}
}
