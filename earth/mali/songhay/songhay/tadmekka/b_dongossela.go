package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东戈塞拉DongosselaBarony struct {
	feud.BaseBarony
}

var BDongossela东戈塞拉 feud.Barony = &东戈塞拉DongosselaBarony{}

func init() {
    f := BDongossela东戈塞拉.(*东戈塞拉DongosselaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dongossela",
		TitleName: "东戈塞拉",
		TitleCode: "b_dongossela",
	}
}
