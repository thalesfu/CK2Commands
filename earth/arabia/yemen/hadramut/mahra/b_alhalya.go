package mahra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒亚AlhalyaBarony struct {
	feud.BaseBarony
}

var BAlhalya哈勒亚 feud.Barony = &哈勒亚AlhalyaBarony{}

func init() {
	f := BAlhalya哈勒亚.(*哈勒亚AlhalyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alhalya",
		TitleName: "哈勒亚",
		TitleCode: "b_alhalya",
	}
}
