package empuries

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉比斯瓦尔登波尔达LabisbaldempordaBarony struct {
	feud.BaseBarony
}

var BLabisbaldemporda拉比斯瓦尔登波尔达 feud.Barony = &拉比斯瓦尔登波尔达LabisbaldempordaBarony{}

func init() {
	f := BLabisbaldemporda拉比斯瓦尔登波尔达.(*拉比斯瓦尔登波尔达LabisbaldempordaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "labisbaldemporda",
		TitleName: "拉比斯瓦尔登波尔达",
		TitleCode: "b_labisbaldemporda",
	}
}
