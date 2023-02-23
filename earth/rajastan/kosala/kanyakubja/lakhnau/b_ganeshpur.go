package lakhnau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽泥舍补罗GaneshpurBarony struct {
	feud.BaseBarony
}

var BGaneshpur伽泥舍补罗 feud.Barony = &伽泥舍补罗GaneshpurBarony{}

func init() {
	f := BGaneshpur伽泥舍补罗.(*伽泥舍补罗GaneshpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ganeshpur",
		TitleName: "伽泥舍补罗",
		TitleCode: "b_ganeshpur",
	}
}
