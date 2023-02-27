package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安克拉姆AnklamBarony struct {
	feud.BaseBarony
}

var BAnklam安克拉姆 feud.Barony = &安克拉姆AnklamBarony{}

func init() {
    f := BAnklam安克拉姆.(*安克拉姆AnklamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anklam",
		TitleName: "安克拉姆",
		TitleCode: "b_anklam",
	}
}
