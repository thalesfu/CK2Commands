package marmaros

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰什克TecsoBarony struct {
	feud.BaseBarony
}

var BTecso泰什克 feud.Barony = &泰什克TecsoBarony{}

func init() {
	f := BTecso泰什克.(*泰什克TecsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tecso",
		TitleName: "泰什克",
		TitleCode: "b_tecso",
	}
}
