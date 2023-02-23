package graz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉茨GrazBarony struct {
	feud.BaseBarony
}

var BGraz格拉茨 feud.Barony = &格拉茨GrazBarony{}

func init() {
	f := BGraz格拉茨.(*格拉茨GrazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "graz",
		TitleName: "格拉茨",
		TitleCode: "b_graz",
	}
}
