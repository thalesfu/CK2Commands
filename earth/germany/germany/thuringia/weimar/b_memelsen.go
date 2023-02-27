package weimar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅梅尔森MemelsenBarony struct {
	feud.BaseBarony
}

var BMemelsen梅梅尔森 feud.Barony = &梅梅尔森MemelsenBarony{}

func init() {
    f := BMemelsen梅梅尔森.(*梅梅尔森MemelsenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "memelsen",
		TitleName: "梅梅尔森",
		TitleCode: "b_memelsen",
	}
}
