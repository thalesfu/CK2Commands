package pokhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毕哩体微那罗延PrithbinarayanBarony struct {
	feud.BaseBarony
}

var BPrithbinarayan毕哩体微那罗延 feud.Barony = &毕哩体微那罗延PrithbinarayanBarony{}

func init() {
    f := BPrithbinarayan毕哩体微那罗延.(*毕哩体微那罗延PrithbinarayanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prithbinarayan",
		TitleName: "毕哩体微那罗延",
		TitleCode: "b_prithbinarayan",
	}
}
