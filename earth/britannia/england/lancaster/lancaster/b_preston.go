package lancaster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普雷斯顿PrestonBarony struct {
	feud.BaseBarony
}

var BPreston普雷斯顿 feud.Barony = &普雷斯顿PrestonBarony{}

func init() {
    f := BPreston普雷斯顿.(*普雷斯顿PrestonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "preston",
		TitleName: "普雷斯顿",
		TitleCode: "b_preston",
	}
}
