package derbent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡拉贾HumrajBarony struct {
	feud.BaseBarony
}

var BHumraj胡拉贾 feud.Barony = &胡拉贾HumrajBarony{}

func init() {
    f := BHumraj胡拉贾.(*胡拉贾HumrajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "humraj",
		TitleName: "胡拉贾",
		TitleCode: "b_humraj",
	}
}
