package montbeliard

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦朗蒂涅ValentigneyBarony struct {
	feud.BaseBarony
}

var BValentigney瓦朗蒂涅 feud.Barony = &瓦朗蒂涅ValentigneyBarony{}

func init() {
    f := BValentigney瓦朗蒂涅.(*瓦朗蒂涅ValentigneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valentigney",
		TitleName: "瓦朗蒂涅",
		TitleCode: "b_valentigney",
	}
}
