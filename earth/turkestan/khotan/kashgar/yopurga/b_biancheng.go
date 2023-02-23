package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 遍城BianchengBarony struct {
	feud.BaseBarony
}

var BBiancheng遍城 feud.Barony = &遍城BianchengBarony{}

func init() {
	f := BBiancheng遍城.(*遍城BianchengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "biancheng",
		TitleName: "遍城",
		TitleCode: "b_biancheng",
	}
}
