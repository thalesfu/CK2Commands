package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库福尔登CoevordenBarony struct {
	feud.BaseBarony
}

var BCoevorden库福尔登 feud.Barony = &库福尔登CoevordenBarony{}

func init() {
	f := BCoevorden库福尔登.(*库福尔登CoevordenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coevorden",
		TitleName: "库福尔登",
		TitleCode: "b_coevorden",
	}
}
