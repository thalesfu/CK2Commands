package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇摩醯因陀罗伐蓝RajamahendravaramBarony struct {
	feud.BaseBarony
}

var BRajamahendravaram罗阇摩醯因陀罗伐蓝 feud.Barony = &罗阇摩醯因陀罗伐蓝RajamahendravaramBarony{}

func init() {
    f := BRajamahendravaram罗阇摩醯因陀罗伐蓝.(*罗阇摩醯因陀罗伐蓝RajamahendravaramBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajamahendravaram",
		TitleName: "罗阇摩醯因陀罗伐蓝",
		TitleCode: "b_rajamahendravaram",
	}
}
