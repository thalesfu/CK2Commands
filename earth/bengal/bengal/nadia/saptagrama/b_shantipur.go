package saptagrama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扇底补罗ShantipurBarony struct {
	feud.BaseBarony
}

var BShantipur扇底补罗 feud.Barony = &扇底补罗ShantipurBarony{}

func init() {
    f := BShantipur扇底补罗.(*扇底补罗ShantipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shantipur",
		TitleName: "扇底补罗",
		TitleCode: "b_shantipur",
	}
}
