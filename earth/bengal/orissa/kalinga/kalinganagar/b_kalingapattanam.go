package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯陵伽波多那KalingapattanamBarony struct {
	feud.BaseBarony
}

var BKalingapattanam羯陵伽波多那 feud.Barony = &羯陵伽波多那KalingapattanamBarony{}

func init() {
    f := BKalingapattanam羯陵伽波多那.(*羯陵伽波多那KalingapattanamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kalingapattanam",
		TitleName: "羯陵伽波多那",
		TitleCode: "b_kalingapattanam",
	}
}
