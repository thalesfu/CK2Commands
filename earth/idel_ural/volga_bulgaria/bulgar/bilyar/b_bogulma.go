package bilyar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布古利马BogulmaBarony struct {
	feud.BaseBarony
}

var BBogulma布古利马 feud.Barony = &布古利马BogulmaBarony{}

func init() {
    f := BBogulma布古利马.(*布古利马BogulmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogulma",
		TitleName: "布古利马",
		TitleCode: "b_bogulma",
	}
}
