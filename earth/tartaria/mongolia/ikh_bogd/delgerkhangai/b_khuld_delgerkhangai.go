package delgerkhangai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼勒德Khuld_delgerkhangaiBarony struct {
	feud.BaseBarony
}

var BKhuld_delgerkhangai呼勒德 feud.Barony = &呼勒德Khuld_delgerkhangaiBarony{}

func init() {
    f := BKhuld_delgerkhangai呼勒德.(*呼勒德Khuld_delgerkhangaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khuld_delgerkhangai",
		TitleName: "呼勒德",
		TitleCode: "b_khuld_delgerkhangai",
	}
}
