package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克托别AqtobeBarony struct {
	feud.BaseBarony
}

var BAqtobe阿克托别 feud.Barony = &阿克托别AqtobeBarony{}

func init() {
    f := BAqtobe阿克托别.(*阿克托别AqtobeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqtobe",
		TitleName: "阿克托别",
		TitleCode: "b_aqtobe",
	}
}
