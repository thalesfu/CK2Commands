package cumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彭里斯PenrithBarony struct {
	feud.BaseBarony
}

var BPenrith彭里斯 feud.Barony = &彭里斯PenrithBarony{}

func init() {
	f := BPenrith彭里斯.(*彭里斯PenrithBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penrith",
		TitleName: "彭里斯",
		TitleCode: "b_penrith",
	}
}
