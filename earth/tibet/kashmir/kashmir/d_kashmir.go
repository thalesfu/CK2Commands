package kashmir

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/kashmir/gilgit"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/kashmir/kasmira"
	"github.com/thalesfu/CK2Commands/earth/tibet/kashmir/kashmir/skardu"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KashmirDuke interface {
	feud.Duke
	CGilgit吉尔吉特() gilgit.GilgitCounty
	CKasmira迦湿弥罗() kasmira.KasmiraCounty
	CSkardu锡卡都() skardu.SkarduCounty
}

type 迦湿弥罗KashmirDuke struct {
	feud.BaseDuke
	吉尔吉特Gilgit  gilgit.GilgitCounty
	迦湿弥罗Kasmira kasmira.KasmiraCounty
	锡卡都Skardu   skardu.SkarduCounty
}

func (d *迦湿弥罗KashmirDuke) CGilgit吉尔吉特() gilgit.GilgitCounty {
	return d.吉尔吉特Gilgit
}

func (d *迦湿弥罗KashmirDuke) CKasmira迦湿弥罗() kasmira.KasmiraCounty {
	return d.迦湿弥罗Kasmira
}

func (d *迦湿弥罗KashmirDuke) CSkardu锡卡都() skardu.SkarduCounty {
	return d.锡卡都Skardu
}

var DKashmir迦湿弥罗 KashmirDuke = &迦湿弥罗KashmirDuke{}

func init() {
	f := DKashmir迦湿弥罗.(*迦湿弥罗KashmirDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kashmir",
		TitleName: "迦湿弥罗",
		TitleCode: "d_kashmir",
		Counties:  map[string]feud.County{},
	}

	f.吉尔吉特Gilgit = gilgit.CGilgit吉尔吉特
	f.吉尔吉特Gilgit.SetParent(f)

	f.迦湿弥罗Kasmira = kasmira.CKasmira迦湿弥罗
	f.迦湿弥罗Kasmira.SetParent(f)

	f.锡卡都Skardu = skardu.CSkardu锡卡都
	f.锡卡都Skardu.SetParent(f)

}
