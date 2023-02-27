package manych

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔戈别克_巴尔卡MalgobekbalkaBarony struct {
	feud.BaseBarony
}

var BMalgobekbalka马尔戈别克_巴尔卡 feud.Barony = &马尔戈别克_巴尔卡MalgobekbalkaBarony{}

func init() {
    f := BMalgobekbalka马尔戈别克_巴尔卡.(*马尔戈别克_巴尔卡MalgobekbalkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malgobekbalka",
		TitleName: "马尔戈别克_巴尔卡",
		TitleCode: "b_malgobekbalka",
	}
}
