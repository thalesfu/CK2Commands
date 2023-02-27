package darum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴桑AbasanBarony struct {
	feud.BaseBarony
}

var BAbasan阿巴桑 feud.Barony = &阿巴桑AbasanBarony{}

func init() {
    f := BAbasan阿巴桑.(*阿巴桑AbasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abasan",
		TitleName: "阿巴桑",
		TitleCode: "b_abasan",
	}
}
