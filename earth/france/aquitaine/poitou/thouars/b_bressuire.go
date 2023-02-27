package thouars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布雷叙尔BressuireBarony struct {
	feud.BaseBarony
}

var BBressuire布雷叙尔 feud.Barony = &布雷叙尔BressuireBarony{}

func init() {
    f := BBressuire布雷叙尔.(*布雷叙尔BressuireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bressuire",
		TitleName: "布雷叙尔",
		TitleCode: "b_bressuire",
	}
}
