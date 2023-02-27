package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格林CorlinBarony struct {
	feud.BaseBarony
}

var BCorlin格林 feud.Barony = &格林CorlinBarony{}

func init() {
    f := BCorlin格林.(*格林CorlinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corlin",
		TitleName: "格林",
		TitleCode: "b_corlin",
	}
}
