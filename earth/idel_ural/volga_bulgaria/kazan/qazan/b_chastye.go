package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰斯特耶ChastyeBarony struct {
	feud.BaseBarony
}

var BChastye恰斯特耶 feud.Barony = &恰斯特耶ChastyeBarony{}

func init() {
    f := BChastye恰斯特耶.(*恰斯特耶ChastyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chastye",
		TitleName: "恰斯特耶",
		TitleCode: "b_chastye",
	}
}
