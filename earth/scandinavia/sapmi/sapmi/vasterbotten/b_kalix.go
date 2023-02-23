package vasterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利克斯KalixBarony struct {
	feud.BaseBarony
}

var BKalix卡利克斯 feud.Barony = &卡利克斯KalixBarony{}

func init() {
	f := BKalix卡利克斯.(*卡利克斯KalixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalix",
		TitleName: "卡利克斯",
		TitleCode: "b_kalix",
	}
}
