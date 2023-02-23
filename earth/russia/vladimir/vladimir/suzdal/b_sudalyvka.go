package suzdal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏达雷夫卡SudalyvkaBarony struct {
	feud.BaseBarony
}

var BSudalyvka苏达雷夫卡 feud.Barony = &苏达雷夫卡SudalyvkaBarony{}

func init() {
	f := BSudalyvka苏达雷夫卡.(*苏达雷夫卡SudalyvkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sudalyvka",
		TitleName: "苏达雷夫卡",
		TitleCode: "b_sudalyvka",
	}
}
