package lombardia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙扎MonzaBarony struct {
	feud.BaseBarony
}

var BMonza蒙扎 feud.Barony = &蒙扎MonzaBarony{}

func init() {
    f := BMonza蒙扎.(*蒙扎MonzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monza",
		TitleName: "蒙扎",
		TitleCode: "b_monza",
	}
}
