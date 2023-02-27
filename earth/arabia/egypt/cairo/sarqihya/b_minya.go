package sarqihya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明亚MinyaBarony struct {
	feud.BaseBarony
}

var BMinya明亚 feud.Barony = &明亚MinyaBarony{}

func init() {
    f := BMinya明亚.(*明亚MinyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "minya",
		TitleName: "明亚",
		TitleCode: "b_minya",
	}
}
