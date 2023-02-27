package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博布鲁伊斯克BabruyskBarony struct {
	feud.BaseBarony
}

var BBabruysk博布鲁伊斯克 feud.Barony = &博布鲁伊斯克BabruyskBarony{}

func init() {
    f := BBabruysk博布鲁伊斯克.(*博布鲁伊斯克BabruyskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "babruysk",
		TitleName: "博布鲁伊斯克",
		TitleCode: "b_babruysk",
	}
}
