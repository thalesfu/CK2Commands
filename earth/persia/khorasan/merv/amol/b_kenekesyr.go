package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯内克斯尔KenekesyrBarony struct {
	feud.BaseBarony
}

var BKenekesyr凯内克斯尔 feud.Barony = &凯内克斯尔KenekesyrBarony{}

func init() {
    f := BKenekesyr凯内克斯尔.(*凯内克斯尔KenekesyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kenekesyr",
		TitleName: "凯内克斯尔",
		TitleCode: "b_kenekesyr",
	}
}
