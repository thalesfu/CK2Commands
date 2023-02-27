package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾兹普泰AizputeBarony struct {
	feud.BaseBarony
}

var BAizpute艾兹普泰 feud.Barony = &艾兹普泰AizputeBarony{}

func init() {
    f := BAizpute艾兹普泰.(*艾兹普泰AizputeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aizpute",
		TitleName: "艾兹普泰",
		TitleCode: "b_aizpute",
	}
}
