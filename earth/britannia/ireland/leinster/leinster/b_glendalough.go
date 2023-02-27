package leinster

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格兰达洛GlendaloughBarony struct {
	feud.BaseBarony
}

var BGlendalough格兰达洛 feud.Barony = &格兰达洛GlendaloughBarony{}

func init() {
    f := BGlendalough格兰达洛.(*格兰达洛GlendaloughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glendalough",
		TitleName: "格兰达洛",
		TitleCode: "b_glendalough",
	}
}
