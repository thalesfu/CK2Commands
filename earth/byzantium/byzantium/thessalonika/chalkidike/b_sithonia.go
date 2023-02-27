package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡索尼亚SithoniaBarony struct {
	feud.BaseBarony
}

var BSithonia锡索尼亚 feud.Barony = &锡索尼亚SithoniaBarony{}

func init() {
    f := BSithonia锡索尼亚.(*锡索尼亚SithoniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sithonia",
		TitleName: "锡索尼亚",
		TitleCode: "b_sithonia",
	}
}
