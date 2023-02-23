package alodia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlodiaCounty interface {
	feud.County
	BAlwa阿勒瓦() feud.Barony
	BBarah巴拉() feud.Barony
	BMalaka马拉卡() feud.Barony
	BOmdurman恩图曼() feud.Barony
	BOmjerky欧米茄尔可() feud.Barony
	BSoba索巴() feud.Barony
	BWawissi瓦维斯() feud.Barony
}

type 阿洛地亚AlodiaCounty struct {
	feud.BaseCounty
	阿勒瓦Alwa      feud.Barony
	巴拉Barah      feud.Barony
	马拉卡Malaka    feud.Barony
	恩图曼Omdurman  feud.Barony
	欧米茄尔可Omjerky feud.Barony
	索巴Soba       feud.Barony
	瓦维斯Wawissi   feud.Barony
}

func (c *阿洛地亚AlodiaCounty) BAlwa阿勒瓦() feud.Barony {
	return c.阿勒瓦Alwa
}

func (c *阿洛地亚AlodiaCounty) BBarah巴拉() feud.Barony {
	return c.巴拉Barah
}

func (c *阿洛地亚AlodiaCounty) BMalaka马拉卡() feud.Barony {
	return c.马拉卡Malaka
}

func (c *阿洛地亚AlodiaCounty) BOmdurman恩图曼() feud.Barony {
	return c.恩图曼Omdurman
}

func (c *阿洛地亚AlodiaCounty) BOmjerky欧米茄尔可() feud.Barony {
	return c.欧米茄尔可Omjerky
}

func (c *阿洛地亚AlodiaCounty) BSoba索巴() feud.Barony {
	return c.索巴Soba
}

func (c *阿洛地亚AlodiaCounty) BWawissi瓦维斯() feud.Barony {
	return c.瓦维斯Wawissi
}

var CAlodia阿洛地亚 AlodiaCounty = &阿洛地亚AlodiaCounty{}

func init() {
	f := CAlodia阿洛地亚.(*阿洛地亚AlodiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1334",
		Title:     "alodia",
		TitleName: "阿洛地亚",
		TitleCode: "c_alodia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿勒瓦Alwa = BAlwa阿勒瓦
	f.阿勒瓦Alwa.SetParent(f)

	f.巴拉Barah = BBarah巴拉
	f.巴拉Barah.SetParent(f)

	f.马拉卡Malaka = BMalaka马拉卡
	f.马拉卡Malaka.SetParent(f)

	f.恩图曼Omdurman = BOmdurman恩图曼
	f.恩图曼Omdurman.SetParent(f)

	f.欧米茄尔可Omjerky = BOmjerky欧米茄尔可
	f.欧米茄尔可Omjerky.SetParent(f)

	f.索巴Soba = BSoba索巴
	f.索巴Soba.SetParent(f)

	f.瓦维斯Wawissi = BWawissi瓦维斯
	f.瓦维斯Wawissi.SetParent(f)

}
