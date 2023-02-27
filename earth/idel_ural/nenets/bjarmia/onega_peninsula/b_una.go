package onega_peninsula

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌纳UnaBarony struct {
	feud.BaseBarony
}

var BUna乌纳 feud.Barony = &乌纳UnaBarony{}

func init() {
    f := BUna乌纳.(*乌纳UnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "una",
		TitleName: "乌纳",
		TitleCode: "b_una",
	}
}
