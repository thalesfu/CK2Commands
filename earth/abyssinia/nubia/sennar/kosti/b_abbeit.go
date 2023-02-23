package kosti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿贝特AbbeitBarony struct {
	feud.BaseBarony
}

var BAbbeit阿贝特 feud.Barony = &阿贝特AbbeitBarony{}

func init() {
	f := BAbbeit阿贝特.(*阿贝特AbbeitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abbeit",
		TitleName: "阿贝特",
		TitleCode: "b_abbeit",
	}
}
