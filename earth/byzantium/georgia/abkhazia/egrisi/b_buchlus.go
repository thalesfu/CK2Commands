package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布赫卢斯BuchlusBarony struct {
	feud.BaseBarony
}

var BBuchlus布赫卢斯 feud.Barony = &布赫卢斯BuchlusBarony{}

func init() {
    f := BBuchlus布赫卢斯.(*布赫卢斯BuchlusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buchlus",
		TitleName: "布赫卢斯",
		TitleCode: "b_buchlus",
	}
}
