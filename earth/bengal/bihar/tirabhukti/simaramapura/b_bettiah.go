package simaramapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠提阿BettiahBarony struct {
	feud.BaseBarony
}

var BBettiah吠提阿 feud.Barony = &吠提阿BettiahBarony{}

func init() {
    f := BBettiah吠提阿.(*吠提阿BettiahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bettiah",
		TitleName: "吠提阿",
		TitleCode: "b_bettiah",
	}
}
