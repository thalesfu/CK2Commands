package kexholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯基萨尔米KakisalmiBarony struct {
	feud.BaseBarony
}

var BKakisalmi凯基萨尔米 feud.Barony = &凯基萨尔米KakisalmiBarony{}

func init() {
	f := BKakisalmi凯基萨尔米.(*凯基萨尔米KakisalmiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kakisalmi",
		TitleName: "凯基萨尔米",
		TitleCode: "b_kakisalmi",
	}
}
