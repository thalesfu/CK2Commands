package trent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里申BrixenBarony struct {
	feud.BaseBarony
}

var BBrixen布里申 feud.Barony = &布里申BrixenBarony{}

func init() {
    f := BBrixen布里申.(*布里申BrixenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brixen",
		TitleName: "布里申",
		TitleCode: "b_brixen",
	}
}
