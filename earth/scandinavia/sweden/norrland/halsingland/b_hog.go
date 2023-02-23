package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫格HogBarony struct {
	feud.BaseBarony
}

var BHog赫格 feud.Barony = &赫格HogBarony{}

func init() {
	f := BHog赫格.(*赫格HogBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hog",
		TitleName: "赫格",
		TitleCode: "b_hog",
	}
}
