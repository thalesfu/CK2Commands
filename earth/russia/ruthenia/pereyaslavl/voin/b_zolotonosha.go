package voin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐洛托诺沙ZolotonoshaBarony struct {
	feud.BaseBarony
}

var BZolotonosha佐洛托诺沙 feud.Barony = &佐洛托诺沙ZolotonoshaBarony{}

func init() {
	f := BZolotonosha佐洛托诺沙.(*佐洛托诺沙ZolotonoshaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zolotonosha",
		TitleName: "佐洛托诺沙",
		TitleCode: "b_zolotonosha",
	}
}
