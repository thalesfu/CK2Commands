package aden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考德AlkawdBarony struct {
	feud.BaseBarony
}

var BAlkawd考德 feud.Barony = &考德AlkawdBarony{}

func init() {
    f := BAlkawd考德.(*考德AlkawdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alkawd",
		TitleName: "考德",
		TitleCode: "b_alkawd",
	}
}
