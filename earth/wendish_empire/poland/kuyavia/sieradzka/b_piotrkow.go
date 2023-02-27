package sieradzka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 彼得库夫PiotrkowBarony struct {
	feud.BaseBarony
}

var BPiotrkow彼得库夫 feud.Barony = &彼得库夫PiotrkowBarony{}

func init() {
    f := BPiotrkow彼得库夫.(*彼得库夫PiotrkowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "piotrkow",
		TitleName: "彼得库夫",
		TitleCode: "b_piotrkow",
	}
}
