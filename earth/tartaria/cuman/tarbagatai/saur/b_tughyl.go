package saur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吐格勒TughylBarony struct {
	feud.BaseBarony
}

var BTughyl吐格勒 feud.Barony = &吐格勒TughylBarony{}

func init() {
    f := BTughyl吐格勒.(*吐格勒TughylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tughyl",
		TitleName: "吐格勒",
		TitleCode: "b_tughyl",
	}
}
