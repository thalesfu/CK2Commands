package qamdo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日通RetongBarony struct {
	feud.BaseBarony
}

var BRetong日通 feud.Barony = &日通RetongBarony{}

func init() {
    f := BRetong日通.(*日通RetongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "retong",
		TitleName: "日通",
		TitleCode: "b_retong",
	}
}
