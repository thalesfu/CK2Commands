package atheniai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AtheniaiCounty interface {
	feud.County
	BAthens雅典() feud.Barony
	BDaphni达夫尼() feud.Barony
	BKarydi卡律季() feud.Barony
	BMarathon马拉松() feud.Barony
	BMegara墨伽拉() feud.Barony
	BPiraeaus比雷埃夫斯() feud.Barony
	BSalamis萨拉米斯() feud.Barony
	BSoula苏拉() feud.Barony
}

type 雅典AtheniaiCounty struct {
	feud.BaseCounty
	雅典Athens      feud.Barony
	达夫尼Daphni     feud.Barony
	卡律季Karydi     feud.Barony
	马拉松Marathon   feud.Barony
	墨伽拉Megara     feud.Barony
	比雷埃夫斯Piraeaus feud.Barony
	萨拉米斯Salamis   feud.Barony
	苏拉Soula       feud.Barony
}

func (c *雅典AtheniaiCounty) BAthens雅典() feud.Barony {
	return c.雅典Athens
}

func (c *雅典AtheniaiCounty) BDaphni达夫尼() feud.Barony {
	return c.达夫尼Daphni
}

func (c *雅典AtheniaiCounty) BKarydi卡律季() feud.Barony {
	return c.卡律季Karydi
}

func (c *雅典AtheniaiCounty) BMarathon马拉松() feud.Barony {
	return c.马拉松Marathon
}

func (c *雅典AtheniaiCounty) BMegara墨伽拉() feud.Barony {
	return c.墨伽拉Megara
}

func (c *雅典AtheniaiCounty) BPiraeaus比雷埃夫斯() feud.Barony {
	return c.比雷埃夫斯Piraeaus
}

func (c *雅典AtheniaiCounty) BSalamis萨拉米斯() feud.Barony {
	return c.萨拉米斯Salamis
}

func (c *雅典AtheniaiCounty) BSoula苏拉() feud.Barony {
	return c.苏拉Soula
}

var CAtheniai雅典 AtheniaiCounty = &雅典AtheniaiCounty{}

func init() {
	f := CAtheniai雅典.(*雅典AtheniaiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "482",
		Title:     "atheniai",
		TitleName: "雅典",
		TitleCode: "c_atheniai",
		Baronies:  map[string]feud.Barony{},
	}

	f.雅典Athens = BAthens雅典
	f.雅典Athens.SetParent(f)

	f.达夫尼Daphni = BDaphni达夫尼
	f.达夫尼Daphni.SetParent(f)

	f.卡律季Karydi = BKarydi卡律季
	f.卡律季Karydi.SetParent(f)

	f.马拉松Marathon = BMarathon马拉松
	f.马拉松Marathon.SetParent(f)

	f.墨伽拉Megara = BMegara墨伽拉
	f.墨伽拉Megara.SetParent(f)

	f.比雷埃夫斯Piraeaus = BPiraeaus比雷埃夫斯
	f.比雷埃夫斯Piraeaus.SetParent(f)

	f.萨拉米斯Salamis = BSalamis萨拉米斯
	f.萨拉米斯Salamis.SetParent(f)

	f.苏拉Soula = BSoula苏拉
	f.苏拉Soula.SetParent(f)

}
