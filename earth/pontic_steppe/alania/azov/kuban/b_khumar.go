package kuban

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡马尔KhumarBarony struct {
	feud.BaseBarony
}

var BKhumar胡马尔 feud.Barony = &胡马尔KhumarBarony{}

func init() {
    f := BKhumar胡马尔.(*胡马尔KhumarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khumar",
		TitleName: "胡马尔",
		TitleCode: "b_khumar",
	}
}
