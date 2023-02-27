package radha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瞿波部摩GopbhumBarony struct {
	feud.BaseBarony
}

var BGopbhum瞿波部摩 feud.Barony = &瞿波部摩GopbhumBarony{}

func init() {
    f := BGopbhum瞿波部摩.(*瞿波部摩GopbhumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gopbhum",
		TitleName: "瞿波部摩",
		TitleCode: "b_gopbhum",
	}
}
