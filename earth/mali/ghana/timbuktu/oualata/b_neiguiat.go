package oualata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 内吉亚特NeiguiatBarony struct {
	feud.BaseBarony
}

var BNeiguiat内吉亚特 feud.Barony = &内吉亚特NeiguiatBarony{}

func init() {
	f := BNeiguiat内吉亚特.(*内吉亚特NeiguiatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "neiguiat",
		TitleName: "内吉亚特",
		TitleCode: "b_neiguiat",
	}
}
