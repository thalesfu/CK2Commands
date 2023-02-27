package aurillac

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧里亚克AurillacBarony struct {
	feud.BaseBarony
}

var BAurillac欧里亚克 feud.Barony = &欧里亚克AurillacBarony{}

func init() {
    f := BAurillac欧里亚克.(*欧里亚克AurillacBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aurillac",
		TitleName: "欧里亚克",
		TitleCode: "b_aurillac",
	}
}
