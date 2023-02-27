package kurumba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦里奥瓦里奥Ouario_ouarioBarony struct {
	feud.BaseBarony
}

var BOuario_ouario瓦里奥瓦里奥 feud.Barony = &瓦里奥瓦里奥Ouario_ouarioBarony{}

func init() {
    f := BOuario_ouario瓦里奥瓦里奥.(*瓦里奥瓦里奥Ouario_ouarioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouario_ouario",
		TitleName: "瓦里奥瓦里奥",
		TitleCode: "b_ouario_ouario",
	}
}
