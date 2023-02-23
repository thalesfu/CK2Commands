package vatsagulma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VatsagulmaCounty interface {
	feud.County
	BJalna阇罗那() feud.Barony
	BLonar卢那罗() feud.Barony
	BManora摩奴罗() feud.Barony
	BNandura难豆罗() feud.Barony
	BPuduchattiram富菟遮底蓝() feud.Barony
	BSatgaon娑多伽罗摩() feud.Barony
	BVatsagulma婆蹉崛摩() feud.Barony
}

type 婆蹉崛摩VatsagulmaCounty struct {
	feud.BaseCounty
	阇罗那Jalna           feud.Barony
	卢那罗Lonar           feud.Barony
	摩奴罗Manora          feud.Barony
	难豆罗Nandura         feud.Barony
	富菟遮底蓝Puduchattiram feud.Barony
	娑多伽罗摩Satgaon       feud.Barony
	婆蹉崛摩Vatsagulma     feud.Barony
}

func (c *婆蹉崛摩VatsagulmaCounty) BJalna阇罗那() feud.Barony {
	return c.阇罗那Jalna
}

func (c *婆蹉崛摩VatsagulmaCounty) BLonar卢那罗() feud.Barony {
	return c.卢那罗Lonar
}

func (c *婆蹉崛摩VatsagulmaCounty) BManora摩奴罗() feud.Barony {
	return c.摩奴罗Manora
}

func (c *婆蹉崛摩VatsagulmaCounty) BNandura难豆罗() feud.Barony {
	return c.难豆罗Nandura
}

func (c *婆蹉崛摩VatsagulmaCounty) BPuduchattiram富菟遮底蓝() feud.Barony {
	return c.富菟遮底蓝Puduchattiram
}

func (c *婆蹉崛摩VatsagulmaCounty) BSatgaon娑多伽罗摩() feud.Barony {
	return c.娑多伽罗摩Satgaon
}

func (c *婆蹉崛摩VatsagulmaCounty) BVatsagulma婆蹉崛摩() feud.Barony {
	return c.婆蹉崛摩Vatsagulma
}

var CVatsagulma婆蹉崛摩 VatsagulmaCounty = &婆蹉崛摩VatsagulmaCounty{}

func init() {
	f := CVatsagulma婆蹉崛摩.(*婆蹉崛摩VatsagulmaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1259",
		Title:     "vatsagulma",
		TitleName: "婆蹉崛摩",
		TitleCode: "c_vatsagulma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阇罗那Jalna = BJalna阇罗那
	f.阇罗那Jalna.SetParent(f)

	f.卢那罗Lonar = BLonar卢那罗
	f.卢那罗Lonar.SetParent(f)

	f.摩奴罗Manora = BManora摩奴罗
	f.摩奴罗Manora.SetParent(f)

	f.难豆罗Nandura = BNandura难豆罗
	f.难豆罗Nandura.SetParent(f)

	f.富菟遮底蓝Puduchattiram = BPuduchattiram富菟遮底蓝
	f.富菟遮底蓝Puduchattiram.SetParent(f)

	f.娑多伽罗摩Satgaon = BSatgaon娑多伽罗摩
	f.娑多伽罗摩Satgaon.SetParent(f)

	f.婆蹉崛摩Vatsagulma = BVatsagulma婆蹉崛摩
	f.婆蹉崛摩Vatsagulma.SetParent(f)

}
