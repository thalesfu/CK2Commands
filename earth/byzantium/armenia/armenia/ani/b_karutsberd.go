package ani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡尔斯贝拉德KarutsberdBarony struct {
	feud.BaseBarony
}

var BKarutsberd卡尔斯贝拉德 feud.Barony = &卡尔斯贝拉德KarutsberdBarony{}

func init() {
    f := BKarutsberd卡尔斯贝拉德.(*卡尔斯贝拉德KarutsberdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karutsberd",
		TitleName: "卡尔斯贝拉德",
		TitleCode: "b_karutsberd",
	}
}
