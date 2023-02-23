package gloucester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格洛斯特GloucesterBarony struct {
	feud.BaseBarony
}

var BGloucester格洛斯特 feud.Barony = &格洛斯特GloucesterBarony{}

func init() {
	f := BGloucester格洛斯特.(*格洛斯特GloucesterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gloucester",
		TitleName: "格洛斯特",
		TitleCode: "b_gloucester",
	}
}
