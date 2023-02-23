package acalapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔那拉NarnalaBarony struct {
	feud.BaseBarony
}

var BNarnala纳尔那拉 feud.Barony = &纳尔那拉NarnalaBarony{}

func init() {
	f := BNarnala纳尔那拉.(*纳尔那拉NarnalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "narnala",
		TitleName: "纳尔那拉",
		TitleCode: "b_narnala",
	}
}
