package agder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里姆斯塔GrimstadBarony struct {
	feud.BaseBarony
}

var BGrimstad格里姆斯塔 feud.Barony = &格里姆斯塔GrimstadBarony{}

func init() {
	f := BGrimstad格里姆斯塔.(*格里姆斯塔GrimstadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grimstad",
		TitleName: "格里姆斯塔",
		TitleCode: "b_grimstad",
	}
}
