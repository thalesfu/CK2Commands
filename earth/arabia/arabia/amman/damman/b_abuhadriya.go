package damman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布哈德里耶AbuhadriyaBarony struct {
	feud.BaseBarony
}

var BAbuhadriya阿布哈德里耶 feud.Barony = &阿布哈德里耶AbuhadriyaBarony{}

func init() {
	f := BAbuhadriya阿布哈德里耶.(*阿布哈德里耶AbuhadriyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abuhadriya",
		TitleName: "阿布哈德里耶",
		TitleCode: "b_abuhadriya",
	}
}
