package uwal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉达JiddaBarony struct {
	feud.BaseBarony
}

var BJidda吉达 feud.Barony = &吉达JiddaBarony{}

func init() {
    f := BJidda吉达.(*吉达JiddaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jidda",
		TitleName: "吉达",
		TitleCode: "b_jidda",
	}
}
