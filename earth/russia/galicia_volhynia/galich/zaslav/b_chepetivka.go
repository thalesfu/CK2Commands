package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍佩蒂夫卡ChepetivkaBarony struct {
	feud.BaseBarony
}

var BChepetivka舍佩蒂夫卡 feud.Barony = &舍佩蒂夫卡ChepetivkaBarony{}

func init() {
    f := BChepetivka舍佩蒂夫卡.(*舍佩蒂夫卡ChepetivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chepetivka",
		TitleName: "舍佩蒂夫卡",
		TitleCode: "b_chepetivka",
	}
}
