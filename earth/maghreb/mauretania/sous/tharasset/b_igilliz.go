package tharasset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊吉利兹IgillizBarony struct {
	feud.BaseBarony
}

var BIgilliz伊吉利兹 feud.Barony = &伊吉利兹IgillizBarony{}

func init() {
    f := BIgilliz伊吉利兹.(*伊吉利兹IgillizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "igilliz",
		TitleName: "伊吉利兹",
		TitleCode: "b_igilliz",
	}
}
