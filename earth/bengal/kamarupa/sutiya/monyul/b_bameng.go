package monyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝门BamengBarony struct {
	feud.BaseBarony
}

var BBameng贝门 feud.Barony = &贝门BamengBarony{}

func init() {
	f := BBameng贝门.(*贝门BamengBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bameng",
		TitleName: "贝门",
		TitleCode: "b_bameng",
	}
}
