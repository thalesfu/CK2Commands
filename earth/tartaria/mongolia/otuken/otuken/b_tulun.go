package otuken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图伦TulunBarony struct {
	feud.BaseBarony
}

var BTulun图伦 feud.Barony = &图伦TulunBarony{}

func init() {
    f := BTulun图伦.(*图伦TulunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulun",
		TitleName: "图伦",
		TitleCode: "b_tulun",
	}
}
