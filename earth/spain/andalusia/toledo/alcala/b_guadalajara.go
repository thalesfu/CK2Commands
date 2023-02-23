package alcala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜达拉哈拉GuadalajaraBarony struct {
	feud.BaseBarony
}

var BGuadalajara瓜达拉哈拉 feud.Barony = &瓜达拉哈拉GuadalajaraBarony{}

func init() {
	f := BGuadalajara瓜达拉哈拉.(*瓜达拉哈拉GuadalajaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guadalajara",
		TitleName: "瓜达拉哈拉",
		TitleCode: "b_guadalajara",
	}
}
