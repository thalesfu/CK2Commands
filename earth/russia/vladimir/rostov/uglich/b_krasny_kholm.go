package uglich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 红霍尔姆Krasny_kholmBarony struct {
	feud.BaseBarony
}

var BKrasny_kholm红霍尔姆 feud.Barony = &红霍尔姆Krasny_kholmBarony{}

func init() {
    f := BKrasny_kholm红霍尔姆.(*红霍尔姆Krasny_kholmBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "krasny_kholm",
		TitleName: "红霍尔姆",
		TitleCode: "b_krasny_kholm",
	}
}
