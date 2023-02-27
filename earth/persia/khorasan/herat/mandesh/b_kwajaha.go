package mandesh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科瓦贾哈KwajahaBarony struct {
	feud.BaseBarony
}

var BKwajaha科瓦贾哈 feud.Barony = &科瓦贾哈KwajahaBarony{}

func init() {
    f := BKwajaha科瓦贾哈.(*科瓦贾哈KwajahaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kwajaha",
		TitleName: "科瓦贾哈",
		TitleCode: "b_kwajaha",
	}
}
