package homs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔马里塔MarmaritaBarony struct {
	feud.BaseBarony
}

var BMarmarita马尔马里塔 feud.Barony = &马尔马里塔MarmaritaBarony{}

func init() {
    f := BMarmarita马尔马里塔.(*马尔马里塔MarmaritaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marmarita",
		TitleName: "马尔马里塔",
		TitleCode: "b_marmarita",
	}
}
