package opava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格武布奇采HlubciceBarony struct {
	feud.BaseBarony
}

var BHlubcice格武布奇采 feud.Barony = &格武布奇采HlubciceBarony{}

func init() {
    f := BHlubcice格武布奇采.(*格武布奇采HlubciceBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hlubcice",
		TitleName: "格武布奇采",
		TitleCode: "b_hlubcice",
	}
}
