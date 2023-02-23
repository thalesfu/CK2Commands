package kundina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 提那须吉阿TinsukiaBarony struct {
	feud.BaseBarony
}

var BTinsukia提那须吉阿 feud.Barony = &提那须吉阿TinsukiaBarony{}

func init() {
	f := BTinsukia提那须吉阿.(*提那须吉阿TinsukiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinsukia",
		TitleName: "提那须吉阿",
		TitleCode: "b_tinsukia",
	}
}
