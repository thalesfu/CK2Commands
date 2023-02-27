package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雍仲林YungdrunglingBarony struct {
	feud.BaseBarony
}

var BYungdrungling雍仲林 feud.Barony = &雍仲林YungdrunglingBarony{}

func init() {
    f := BYungdrungling雍仲林.(*雍仲林YungdrunglingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yungdrungling",
		TitleName: "雍仲林",
		TitleCode: "b_yungdrungling",
	}
}
