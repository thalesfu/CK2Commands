package reval

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博雷阿尔BorpealBarony struct {
	feud.BaseBarony
}

var BBorpeal博雷阿尔 feud.Barony = &博雷阿尔BorpealBarony{}

func init() {
	f := BBorpeal博雷阿尔.(*博雷阿尔BorpealBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borpeal",
		TitleName: "博雷阿尔",
		TitleCode: "b_borpeal",
	}
}
