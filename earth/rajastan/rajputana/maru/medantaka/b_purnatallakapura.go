package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富楼那呾罗迦补罗PurnatallakapuraBarony struct {
	feud.BaseBarony
}

var BPurnatallakapura富楼那呾罗迦补罗 feud.Barony = &富楼那呾罗迦补罗PurnatallakapuraBarony{}

func init() {
	f := BPurnatallakapura富楼那呾罗迦补罗.(*富楼那呾罗迦补罗PurnatallakapuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "purnatallakapura",
		TitleName: "富楼那呾罗迦补罗",
		TitleCode: "b_purnatallakapura",
	}
}
