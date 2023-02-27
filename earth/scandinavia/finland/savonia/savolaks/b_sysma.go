package savolaks

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叙斯迈SysmaBarony struct {
	feud.BaseBarony
}

var BSysma叙斯迈 feud.Barony = &叙斯迈SysmaBarony{}

func init() {
    f := BSysma叙斯迈.(*叙斯迈SysmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sysma",
		TitleName: "叙斯迈",
		TitleCode: "b_sysma",
	}
}
