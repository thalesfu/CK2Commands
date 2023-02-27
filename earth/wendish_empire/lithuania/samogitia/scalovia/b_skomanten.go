package scalovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯科曼滕SkomantenBarony struct {
	feud.BaseBarony
}

var BSkomanten斯科曼滕 feud.Barony = &斯科曼滕SkomantenBarony{}

func init() {
    f := BSkomanten斯科曼滕.(*斯科曼滕SkomantenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skomanten",
		TitleName: "斯科曼滕",
		TitleCode: "b_skomanten",
	}
}
