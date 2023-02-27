package bure

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡吉里SiguiriBarony struct {
	feud.BaseBarony
}

var BSiguiri锡吉里 feud.Barony = &锡吉里SiguiriBarony{}

func init() {
    f := BSiguiri锡吉里.(*锡吉里SiguiriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siguiri",
		TitleName: "锡吉里",
		TitleCode: "b_siguiri",
	}
}
