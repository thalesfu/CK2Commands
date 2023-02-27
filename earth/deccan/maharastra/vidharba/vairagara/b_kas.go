package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽斯KasBarony struct {
	feud.BaseBarony
}

var BKas伽斯 feud.Barony = &伽斯KasBarony{}

func init() {
    f := BKas伽斯.(*伽斯KasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kas",
		TitleName: "伽斯",
		TitleCode: "b_kas",
	}
}
