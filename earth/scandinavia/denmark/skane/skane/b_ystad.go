package skane

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于斯塔德YstadBarony struct {
	feud.BaseBarony
}

var BYstad于斯塔德 feud.Barony = &于斯塔德YstadBarony{}

func init() {
	f := BYstad于斯塔德.(*于斯塔德YstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ystad",
		TitleName: "于斯塔德",
		TitleCode: "b_ystad",
	}
}
