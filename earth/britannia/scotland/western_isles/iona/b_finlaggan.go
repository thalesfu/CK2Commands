package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 芬拉甘FinlagganBarony struct {
	feud.BaseBarony
}

var BFinlaggan芬拉甘 feud.Barony = &芬拉甘FinlagganBarony{}

func init() {
    f := BFinlaggan芬拉甘.(*芬拉甘FinlagganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "finlaggan",
		TitleName: "芬拉甘",
		TitleCode: "b_finlaggan",
	}
}
