package lukomorie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呕基夫OnkhivBarony struct {
	feud.BaseBarony
}

var BOnkhiv呕基夫 feud.Barony = &呕基夫OnkhivBarony{}

func init() {
    f := BOnkhiv呕基夫.(*呕基夫OnkhivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "onkhiv",
		TitleName: "呕基夫",
		TitleCode: "b_onkhiv",
	}
}
