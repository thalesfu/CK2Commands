package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩尔格基兴PergkirchenBarony struct {
	feud.BaseBarony
}

var BPergkirchen佩尔格基兴 feud.Barony = &佩尔格基兴PergkirchenBarony{}

func init() {
    f := BPergkirchen佩尔格基兴.(*佩尔格基兴PergkirchenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pergkirchen",
		TitleName: "佩尔格基兴",
		TitleCode: "b_pergkirchen",
	}
}
