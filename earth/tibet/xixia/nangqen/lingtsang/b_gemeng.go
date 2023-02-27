package lingtsang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格孟GemengBarony struct {
	feud.BaseBarony
}

var BGemeng格孟 feud.Barony = &格孟GemengBarony{}

func init() {
    f := BGemeng格孟.(*格孟GemengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gemeng",
		TitleName: "格孟",
		TitleCode: "b_gemeng",
	}
}
