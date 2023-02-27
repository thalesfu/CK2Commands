package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉登巴赫GladenbachBarony struct {
	feud.BaseBarony
}

var BGladenbach格拉登巴赫 feud.Barony = &格拉登巴赫GladenbachBarony{}

func init() {
    f := BGladenbach格拉登巴赫.(*格拉登巴赫GladenbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gladenbach",
		TitleName: "格拉登巴赫",
		TitleCode: "b_gladenbach",
	}
}
