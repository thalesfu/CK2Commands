package dyrrachion

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/dyrrachion/aulon"
	"github.com/thalesfu/CK2Commands/earth/byzantium/epirus/dyrrachion/dyrrachion"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DyrrachionDuke interface {
	feud.Duke
	CAulon阿夫洛纳斯() aulon.AulonCounty
	CDyrrachion都拉齐翁() dyrrachion.DyrrachionCounty
}

type 都拉齐翁DyrrachionDuke struct {
	feud.BaseDuke
	阿夫洛纳斯Aulon     aulon.AulonCounty
	都拉齐翁Dyrrachion dyrrachion.DyrrachionCounty
}

func (d *都拉齐翁DyrrachionDuke) CAulon阿夫洛纳斯() aulon.AulonCounty {
	return d.阿夫洛纳斯Aulon
}

func (d *都拉齐翁DyrrachionDuke) CDyrrachion都拉齐翁() dyrrachion.DyrrachionCounty {
	return d.都拉齐翁Dyrrachion
}

var DDyrrachion都拉齐翁 DyrrachionDuke = &都拉齐翁DyrrachionDuke{}

func init() {
	f := DDyrrachion都拉齐翁.(*都拉齐翁DyrrachionDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dyrrachion",
		TitleName: "都拉齐翁",
		TitleCode: "d_dyrrachion",
		Counties:  map[string]feud.County{},
	}

	f.阿夫洛纳斯Aulon = aulon.CAulon阿夫洛纳斯
	f.阿夫洛纳斯Aulon.SetParent(f)

	f.都拉齐翁Dyrrachion = dyrrachion.CDyrrachion都拉齐翁
	f.都拉齐翁Dyrrachion.SetParent(f)

}
