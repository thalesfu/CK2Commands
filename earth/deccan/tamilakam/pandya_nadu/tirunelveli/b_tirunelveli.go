package tirunelveli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 底卢内勒吠梨TirunelveliBarony struct {
	feud.BaseBarony
}

var BTirunelveli底卢内勒吠梨 feud.Barony = &底卢内勒吠梨TirunelveliBarony{}

func init() {
    f := BTirunelveli底卢内勒吠梨.(*底卢内勒吠梨TirunelveliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tirunelveli",
		TitleName: "底卢内勒吠梨",
		TitleCode: "b_tirunelveli",
	}
}
