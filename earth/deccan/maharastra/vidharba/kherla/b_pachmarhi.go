package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯杰默里PachmarhiBarony struct {
	feud.BaseBarony
}

var BPachmarhi伯杰默里 feud.Barony = &伯杰默里PachmarhiBarony{}

func init() {
    f := BPachmarhi伯杰默里.(*伯杰默里PachmarhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pachmarhi",
		TitleName: "伯杰默里",
		TitleCode: "b_pachmarhi",
	}
}
