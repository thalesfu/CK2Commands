package turgay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔雷基UrlykiBarony struct {
	feud.BaseBarony
}

var BUrlyki乌尔雷基 feud.Barony = &乌尔雷基UrlykiBarony{}

func init() {
	f := BUrlyki乌尔雷基.(*乌尔雷基UrlykiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "urlyki",
		TitleName: "乌尔雷基",
		TitleCode: "b_urlyki",
	}
}
