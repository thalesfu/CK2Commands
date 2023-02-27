package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提迦波汉提DigapahandiBarony struct {
	feud.BaseBarony
}

var BDigapahandi提迦波汉提 feud.Barony = &提迦波汉提DigapahandiBarony{}

func init() {
    f := BDigapahandi提迦波汉提.(*提迦波汉提DigapahandiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "digapahandi",
		TitleName: "提迦波汉提",
		TitleCode: "b_digapahandi",
	}
}
