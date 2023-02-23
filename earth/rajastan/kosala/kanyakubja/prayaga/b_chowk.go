package prayaga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 朝克ChowkBarony struct {
	feud.BaseBarony
}

var BChowk朝克 feud.Barony = &朝克ChowkBarony{}

func init() {
	f := BChowk朝克.(*朝克ChowkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chowk",
		TitleName: "朝克",
		TitleCode: "b_chowk",
	}
}
