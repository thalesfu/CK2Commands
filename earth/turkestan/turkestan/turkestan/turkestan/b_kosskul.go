package turkestan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯_库尔KosskulBarony struct {
	feud.BaseBarony
}

var BKosskul科斯_库尔 feud.Barony = &科斯_库尔KosskulBarony{}

func init() {
    f := BKosskul科斯_库尔.(*科斯_库尔KosskulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosskul",
		TitleName: "科斯_库尔",
		TitleCode: "b_kosskul",
	}
}
