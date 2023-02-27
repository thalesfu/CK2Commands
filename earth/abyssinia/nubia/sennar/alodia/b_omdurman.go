package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩图曼OmdurmanBarony struct {
	feud.BaseBarony
}

var BOmdurman恩图曼 feud.Barony = &恩图曼OmdurmanBarony{}

func init() {
    f := BOmdurman恩图曼.(*恩图曼OmdurmanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "omdurman",
		TitleName: "恩图曼",
		TitleCode: "b_omdurman",
	}
}
