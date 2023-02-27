package split

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯普利特SplitBarony struct {
	feud.BaseBarony
}

var BSplit斯普利特 feud.Barony = &斯普利特SplitBarony{}

func init() {
    f := BSplit斯普利特.(*斯普利特SplitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "split",
		TitleName: "斯普利特",
		TitleCode: "b_split",
	}
}
