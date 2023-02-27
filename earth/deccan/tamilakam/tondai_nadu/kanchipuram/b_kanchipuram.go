package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 建志补罗KanchipuramBarony struct {
	feud.BaseBarony
}

var BKanchipuram建志补罗 feud.Barony = &建志补罗KanchipuramBarony{}

func init() {
    f := BKanchipuram建志补罗.(*建志补罗KanchipuramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanchipuram",
		TitleName: "建志补罗",
		TitleCode: "b_kanchipuram",
	}
}
