package bhakkar

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/aror"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/bhakkar"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/kandail"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/quzdar"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/rajanpur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar/sibi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BhakkarDuke interface {
	feud.Duke
	CAror阿卢梨() aror.ArorCounty
	CBhakkar珀格尔() bhakkar.BhakkarCounty
	CKandail甘代尔() kandail.KandailCounty
	CQuzdar古兹达尔() quzdar.QuzdarCounty
	CRajanpur罗旬补罗() rajanpur.RajanpurCounty
	CSibi尸毗() sibi.SibiCounty
}

type 珀格尔BhakkarDuke struct {
	feud.BaseDuke
	阿卢梨Aror      aror.ArorCounty
	珀格尔Bhakkar   bhakkar.BhakkarCounty
	甘代尔Kandail   kandail.KandailCounty
	古兹达尔Quzdar   quzdar.QuzdarCounty
	罗旬补罗Rajanpur rajanpur.RajanpurCounty
	尸毗Sibi       sibi.SibiCounty
}

func (d *珀格尔BhakkarDuke) CAror阿卢梨() aror.ArorCounty {
	return d.阿卢梨Aror
}

func (d *珀格尔BhakkarDuke) CBhakkar珀格尔() bhakkar.BhakkarCounty {
	return d.珀格尔Bhakkar
}

func (d *珀格尔BhakkarDuke) CKandail甘代尔() kandail.KandailCounty {
	return d.甘代尔Kandail
}

func (d *珀格尔BhakkarDuke) CQuzdar古兹达尔() quzdar.QuzdarCounty {
	return d.古兹达尔Quzdar
}

func (d *珀格尔BhakkarDuke) CRajanpur罗旬补罗() rajanpur.RajanpurCounty {
	return d.罗旬补罗Rajanpur
}

func (d *珀格尔BhakkarDuke) CSibi尸毗() sibi.SibiCounty {
	return d.尸毗Sibi
}

var DBhakkar珀格尔 BhakkarDuke = &珀格尔BhakkarDuke{}

func init() {
	f := DBhakkar珀格尔.(*珀格尔BhakkarDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "bhakkar",
		TitleName: "珀格尔",
		TitleCode: "d_bhakkar",
		Counties:  map[string]feud.County{},
	}

	f.阿卢梨Aror = aror.CAror阿卢梨
	f.阿卢梨Aror.SetParent(f)

	f.珀格尔Bhakkar = bhakkar.CBhakkar珀格尔
	f.珀格尔Bhakkar.SetParent(f)

	f.甘代尔Kandail = kandail.CKandail甘代尔
	f.甘代尔Kandail.SetParent(f)

	f.古兹达尔Quzdar = quzdar.CQuzdar古兹达尔
	f.古兹达尔Quzdar.SetParent(f)

	f.罗旬补罗Rajanpur = rajanpur.CRajanpur罗旬补罗
	f.罗旬补罗Rajanpur.SetParent(f)

	f.尸毗Sibi = sibi.CSibi尸毗
	f.尸毗Sibi.SetParent(f)

}
