package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍伯恩SherborneBarony struct {
	feud.BaseBarony
}

var BSherborne舍伯恩 feud.Barony = &舍伯恩SherborneBarony{}

func init() {
    f := BSherborne舍伯恩.(*舍伯恩SherborneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sherborne",
		TitleName: "舍伯恩",
		TitleCode: "b_sherborne",
	}
}
