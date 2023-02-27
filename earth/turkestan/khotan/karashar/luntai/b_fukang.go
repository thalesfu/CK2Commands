package luntai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阜康FukangBarony struct {
	feud.BaseBarony
}

var BFukang阜康 feud.Barony = &阜康FukangBarony{}

func init() {
    f := BFukang阜康.(*阜康FukangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fukang",
		TitleName: "阜康",
		TitleCode: "b_fukang",
	}
}
