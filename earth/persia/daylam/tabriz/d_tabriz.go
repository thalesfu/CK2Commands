package tabriz

import (
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/tabriz/maragha"
	"github.com/thalesfu/CK2Commands/earth/persia/daylam/tabriz/tabriz"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TabrizDuke interface {
	feud.Duke
	CMaragha蔑剌哈() maragha.MaraghaCounty
	CTabriz大不里士() tabriz.TabrizCounty
}

type 大不里士TabrizDuke struct {
	feud.BaseDuke
	蔑剌哈Maragha maragha.MaraghaCounty
	大不里士Tabriz tabriz.TabrizCounty
}

func (d *大不里士TabrizDuke) CMaragha蔑剌哈() maragha.MaraghaCounty {
	return d.蔑剌哈Maragha
}

func (d *大不里士TabrizDuke) CTabriz大不里士() tabriz.TabrizCounty {
	return d.大不里士Tabriz
}

var DTabriz大不里士 TabrizDuke = &大不里士TabrizDuke{}

func init() {
	f := DTabriz大不里士.(*大不里士TabrizDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tabriz",
		TitleName: "大不里士",
		TitleCode: "d_tabriz",
		Counties:  map[string]feud.County{},
	}

	f.蔑剌哈Maragha = maragha.CMaragha蔑剌哈
	f.蔑剌哈Maragha.SetParent(f)

	f.大不里士Tabriz = tabriz.CTabriz大不里士
	f.大不里士Tabriz.SetParent(f)

}
