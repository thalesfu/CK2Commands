package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达木布克DambukBarony struct {
	feud.BaseBarony
}

var BDambuk达木布克 feud.Barony = &达木布克DambukBarony{}

func init() {
	f := BDambuk达木布克.(*达木布克DambukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dambuk",
		TitleName: "达木布克",
		TitleCode: "b_dambuk",
	}
}
