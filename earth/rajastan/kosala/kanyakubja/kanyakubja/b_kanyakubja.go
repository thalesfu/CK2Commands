package kanyakubja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲女城KanyakubjaBarony struct {
	feud.BaseBarony
}

var BKanyakubja曲女城 feud.Barony = &曲女城KanyakubjaBarony{}

func init() {
    f := BKanyakubja曲女城.(*曲女城KanyakubjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanyakubja",
		TitleName: "曲女城",
		TitleCode: "b_kanyakubja",
	}
}
