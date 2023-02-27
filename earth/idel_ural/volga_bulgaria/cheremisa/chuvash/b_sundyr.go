package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孙德里SundyrBarony struct {
	feud.BaseBarony
}

var BSundyr孙德里 feud.Barony = &孙德里SundyrBarony{}

func init() {
    f := BSundyr孙德里.(*孙德里SundyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sundyr",
		TitleName: "孙德里",
		TitleCode: "b_sundyr",
	}
}
