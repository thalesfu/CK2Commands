package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗拉丁根VlaardingenBarony struct {
	feud.BaseBarony
}

var BVlaardingen弗拉丁根 feud.Barony = &弗拉丁根VlaardingenBarony{}

func init() {
    f := BVlaardingen弗拉丁根.(*弗拉丁根VlaardingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vlaardingen",
		TitleName: "弗拉丁根",
		TitleCode: "b_vlaardingen",
	}
}
