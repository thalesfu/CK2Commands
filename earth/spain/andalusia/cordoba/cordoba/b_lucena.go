package cordoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢塞纳LucenaBarony struct {
	feud.BaseBarony
}

var BLucena卢塞纳 feud.Barony = &卢塞纳LucenaBarony{}

func init() {
	f := BLucena卢塞纳.(*卢塞纳LucenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lucena",
		TitleName: "卢塞纳",
		TitleCode: "b_lucena",
	}
}
