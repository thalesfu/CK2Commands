package kumtag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 果佳GuojiaBarony struct {
	feud.BaseBarony
}

var BGuojia果佳 feud.Barony = &果佳GuojiaBarony{}

func init() {
	f := BGuojia果佳.(*果佳GuojiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guojia",
		TitleName: "果佳",
		TitleCode: "b_guojia",
	}
}
