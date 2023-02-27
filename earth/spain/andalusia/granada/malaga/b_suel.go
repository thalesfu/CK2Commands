package malaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏埃尔SuelBarony struct {
	feud.BaseBarony
}

var BSuel苏埃尔 feud.Barony = &苏埃尔SuelBarony{}

func init() {
    f := BSuel苏埃尔.(*苏埃尔SuelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "suel",
		TitleName: "苏埃尔",
		TitleCode: "b_suel",
	}
}
