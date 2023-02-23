package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌里UriBarony struct {
	feud.BaseBarony
}

var BUri乌里 feud.Barony = &乌里UriBarony{}

func init() {
	f := BUri乌里.(*乌里UriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uri",
		TitleName: "乌里",
		TitleCode: "b_uri",
	}
}
