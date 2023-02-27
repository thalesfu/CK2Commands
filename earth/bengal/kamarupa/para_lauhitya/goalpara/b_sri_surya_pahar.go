package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利苏利耶钵罗诃罗Sri_surya_paharBarony struct {
	feud.BaseBarony
}

var BSri_surya_pahar室利苏利耶钵罗诃罗 feud.Barony = &室利苏利耶钵罗诃罗Sri_surya_paharBarony{}

func init() {
    f := BSri_surya_pahar室利苏利耶钵罗诃罗.(*室利苏利耶钵罗诃罗Sri_surya_paharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sri_surya_pahar",
		TitleName: "室利苏利耶钵罗诃罗",
		TitleCode: "b_sri_surya_pahar",
	}
}
