package lyzha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎哈尔万ZakharvanBarony struct {
	feud.BaseBarony
}

var BZakharvan扎哈尔万 feud.Barony = &扎哈尔万ZakharvanBarony{}

func init() {
    f := BZakharvan扎哈尔万.(*扎哈尔万ZakharvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zakharvan",
		TitleName: "扎哈尔万",
		TitleCode: "b_zakharvan",
	}
}
