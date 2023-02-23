package mali

import (
	"github.com/thalesfu/CK2Commands/earth/mali/mali/mali/bure"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/mali/dodugu"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/mali/kiri"
	"github.com/thalesfu/CK2Commands/earth/mali/mali/mali/mali"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaliDuke interface {
	feud.Duke
	CBure布尔() bure.BureCounty
	CDodugu多杜古() dodugu.DoduguCounty
	CKiri基里() kiri.KiriCounty
	CMali尼亚尼() mali.MaliCounty
}

type 曼丁MaliDuke struct {
	feud.BaseDuke
	布尔Bure    bure.BureCounty
	多杜古Dodugu dodugu.DoduguCounty
	基里Kiri    kiri.KiriCounty
	尼亚尼Mali   mali.MaliCounty
}

func (d *曼丁MaliDuke) CBure布尔() bure.BureCounty {
	return d.布尔Bure
}

func (d *曼丁MaliDuke) CDodugu多杜古() dodugu.DoduguCounty {
	return d.多杜古Dodugu
}

func (d *曼丁MaliDuke) CKiri基里() kiri.KiriCounty {
	return d.基里Kiri
}

func (d *曼丁MaliDuke) CMali尼亚尼() mali.MaliCounty {
	return d.尼亚尼Mali
}

var DMali曼丁 MaliDuke = &曼丁MaliDuke{}

func init() {
	f := DMali曼丁.(*曼丁MaliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mali",
		TitleName: "曼丁",
		TitleCode: "d_mali",
		Counties:  map[string]feud.County{},
	}

	f.布尔Bure = bure.CBure布尔
	f.布尔Bure.SetParent(f)

	f.多杜古Dodugu = dodugu.CDodugu多杜古
	f.多杜古Dodugu.SetParent(f)

	f.基里Kiri = kiri.CKiri基里
	f.基里Kiri.SetParent(f)

	f.尼亚尼Mali = mali.CMali尼亚尼
	f.尼亚尼Mali.SetParent(f)

}
