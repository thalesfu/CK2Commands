package sonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莫乔特AmarkotBarony struct {
	feud.BaseBarony
}

var BAmarkot阿莫乔特 feud.Barony = &阿莫乔特AmarkotBarony{}

func init() {
	f := BAmarkot阿莫乔特.(*阿莫乔特AmarkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amarkot",
		TitleName: "阿莫乔特",
		TitleCode: "b_amarkot",
	}
}
