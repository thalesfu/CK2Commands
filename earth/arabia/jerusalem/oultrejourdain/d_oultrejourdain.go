package oultrejourdain

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/oultrejourdain/kerak"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/oultrejourdain/madaba"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/oultrejourdain/monreal"
	"github.com/thalesfu/CK2Commands/earth/arabia/jerusalem/oultrejourdain/negev"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OultrejourdainDuke interface {
	feud.Duke
	CKerak克拉克() kerak.KerakCounty
	CMadaba马德巴() madaba.MadabaCounty
	CMonreal蒙雷阿尔() monreal.MonrealCounty
	CNegev内盖夫() negev.NegevCounty
}

type 外约旦OultrejourdainDuke struct {
	feud.BaseDuke
	克拉克Kerak    kerak.KerakCounty
	马德巴Madaba   madaba.MadabaCounty
	蒙雷阿尔Monreal monreal.MonrealCounty
	内盖夫Negev    negev.NegevCounty
}

func (d *外约旦OultrejourdainDuke) CKerak克拉克() kerak.KerakCounty {
	return d.克拉克Kerak
}

func (d *外约旦OultrejourdainDuke) CMadaba马德巴() madaba.MadabaCounty {
	return d.马德巴Madaba
}

func (d *外约旦OultrejourdainDuke) CMonreal蒙雷阿尔() monreal.MonrealCounty {
	return d.蒙雷阿尔Monreal
}

func (d *外约旦OultrejourdainDuke) CNegev内盖夫() negev.NegevCounty {
	return d.内盖夫Negev
}

var DOultrejourdain外约旦 OultrejourdainDuke = &外约旦OultrejourdainDuke{}

func init() {
	f := DOultrejourdain外约旦.(*外约旦OultrejourdainDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "oultrejourdain",
		TitleName: "外约旦",
		TitleCode: "d_oultrejourdain",
		Counties:  map[string]feud.County{},
	}

	f.克拉克Kerak = kerak.CKerak克拉克
	f.克拉克Kerak.SetParent(f)

	f.马德巴Madaba = madaba.CMadaba马德巴
	f.马德巴Madaba.SetParent(f)

	f.蒙雷阿尔Monreal = monreal.CMonreal蒙雷阿尔
	f.蒙雷阿尔Monreal.SetParent(f)

	f.内盖夫Negev = negev.CNegev内盖夫
	f.内盖夫Negev.SetParent(f)

}
