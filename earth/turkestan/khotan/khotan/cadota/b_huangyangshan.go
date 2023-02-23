package cadota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黄杨山HuangyangshanBarony struct {
	feud.BaseBarony
}

var BHuangyangshan黄杨山 feud.Barony = &黄杨山HuangyangshanBarony{}

func init() {
	f := BHuangyangshan黄杨山.(*黄杨山HuangyangshanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "huangyangshan",
		TitleName: "黄杨山",
		TitleCode: "b_huangyangshan",
	}
}
