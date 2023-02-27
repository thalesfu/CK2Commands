package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博乔尔马BochormaBarony struct {
	feud.BaseBarony
}

var BBochorma博乔尔马 feud.Barony = &博乔尔马BochormaBarony{}

func init() {
    f := BBochorma博乔尔马.(*博乔尔马BochormaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bochorma",
		TitleName: "博乔尔马",
		TitleCode: "b_bochorma",
	}
}
