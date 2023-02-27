package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷巴基夫卡RybakivkaBarony struct {
	feud.BaseBarony
}

var BRybakivka雷巴基夫卡 feud.Barony = &雷巴基夫卡RybakivkaBarony{}

func init() {
    f := BRybakivka雷巴基夫卡.(*雷巴基夫卡RybakivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rybakivka",
		TitleName: "雷巴基夫卡",
		TitleCode: "b_rybakivka",
	}
}
