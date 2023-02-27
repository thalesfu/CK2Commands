package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔卢克UrlukBarony struct {
	feud.BaseBarony
}

var BUrluk乌尔卢克 feud.Barony = &乌尔卢克UrlukBarony{}

func init() {
    f := BUrluk乌尔卢克.(*乌尔卢克UrlukBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urluk",
		TitleName: "乌尔卢克",
		TitleCode: "b_urluk",
	}
}
