package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格伦赫根GronhogenBarony struct {
	feud.BaseBarony
}

var BGronhogen格伦赫根 feud.Barony = &格伦赫根GronhogenBarony{}

func init() {
    f := BGronhogen格伦赫根.(*格伦赫根GronhogenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gronhogen",
		TitleName: "格伦赫根",
		TitleCode: "b_gronhogen",
	}
}
