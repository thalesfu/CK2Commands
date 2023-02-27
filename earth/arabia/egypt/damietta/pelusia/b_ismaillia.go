package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯梅利亚IsmailliaBarony struct {
	feud.BaseBarony
}

var BIsmaillia伊斯梅利亚 feud.Barony = &伊斯梅利亚IsmailliaBarony{}

func init() {
    f := BIsmaillia伊斯梅利亚.(*伊斯梅利亚IsmailliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ismaillia",
		TitleName: "伊斯梅利亚",
		TitleCode: "b_ismaillia",
	}
}
