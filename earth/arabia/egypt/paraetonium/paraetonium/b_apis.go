package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿庇斯ApisBarony struct {
	feud.BaseBarony
}

var BApis阿庇斯 feud.Barony = &阿庇斯ApisBarony{}

func init() {
    f := BApis阿庇斯.(*阿庇斯ApisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apis",
		TitleName: "阿庇斯",
		TitleCode: "b_apis",
	}
}
