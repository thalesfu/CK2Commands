package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桂花园GuihuayuanBarony struct {
	feud.BaseBarony
}

var BGuihuayuan桂花园 feud.Barony = &桂花园GuihuayuanBarony{}

func init() {
	f := BGuihuayuan桂花园.(*桂花园GuihuayuanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guihuayuan",
		TitleName: "桂花园",
		TitleCode: "b_guihuayuan",
	}
}
