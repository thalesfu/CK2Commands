package barkul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 小关田XiaoguantianBarony struct {
	feud.BaseBarony
}

var BXiaoguantian小关田 feud.Barony = &小关田XiaoguantianBarony{}

func init() {
    f := BXiaoguantian小关田.(*小关田XiaoguantianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xiaoguantian",
		TitleName: "小关田",
		TitleCode: "b_xiaoguantian",
	}
}
