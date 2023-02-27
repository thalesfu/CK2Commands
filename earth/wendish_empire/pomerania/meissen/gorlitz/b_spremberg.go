package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施普伦贝格SprembergBarony struct {
	feud.BaseBarony
}

var BSpremberg施普伦贝格 feud.Barony = &施普伦贝格SprembergBarony{}

func init() {
    f := BSpremberg施普伦贝格.(*施普伦贝格SprembergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spremberg",
		TitleName: "施普伦贝格",
		TitleCode: "b_spremberg",
	}
}
