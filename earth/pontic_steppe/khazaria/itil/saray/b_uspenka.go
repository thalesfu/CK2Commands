package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯片卡UspenkaBarony struct {
	feud.BaseBarony
}

var BUspenka乌斯片卡 feud.Barony = &乌斯片卡UspenkaBarony{}

func init() {
    f := BUspenka乌斯片卡.(*乌斯片卡UspenkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uspenka",
		TitleName: "乌斯片卡",
		TitleCode: "b_uspenka",
	}
}
