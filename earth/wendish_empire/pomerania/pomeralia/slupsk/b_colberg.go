package slupsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔贝格ColbergBarony struct {
	feud.BaseBarony
}

var BColberg科尔贝格 feud.Barony = &科尔贝格ColbergBarony{}

func init() {
    f := BColberg科尔贝格.(*科尔贝格ColbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "colberg",
		TitleName: "科尔贝格",
		TitleCode: "b_colberg",
	}
}
