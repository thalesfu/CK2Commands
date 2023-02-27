package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌德比纳UdbinaBarony struct {
	feud.BaseBarony
}

var BUdbina乌德比纳 feud.Barony = &乌德比纳UdbinaBarony{}

func init() {
    f := BUdbina乌德比纳.(*乌德比纳UdbinaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udbina",
		TitleName: "乌德比纳",
		TitleCode: "b_udbina",
	}
}
