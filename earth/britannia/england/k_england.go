package england

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/bedford"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/canterbury"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/cumbria"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/gloucester"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hampshire"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hereford"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/lancaster"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/norfolk"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/northumberland"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/somerset"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/york"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EnglandKingdom interface {
	feud.Kingdom
	DBedford埃塞克斯() bedford.BedfordDuke
	DCanterbury肯特() canterbury.CanterburyDuke
	DCumbria坎布里亚() cumbria.CumbriaDuke
	DGloucester赫威赛() gloucester.GloucesterDuke
	DHampshire威塞克斯() hampshire.HampshireDuke
	DHereford麦西亚() hereford.HerefordDuke
	DLancaster兰开斯特() lancaster.LancasterDuke
	DNorfolk东盎格利亚() norfolk.NorfolkDuke
	DNorthumberland诺森伯兰() northumberland.NorthumberlandDuke
	DSomerset萨默塞特() somerset.SomersetDuke
	DYork约克() york.YorkDuke
}

type 英格兰EnglandKingdom struct {
	feud.BaseKingdom
	埃塞克斯Bedford        bedford.BedfordDuke
	肯特Canterbury       canterbury.CanterburyDuke
	坎布里亚Cumbria        cumbria.CumbriaDuke
	赫威赛Gloucester      gloucester.GloucesterDuke
	威塞克斯Hampshire      hampshire.HampshireDuke
	麦西亚Hereford        hereford.HerefordDuke
	兰开斯特Lancaster      lancaster.LancasterDuke
	东盎格利亚Norfolk       norfolk.NorfolkDuke
	诺森伯兰Northumberland northumberland.NorthumberlandDuke
	萨默塞特Somerset       somerset.SomersetDuke
	约克York             york.YorkDuke
}

func (k *英格兰EnglandKingdom) DBedford埃塞克斯() bedford.BedfordDuke {
	return k.埃塞克斯Bedford
}

func (k *英格兰EnglandKingdom) DCanterbury肯特() canterbury.CanterburyDuke {
	return k.肯特Canterbury
}

func (k *英格兰EnglandKingdom) DCumbria坎布里亚() cumbria.CumbriaDuke {
	return k.坎布里亚Cumbria
}

func (k *英格兰EnglandKingdom) DGloucester赫威赛() gloucester.GloucesterDuke {
	return k.赫威赛Gloucester
}

func (k *英格兰EnglandKingdom) DHampshire威塞克斯() hampshire.HampshireDuke {
	return k.威塞克斯Hampshire
}

func (k *英格兰EnglandKingdom) DHereford麦西亚() hereford.HerefordDuke {
	return k.麦西亚Hereford
}

func (k *英格兰EnglandKingdom) DLancaster兰开斯特() lancaster.LancasterDuke {
	return k.兰开斯特Lancaster
}

func (k *英格兰EnglandKingdom) DNorfolk东盎格利亚() norfolk.NorfolkDuke {
	return k.东盎格利亚Norfolk
}

func (k *英格兰EnglandKingdom) DNorthumberland诺森伯兰() northumberland.NorthumberlandDuke {
	return k.诺森伯兰Northumberland
}

func (k *英格兰EnglandKingdom) DSomerset萨默塞特() somerset.SomersetDuke {
	return k.萨默塞特Somerset
}

func (k *英格兰EnglandKingdom) DYork约克() york.YorkDuke {
	return k.约克York
}

var KEngland英格兰 EnglandKingdom = &英格兰EnglandKingdom{}

func init() {
	f := KEngland英格兰.(*英格兰EnglandKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "england",
		TitleName: "英格兰",
		TitleCode: "k_england",
		Dukes:     map[string]feud.Duke{},
	}

	f.埃塞克斯Bedford = bedford.DBedford埃塞克斯
	f.埃塞克斯Bedford.SetParent(f)

	f.肯特Canterbury = canterbury.DCanterbury肯特
	f.肯特Canterbury.SetParent(f)

	f.坎布里亚Cumbria = cumbria.DCumbria坎布里亚
	f.坎布里亚Cumbria.SetParent(f)

	f.赫威赛Gloucester = gloucester.DGloucester赫威赛
	f.赫威赛Gloucester.SetParent(f)

	f.威塞克斯Hampshire = hampshire.DHampshire威塞克斯
	f.威塞克斯Hampshire.SetParent(f)

	f.麦西亚Hereford = hereford.DHereford麦西亚
	f.麦西亚Hereford.SetParent(f)

	f.兰开斯特Lancaster = lancaster.DLancaster兰开斯特
	f.兰开斯特Lancaster.SetParent(f)

	f.东盎格利亚Norfolk = norfolk.DNorfolk东盎格利亚
	f.东盎格利亚Norfolk.SetParent(f)

	f.诺森伯兰Northumberland = northumberland.DNorthumberland诺森伯兰
	f.诺森伯兰Northumberland.SetParent(f)

	f.萨默塞特Somerset = somerset.DSomerset萨默塞特
	f.萨默塞特Somerset.SetParent(f)

	f.约克York = york.DYork约克
	f.约克York.SetParent(f)

}
