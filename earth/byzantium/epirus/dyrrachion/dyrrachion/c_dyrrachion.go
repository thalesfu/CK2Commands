package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DyrrachionCounty interface {
	feud.County
	BBeat贝特() feud.Barony
	BChounavia库纳维亚() feud.Barony
	BDurazzo都拉斯() feud.Barony
	BElbasan埃尔巴桑() feud.Barony
	BGeziq格齐奇() feud.Barony
	BKruje克鲁耶() feud.Barony
	BPetrela佩特雷勒() feud.Barony
}

type 都拉齐翁DyrrachionCounty struct {
	feud.BaseCounty
	贝特Beat        feud.Barony
	库纳维亚Chounavia feud.Barony
	都拉斯Durazzo    feud.Barony
	埃尔巴桑Elbasan   feud.Barony
	格齐奇Geziq      feud.Barony
	克鲁耶Kruje      feud.Barony
	佩特雷勒Petrela   feud.Barony
}

func (c *都拉齐翁DyrrachionCounty) BBeat贝特() feud.Barony {
	return c.贝特Beat
}

func (c *都拉齐翁DyrrachionCounty) BChounavia库纳维亚() feud.Barony {
	return c.库纳维亚Chounavia
}

func (c *都拉齐翁DyrrachionCounty) BDurazzo都拉斯() feud.Barony {
	return c.都拉斯Durazzo
}

func (c *都拉齐翁DyrrachionCounty) BElbasan埃尔巴桑() feud.Barony {
	return c.埃尔巴桑Elbasan
}

func (c *都拉齐翁DyrrachionCounty) BGeziq格齐奇() feud.Barony {
	return c.格齐奇Geziq
}

func (c *都拉齐翁DyrrachionCounty) BKruje克鲁耶() feud.Barony {
	return c.克鲁耶Kruje
}

func (c *都拉齐翁DyrrachionCounty) BPetrela佩特雷勒() feud.Barony {
	return c.佩特雷勒Petrela
}

var CDyrrachion都拉齐翁 DyrrachionCounty = &都拉齐翁DyrrachionCounty{}

func init() {
	f := CDyrrachion都拉齐翁.(*都拉齐翁DyrrachionCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "470",
		Title:     "dyrrachion",
		TitleName: "都拉齐翁",
		TitleCode: "c_dyrrachion",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝特Beat = BBeat贝特
	f.贝特Beat.SetParent(f)

	f.库纳维亚Chounavia = BChounavia库纳维亚
	f.库纳维亚Chounavia.SetParent(f)

	f.都拉斯Durazzo = BDurazzo都拉斯
	f.都拉斯Durazzo.SetParent(f)

	f.埃尔巴桑Elbasan = BElbasan埃尔巴桑
	f.埃尔巴桑Elbasan.SetParent(f)

	f.格齐奇Geziq = BGeziq格齐奇
	f.格齐奇Geziq.SetParent(f)

	f.克鲁耶Kruje = BKruje克鲁耶
	f.克鲁耶Kruje.SetParent(f)

	f.佩特雷勒Petrela = BPetrela佩特雷勒
	f.佩特雷勒Petrela.SetParent(f)

}
