package saintois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布里克塞BrixeyBarony struct {
	feud.BaseBarony
}

var BBrixey布里克塞 feud.Barony = &布里克塞BrixeyBarony{}

func init() {
    f := BBrixey布里克塞.(*布里克塞BrixeyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brixey",
		TitleName: "布里克塞",
		TitleCode: "b_brixey",
	}
}
