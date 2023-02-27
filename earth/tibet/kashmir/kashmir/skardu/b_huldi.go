package skardu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尔迪HuldiBarony struct {
	feud.BaseBarony
}

var BHuldi哈尔迪 feud.Barony = &哈尔迪HuldiBarony{}

func init() {
    f := BHuldi哈尔迪.(*哈尔迪HuldiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huldi",
		TitleName: "哈尔迪",
		TitleCode: "b_huldi",
	}
}
