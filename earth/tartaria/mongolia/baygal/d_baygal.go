package baygal

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal/baygal"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal/chikoi"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal/khilok"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal/selenge"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia/baygal/uda"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BaygalDuke interface {
	feud.Duke
	CBaygal贝加尔() baygal.BaygalCounty
	CChikoi赤苦() chikoi.ChikoiCounty
	CKhilok勤勒豁() khilok.KhilokCounty
	CSelenge娑陵() selenge.SelengeCounty
	CUda兀的() uda.UdaCounty
}

type 贝加尔BaygalDuke struct {
	feud.BaseDuke
	贝加尔Baygal baygal.BaygalCounty
	赤苦Chikoi  chikoi.ChikoiCounty
	勤勒豁Khilok khilok.KhilokCounty
	娑陵Selenge selenge.SelengeCounty
	兀的Uda     uda.UdaCounty
}

func (d *贝加尔BaygalDuke) CBaygal贝加尔() baygal.BaygalCounty {
	return d.贝加尔Baygal
}

func (d *贝加尔BaygalDuke) CChikoi赤苦() chikoi.ChikoiCounty {
	return d.赤苦Chikoi
}

func (d *贝加尔BaygalDuke) CKhilok勤勒豁() khilok.KhilokCounty {
	return d.勤勒豁Khilok
}

func (d *贝加尔BaygalDuke) CSelenge娑陵() selenge.SelengeCounty {
	return d.娑陵Selenge
}

func (d *贝加尔BaygalDuke) CUda兀的() uda.UdaCounty {
	return d.兀的Uda
}

var DBaygal贝加尔 BaygalDuke = &贝加尔BaygalDuke{}

func init() {
	f := DBaygal贝加尔.(*贝加尔BaygalDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "baygal",
		TitleName: "贝加尔",
		TitleCode: "d_baygal",
		Counties:  map[string]feud.County{},
	}

	f.贝加尔Baygal = baygal.CBaygal贝加尔
	f.贝加尔Baygal.SetParent(f)

	f.赤苦Chikoi = chikoi.CChikoi赤苦
	f.赤苦Chikoi.SetParent(f)

	f.勤勒豁Khilok = khilok.CKhilok勤勒豁
	f.勤勒豁Khilok.SetParent(f)

	f.娑陵Selenge = selenge.CSelenge娑陵
	f.娑陵Selenge.SetParent(f)

	f.兀的Uda = uda.CUda兀的
	f.兀的Uda.SetParent(f)

}
