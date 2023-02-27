package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫卧儿MughalBarony struct {
	feud.BaseBarony
}

var BMughal莫卧儿 feud.Barony = &莫卧儿MughalBarony{}

func init() {
    f := BMughal莫卧儿.(*莫卧儿MughalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mughal",
		TitleName: "莫卧儿",
		TitleCode: "b_mughal",
	}
}
