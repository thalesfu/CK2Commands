package mezen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基姆扎KimzhaBarony struct {
	feud.BaseBarony
}

var BKimzha基姆扎 feud.Barony = &基姆扎KimzhaBarony{}

func init() {
    f := BKimzha基姆扎.(*基姆扎KimzhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kimzha",
		TitleName: "基姆扎",
		TitleCode: "b_kimzha",
	}
}
