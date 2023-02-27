package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿梅利亚AmeliaBarony struct {
	feud.BaseBarony
}

var BAmelia阿梅利亚 feud.Barony = &阿梅利亚AmeliaBarony{}

func init() {
    f := BAmelia阿梅利亚.(*阿梅利亚AmeliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amelia",
		TitleName: "阿梅利亚",
		TitleCode: "b_amelia",
	}
}
