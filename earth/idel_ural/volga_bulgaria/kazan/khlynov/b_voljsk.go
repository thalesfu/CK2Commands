package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃尔日斯克VoljskBarony struct {
	feud.BaseBarony
}

var BVoljsk沃尔日斯克 feud.Barony = &沃尔日斯克VoljskBarony{}

func init() {
    f := BVoljsk沃尔日斯克.(*沃尔日斯克VoljskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voljsk",
		TitleName: "沃尔日斯克",
		TitleCode: "b_voljsk",
	}
}
