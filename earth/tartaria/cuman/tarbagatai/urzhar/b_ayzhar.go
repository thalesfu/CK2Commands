package urzhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾扎尔AyzharBarony struct {
	feud.BaseBarony
}

var BAyzhar艾扎尔 feud.Barony = &艾扎尔AyzharBarony{}

func init() {
	f := BAyzhar艾扎尔.(*艾扎尔AyzharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ayzhar",
		TitleName: "艾扎尔",
		TitleCode: "b_ayzhar",
	}
}
