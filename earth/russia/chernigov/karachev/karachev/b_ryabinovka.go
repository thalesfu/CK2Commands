package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里亚比诺夫卡RyabinovkaBarony struct {
	feud.BaseBarony
}

var BRyabinovka里亚比诺夫卡 feud.Barony = &里亚比诺夫卡RyabinovkaBarony{}

func init() {
    f := BRyabinovka里亚比诺夫卡.(*里亚比诺夫卡RyabinovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ryabinovka",
		TitleName: "里亚比诺夫卡",
		TitleCode: "b_ryabinovka",
	}
}
