package barskhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利片卡LipenkaBarony struct {
	feud.BaseBarony
}

var BLipenka利片卡 feud.Barony = &利片卡LipenkaBarony{}

func init() {
	f := BLipenka利片卡.(*利片卡LipenkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lipenka",
		TitleName: "利片卡",
		TitleCode: "b_lipenka",
	}
}
