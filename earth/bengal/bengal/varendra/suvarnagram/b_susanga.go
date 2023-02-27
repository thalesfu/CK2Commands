package suvarnagram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏僧伽SusangaBarony struct {
	feud.BaseBarony
}

var BSusanga苏僧伽 feud.Barony = &苏僧伽SusangaBarony{}

func init() {
    f := BSusanga苏僧伽.(*苏僧伽SusangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "susanga",
		TitleName: "苏僧伽",
		TitleCode: "b_susanga",
	}
}
