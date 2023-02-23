package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格迪耶勒GdyelBarony struct {
	feud.BaseBarony
}

var BGdyel格迪耶勒 feud.Barony = &格迪耶勒GdyelBarony{}

func init() {
	f := BGdyel格迪耶勒.(*格迪耶勒GdyelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gdyel",
		TitleName: "格迪耶勒",
		TitleCode: "b_gdyel",
	}
}
