package naberezhnye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布古利马BugulmaBarony struct {
	feud.BaseBarony
}

var BBugulma布古利马 feud.Barony = &布古利马BugulmaBarony{}

func init() {
    f := BBugulma布古利马.(*布古利马BugulmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bugulma",
		TitleName: "布古利马",
		TitleCode: "b_bugulma",
	}
}
