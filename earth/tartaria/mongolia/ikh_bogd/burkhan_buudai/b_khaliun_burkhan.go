package burkhan_buudai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈里温Khaliun_burkhanBarony struct {
	feud.BaseBarony
}

var BKhaliun_burkhan哈里温 feud.Barony = &哈里温Khaliun_burkhanBarony{}

func init() {
    f := BKhaliun_burkhan哈里温.(*哈里温Khaliun_burkhanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khaliun_burkhan",
		TitleName: "哈里温",
		TitleCode: "b_khaliun_burkhan",
	}
}
