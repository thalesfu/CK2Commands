package domazlice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗森贝格RosenbergBarony struct {
	feud.BaseBarony
}

var BRosenberg罗森贝格 feud.Barony = &罗森贝格RosenbergBarony{}

func init() {
    f := BRosenberg罗森贝格.(*罗森贝格RosenbergBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rosenberg",
		TitleName: "罗森贝格",
		TitleCode: "b_rosenberg",
	}
}
