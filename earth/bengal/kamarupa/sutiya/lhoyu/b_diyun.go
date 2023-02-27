package lhoyu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂芸DiyunBarony struct {
	feud.BaseBarony
}

var BDiyun蒂芸 feud.Barony = &蒂芸DiyunBarony{}

func init() {
    f := BDiyun蒂芸.(*蒂芸DiyunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "diyun",
		TitleName: "蒂芸",
		TitleCode: "b_diyun",
	}
}
