package kanchipuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多伽卢摩TakkalomaBarony struct {
	feud.BaseBarony
}

var BTakkaloma多伽卢摩 feud.Barony = &多伽卢摩TakkalomaBarony{}

func init() {
    f := BTakkaloma多伽卢摩.(*多伽卢摩TakkalomaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "takkaloma",
		TitleName: "多伽卢摩",
		TitleCode: "b_takkaloma",
	}
}
