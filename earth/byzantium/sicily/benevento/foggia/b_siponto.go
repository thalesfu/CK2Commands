package foggia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西蓬托SipontoBarony struct {
	feud.BaseBarony
}

var BSiponto西蓬托 feud.Barony = &西蓬托SipontoBarony{}

func init() {
	f := BSiponto西蓬托.(*西蓬托SipontoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siponto",
		TitleName: "西蓬托",
		TitleCode: "b_siponto",
	}
}
