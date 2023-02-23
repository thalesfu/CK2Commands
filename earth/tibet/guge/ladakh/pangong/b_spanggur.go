package pangong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯潘古尔SpanggurBarony struct {
	feud.BaseBarony
}

var BSpanggur斯潘古尔 feud.Barony = &斯潘古尔SpanggurBarony{}

func init() {
	f := BSpanggur斯潘古尔.(*斯潘古尔SpanggurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spanggur",
		TitleName: "斯潘古尔",
		TitleCode: "b_spanggur",
	}
}
