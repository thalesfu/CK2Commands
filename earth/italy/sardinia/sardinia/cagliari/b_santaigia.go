package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣伊吉亚SantaigiaBarony struct {
	feud.BaseBarony
}

var BSantaigia圣伊吉亚 feud.Barony = &圣伊吉亚SantaigiaBarony{}

func init() {
    f := BSantaigia圣伊吉亚.(*圣伊吉亚SantaigiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santaigia",
		TitleName: "圣伊吉亚",
		TitleCode: "b_santaigia",
	}
}
