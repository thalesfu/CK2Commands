package finnmark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗姆瑟TromsoBarony struct {
	feud.BaseBarony
}

var BTromso特罗姆瑟 feud.Barony = &特罗姆瑟TromsoBarony{}

func init() {
    f := BTromso特罗姆瑟.(*特罗姆瑟TromsoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tromso",
		TitleName: "特罗姆瑟",
		TitleCode: "b_tromso",
	}
}
