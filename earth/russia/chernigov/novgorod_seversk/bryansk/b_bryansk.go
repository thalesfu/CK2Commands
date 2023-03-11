package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布良斯克BryanskBarony struct {
	feud.BaseBarony
}

var BBryansk布良斯克 feud.Barony = &布良斯克BryanskBarony{}

func init() {
    f := BBryansk布良斯克.(*布良斯克BryanskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bryansk",
		TitleName: "布良斯克",
		TitleCode: "b_bryansk",
	}
}
