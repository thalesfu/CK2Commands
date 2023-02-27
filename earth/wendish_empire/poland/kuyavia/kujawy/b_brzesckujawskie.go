package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库亚维地区布热希奇BrzesckujawskieBarony struct {
	feud.BaseBarony
}

var BBrzesckujawskie库亚维地区布热希奇 feud.Barony = &库亚维地区布热希奇BrzesckujawskieBarony{}

func init() {
    f := BBrzesckujawskie库亚维地区布热希奇.(*库亚维地区布热希奇BrzesckujawskieBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brzesckujawskie",
		TitleName: "库亚维地区布热希奇",
		TitleCode: "b_brzesckujawskie",
	}
}
