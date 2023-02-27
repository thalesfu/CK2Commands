package chernigov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅纳MenaBarony struct {
	feud.BaseBarony
}

var BMena梅纳 feud.Barony = &梅纳MenaBarony{}

func init() {
    f := BMena梅纳.(*梅纳MenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mena",
		TitleName: "梅纳",
		TitleCode: "b_mena",
	}
}
