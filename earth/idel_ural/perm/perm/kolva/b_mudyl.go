package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆德利MudylBarony struct {
	feud.BaseBarony
}

var BMudyl穆德利 feud.Barony = &穆德利MudylBarony{}

func init() {
    f := BMudyl穆德利.(*穆德利MudylBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mudyl",
		TitleName: "穆德利",
		TitleCode: "b_mudyl",
	}
}
