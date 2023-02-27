package chortitza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚历山德罗夫斯克AlexandrowskBarony struct {
	feud.BaseBarony
}

var BAlexandrowsk亚历山德罗夫斯克 feud.Barony = &亚历山德罗夫斯克AlexandrowskBarony{}

func init() {
    f := BAlexandrowsk亚历山德罗夫斯克.(*亚历山德罗夫斯克AlexandrowskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alexandrowsk",
		TitleName: "亚历山德罗夫斯克",
		TitleCode: "b_alexandrowsk",
	}
}
