package lugo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥伦塞OrenseBarony struct {
	feud.BaseBarony
}

var BOrense奥伦塞 feud.Barony = &奥伦塞OrenseBarony{}

func init() {
    f := BOrense奥伦塞.(*奥伦塞OrenseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orense",
		TitleName: "奥伦塞",
		TitleCode: "b_orense",
	}
}
