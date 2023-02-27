package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斑波尔BanbhoreBarony struct {
	feud.BaseBarony
}

var BBanbhore斑波尔 feud.Barony = &斑波尔BanbhoreBarony{}

func init() {
    f := BBanbhore斑波尔.(*斑波尔BanbhoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "banbhore",
		TitleName: "斑波尔",
		TitleCode: "b_banbhore",
	}
}
