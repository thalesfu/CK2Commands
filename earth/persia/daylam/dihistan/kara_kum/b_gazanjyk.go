package kara_kum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡赞吉克GazanjykBarony struct {
	feud.BaseBarony
}

var BGazanjyk卡赞吉克 feud.Barony = &卡赞吉克GazanjykBarony{}

func init() {
    f := BGazanjyk卡赞吉克.(*卡赞吉克GazanjykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gazanjyk",
		TitleName: "卡赞吉克",
		TitleCode: "b_gazanjyk",
	}
}
