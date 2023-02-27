package nakhshab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博扎里克BozariqBarony struct {
	feud.BaseBarony
}

var BBozariq博扎里克 feud.Barony = &博扎里克BozariqBarony{}

func init() {
    f := BBozariq博扎里克.(*博扎里克BozariqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bozariq",
		TitleName: "博扎里克",
		TitleCode: "b_bozariq",
	}
}
