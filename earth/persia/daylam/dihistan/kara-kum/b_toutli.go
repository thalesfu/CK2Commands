package kara-kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托特利ToutliBarony struct {
	feud.BaseBarony
}

var BToutli托特利 feud.Barony = &托特利ToutliBarony{}

func init() {
    f := BToutli托特利.(*托特利ToutliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toutli",
		TitleName: "托特利",
		TitleCode: "b_toutli",
	}
}
