package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯撒利亚SyrcaesareaBarony struct {
	feud.BaseBarony
}

var BSyrcaesarea凯撒利亚 feud.Barony = &凯撒利亚SyrcaesareaBarony{}

func init() {
	f := BSyrcaesarea凯撒利亚.(*凯撒利亚SyrcaesareaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "syrcaesarea",
		TitleName: "凯撒利亚",
		TitleCode: "b_syrcaesarea",
	}
}
