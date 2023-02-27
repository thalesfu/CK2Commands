package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔BureBarony struct {
	feud.BaseBarony
}

var BBure布尔 feud.Barony = &布尔BureBarony{}

func init() {
    f := BBure布尔.(*布尔BureBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bure",
		TitleName: "布尔",
		TitleCode: "b_bure",
	}
}
