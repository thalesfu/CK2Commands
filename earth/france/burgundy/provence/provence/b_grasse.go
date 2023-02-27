package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉斯GrasseBarony struct {
	feud.BaseBarony
}

var BGrasse格拉斯 feud.Barony = &格拉斯GrasseBarony{}

func init() {
    f := BGrasse格拉斯.(*格拉斯GrasseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grasse",
		TitleName: "格拉斯",
		TitleCode: "b_grasse",
	}
}
