package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 殊迦JhodgaBarony struct {
	feud.BaseBarony
}

var BJhodga殊迦 feud.Barony = &殊迦JhodgaBarony{}

func init() {
	f := BJhodga殊迦.(*殊迦JhodgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jhodga",
		TitleName: "殊迦",
		TitleCode: "b_jhodga",
	}
}
