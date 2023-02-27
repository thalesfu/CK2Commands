package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦尔扎扎特OuarzazateBarony struct {
	feud.BaseBarony
}

var BOuarzazate瓦尔扎扎特 feud.Barony = &瓦尔扎扎特OuarzazateBarony{}

func init() {
    f := BOuarzazate瓦尔扎扎特.(*瓦尔扎扎特OuarzazateBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouarzazate",
		TitleName: "瓦尔扎扎特",
		TitleCode: "b_ouarzazate",
	}
}
