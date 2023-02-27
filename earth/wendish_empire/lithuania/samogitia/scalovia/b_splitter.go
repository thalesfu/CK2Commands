package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯普利特SplitterBarony struct {
	feud.BaseBarony
}

var BSplitter斯普利特 feud.Barony = &斯普利特SplitterBarony{}

func init() {
    f := BSplitter斯普利特.(*斯普利特SplitterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "splitter",
		TitleName: "斯普利特",
		TitleCode: "b_splitter",
	}
}
