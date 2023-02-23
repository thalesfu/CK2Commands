package tripoli

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/tripoli/baalbek"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/tripoli/tortosa"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/tripoli/tripoli"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TripoliDuke interface {
	feud.Duke
	CBaalbek巴勒贝克() baalbek.BaalbekCounty
	CTortosa托尔托萨() tortosa.TortosaCounty
	CTripoli的黎波里() tripoli.TripoliCounty
}

type 的黎波里TripoliDuke struct {
	feud.BaseDuke
	巴勒贝克Baalbek baalbek.BaalbekCounty
	托尔托萨Tortosa tortosa.TortosaCounty
	的黎波里Tripoli tripoli.TripoliCounty
}

func (d *的黎波里TripoliDuke) CBaalbek巴勒贝克() baalbek.BaalbekCounty {
	return d.巴勒贝克Baalbek
}

func (d *的黎波里TripoliDuke) CTortosa托尔托萨() tortosa.TortosaCounty {
	return d.托尔托萨Tortosa
}

func (d *的黎波里TripoliDuke) CTripoli的黎波里() tripoli.TripoliCounty {
	return d.的黎波里Tripoli
}

var DTripoli的黎波里 TripoliDuke = &的黎波里TripoliDuke{}

func init() {
	f := DTripoli的黎波里.(*的黎波里TripoliDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "tripoli",
		TitleName: "的黎波里",
		TitleCode: "d_tripoli",
		Counties:  map[string]feud.County{},
	}

	f.巴勒贝克Baalbek = baalbek.CBaalbek巴勒贝克
	f.巴勒贝克Baalbek.SetParent(f)

	f.托尔托萨Tortosa = tortosa.CTortosa托尔托萨
	f.托尔托萨Tortosa.SetParent(f)

	f.的黎波里Tripoli = tripoli.CTripoli的黎波里
	f.的黎波里Tripoli.SetParent(f)

}
