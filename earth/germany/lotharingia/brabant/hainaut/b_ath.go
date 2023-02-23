package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特AthBarony struct {
	feud.BaseBarony
}

var BAth阿特 feud.Barony = &阿特AthBarony{}

func init() {
	f := BAth阿特.(*阿特AthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ath",
		TitleName: "阿特",
		TitleCode: "b_ath",
	}
}
