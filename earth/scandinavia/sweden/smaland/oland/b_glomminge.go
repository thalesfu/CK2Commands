package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格勒明厄GlommingeBarony struct {
	feud.BaseBarony
}

var BGlomminge格勒明厄 feud.Barony = &格勒明厄GlommingeBarony{}

func init() {
    f := BGlomminge格勒明厄.(*格勒明厄GlommingeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glomminge",
		TitleName: "格勒明厄",
		TitleCode: "b_glomminge",
	}
}
