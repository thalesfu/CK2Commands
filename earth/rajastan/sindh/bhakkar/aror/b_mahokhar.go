package aror

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩呼伽MahokharBarony struct {
	feud.BaseBarony
}

var BMahokhar摩呼伽 feud.Barony = &摩呼伽MahokharBarony{}

func init() {
    f := BMahokhar摩呼伽.(*摩呼伽MahokharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mahokhar",
		TitleName: "摩呼伽",
		TitleCode: "b_mahokhar",
	}
}
