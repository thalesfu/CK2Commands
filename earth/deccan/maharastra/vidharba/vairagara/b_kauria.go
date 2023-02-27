package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乔梨KauriaBarony struct {
	feud.BaseBarony
}

var BKauria乔梨 feud.Barony = &乔梨KauriaBarony{}

func init() {
    f := BKauria乔梨.(*乔梨KauriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kauria",
		TitleName: "乔梨",
		TitleCode: "b_kauria",
	}
}
