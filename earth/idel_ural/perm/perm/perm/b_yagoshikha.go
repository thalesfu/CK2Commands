package perm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚戈希哈YagoshikhaBarony struct {
	feud.BaseBarony
}

var BYagoshikha亚戈希哈 feud.Barony = &亚戈希哈YagoshikhaBarony{}

func init() {
    f := BYagoshikha亚戈希哈.(*亚戈希哈YagoshikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yagoshikha",
		TitleName: "亚戈希哈",
		TitleCode: "b_yagoshikha",
	}
}
