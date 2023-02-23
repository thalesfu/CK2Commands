package garhwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 北迦尸UttarkashiBarony struct {
	feud.BaseBarony
}

var BUttarkashi北迦尸 feud.Barony = &北迦尸UttarkashiBarony{}

func init() {
	f := BUttarkashi北迦尸.(*北迦尸UttarkashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uttarkashi",
		TitleName: "北迦尸",
		TitleCode: "b_uttarkashi",
	}
}
