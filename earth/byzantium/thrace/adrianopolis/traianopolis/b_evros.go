package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃夫罗斯EvrosBarony struct {
	feud.BaseBarony
}

var BEvros埃夫罗斯 feud.Barony = &埃夫罗斯EvrosBarony{}

func init() {
    f := BEvros埃夫罗斯.(*埃夫罗斯EvrosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evros",
		TitleName: "埃夫罗斯",
		TitleCode: "b_evros",
	}
}
