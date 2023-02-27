package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 顺古特ShungutBarony struct {
	feud.BaseBarony
}

var BShungut顺古特 feud.Barony = &顺古特ShungutBarony{}

func init() {
    f := BShungut顺古特.(*顺古特ShungutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shungut",
		TitleName: "顺古特",
		TitleCode: "b_shungut",
	}
}
