package saris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔特褔BartfaBarony struct {
	feud.BaseBarony
}

var BBartfa巴尔特褔 feud.Barony = &巴尔特褔BartfaBarony{}

func init() {
	f := BBartfa巴尔特褔.(*巴尔特褔BartfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bartfa",
		TitleName: "巴尔特褔",
		TitleCode: "b_bartfa",
	}
}
