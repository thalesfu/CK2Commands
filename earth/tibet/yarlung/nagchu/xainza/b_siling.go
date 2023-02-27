package xainza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 色林SilingBarony struct {
	feud.BaseBarony
}

var BSiling色林 feud.Barony = &色林SilingBarony{}

func init() {
    f := BSiling色林.(*色林SilingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siling",
		TitleName: "色林",
		TitleCode: "b_siling",
	}
}
