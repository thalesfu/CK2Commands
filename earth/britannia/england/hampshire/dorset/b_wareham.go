package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦勒姆WarehamBarony struct {
	feud.BaseBarony
}

var BWareham韦勒姆 feud.Barony = &韦勒姆WarehamBarony{}

func init() {
	f := BWareham韦勒姆.(*韦勒姆WarehamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wareham",
		TitleName: "韦勒姆",
		TitleCode: "b_wareham",
	}
}
