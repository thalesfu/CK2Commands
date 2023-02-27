package kara_bogaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格克赛GeksayBarony struct {
	feud.BaseBarony
}

var BGeksay格克赛 feud.Barony = &格克赛GeksayBarony{}

func init() {
    f := BGeksay格克赛.(*格克赛GeksayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geksay",
		TitleName: "格克赛",
		TitleCode: "b_geksay",
	}
}
