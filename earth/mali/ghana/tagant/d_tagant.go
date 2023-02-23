package tagant

import (
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/tagant/aoudaghost"
	"github.com/thalesfu/CK2Commands/earth/mali/ghana/tagant/tagant"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TagantDuke interface {
	feud.Duke
	CAoudaghost奥达戈斯特() aoudaghost.AoudaghostCounty
	CTagant塔甘特() tagant.TagantCounty
}

type 塔甘特TagantDuke struct {
	feud.BaseDuke
	奥达戈斯特Aoudaghost aoudaghost.AoudaghostCounty
	塔甘特Tagant       tagant.TagantCounty
}

func (d *塔甘特TagantDuke) CAoudaghost奥达戈斯特() aoudaghost.AoudaghostCounty {
	return d.奥达戈斯特Aoudaghost
}

func (d *塔甘特TagantDuke) CTagant塔甘特() tagant.TagantCounty {
	return d.塔甘特Tagant
}

var DTagant塔甘特 TagantDuke = &塔甘特TagantDuke{}

func init() {
	f := DTagant塔甘特.(*塔甘特TagantDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tagant",
		TitleName: "塔甘特",
		TitleCode: "d_tagant",
		Counties:  map[string]feud.County{},
	}

	f.奥达戈斯特Aoudaghost = aoudaghost.CAoudaghost奥达戈斯特
	f.奥达戈斯特Aoudaghost.SetParent(f)

	f.塔甘特Tagant = tagant.CTagant塔甘特
	f.塔甘特Tagant.SetParent(f)

}
