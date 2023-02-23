package yperen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗泽贝克RosebekeBarony struct {
	feud.BaseBarony
}

var BRosebeke罗泽贝克 feud.Barony = &罗泽贝克RosebekeBarony{}

func init() {
	f := BRosebeke罗泽贝克.(*罗泽贝克RosebekeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosebeke",
		TitleName: "罗泽贝克",
		TitleCode: "b_rosebeke",
	}
}
