package anupa

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/anupa/asirgarh"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/anupa/burhanpur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/anupa/mandapika"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/anupa/thalner"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnupaDuke interface {
	feud.Duke
	CAsirgarh阿斯梨姞利呬() asirgarh.AsirgarhCounty
	CBurhanpur补罗汉补罗() burhanpur.BurhanpurCounty
	CMandapika曼荼卑迦() mandapika.MandapikaCounty
	CThalner闼罗内() thalner.ThalnerCounty
}

type 阿怒波AnupaDuke struct {
	feud.BaseDuke
	阿斯梨姞利呬Asirgarh asirgarh.AsirgarhCounty
	补罗汉补罗Burhanpur burhanpur.BurhanpurCounty
	曼荼卑迦Mandapika  mandapika.MandapikaCounty
	闼罗内Thalner     thalner.ThalnerCounty
}

func (d *阿怒波AnupaDuke) CAsirgarh阿斯梨姞利呬() asirgarh.AsirgarhCounty {
	return d.阿斯梨姞利呬Asirgarh
}

func (d *阿怒波AnupaDuke) CBurhanpur补罗汉补罗() burhanpur.BurhanpurCounty {
	return d.补罗汉补罗Burhanpur
}

func (d *阿怒波AnupaDuke) CMandapika曼荼卑迦() mandapika.MandapikaCounty {
	return d.曼荼卑迦Mandapika
}

func (d *阿怒波AnupaDuke) CThalner闼罗内() thalner.ThalnerCounty {
	return d.闼罗内Thalner
}

var DAnupa阿怒波 AnupaDuke = &阿怒波AnupaDuke{}

func init() {
	f := DAnupa阿怒波.(*阿怒波AnupaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "anupa",
		TitleName: "阿怒波",
		TitleCode: "d_anupa",
		Counties:  map[string]feud.County{},
	}

	f.阿斯梨姞利呬Asirgarh = asirgarh.CAsirgarh阿斯梨姞利呬
	f.阿斯梨姞利呬Asirgarh.SetParent(f)

	f.补罗汉补罗Burhanpur = burhanpur.CBurhanpur补罗汉补罗
	f.补罗汉补罗Burhanpur.SetParent(f)

	f.曼荼卑迦Mandapika = mandapika.CMandapika曼荼卑迦
	f.曼荼卑迦Mandapika.SetParent(f)

	f.闼罗内Thalner = thalner.CThalner闼罗内
	f.闼罗内Thalner.SetParent(f)

}
