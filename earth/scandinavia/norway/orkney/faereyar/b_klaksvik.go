package faereyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克拉克斯维克KlaksvikBarony struct {
	feud.BaseBarony
}

var BKlaksvik克拉克斯维克 feud.Barony = &克拉克斯维克KlaksvikBarony{}

func init() {
    f := BKlaksvik克拉克斯维克.(*克拉克斯维克KlaksvikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "klaksvik",
		TitleName: "克拉克斯维克",
		TitleCode: "b_klaksvik",
	}
}
