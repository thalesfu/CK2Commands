package laodikeia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格尔德斯GordesBarony struct {
	feud.BaseBarony
}

var BGordes格尔德斯 feud.Barony = &格尔德斯GordesBarony{}

func init() {
    f := BGordes格尔德斯.(*格尔德斯GordesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gordes",
		TitleName: "格尔德斯",
		TitleCode: "b_gordes",
	}
}
