package kherla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安禅那伽罗摩AnjangaonBarony struct {
	feud.BaseBarony
}

var BAnjangaon安禅那伽罗摩 feud.Barony = &安禅那伽罗摩AnjangaonBarony{}

func init() {
    f := BAnjangaon安禅那伽罗摩.(*安禅那伽罗摩AnjangaonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anjangaon",
		TitleName: "安禅那伽罗摩",
		TitleCode: "b_anjangaon",
	}
}
