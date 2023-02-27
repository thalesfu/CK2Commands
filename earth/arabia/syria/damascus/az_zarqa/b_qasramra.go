package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆拉堡QasramraBarony struct {
	feud.BaseBarony
}

var BQasramra阿姆拉堡 feud.Barony = &阿姆拉堡QasramraBarony{}

func init() {
    f := BQasramra阿姆拉堡.(*阿姆拉堡QasramraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qasramra",
		TitleName: "阿姆拉堡",
		TitleCode: "b_qasramra",
	}
}
