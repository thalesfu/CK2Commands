package provence

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里莫GrimaudBarony struct {
	feud.BaseBarony
}

var BGrimaud格里莫 feud.Barony = &格里莫GrimaudBarony{}

func init() {
	f := BGrimaud格里莫.(*格里莫GrimaudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grimaud",
		TitleName: "格里莫",
		TitleCode: "b_grimaud",
	}
}
