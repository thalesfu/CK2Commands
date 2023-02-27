package aqtobe

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎尔雷_赛ZharlysayBarony struct {
	feud.BaseBarony
}

var BZharlysay扎尔雷_赛 feud.Barony = &扎尔雷_赛ZharlysayBarony{}

func init() {
    f := BZharlysay扎尔雷_赛.(*扎尔雷_赛ZharlysayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zharlysay",
		TitleName: "扎尔雷_赛",
		TitleCode: "b_zharlysay",
	}
}
