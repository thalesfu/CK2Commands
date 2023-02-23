package hereford

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hereford/leicester"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hereford/lincoln"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hereford/warwick"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/hereford/worcester"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HerefordDuke interface {
	feud.Duke
	CLeicester莱斯特() leicester.LeicesterCounty
	CLincoln林肯() lincoln.LincolnCounty
	CWarwick沃里克() warwick.WarwickCounty
	CWorcester伍斯特() worcester.WorcesterCounty
}

type 麦西亚HerefordDuke struct {
	feud.BaseDuke
	莱斯特Leicester leicester.LeicesterCounty
	林肯Lincoln    lincoln.LincolnCounty
	沃里克Warwick   warwick.WarwickCounty
	伍斯特Worcester worcester.WorcesterCounty
}

func (d *麦西亚HerefordDuke) CLeicester莱斯特() leicester.LeicesterCounty {
	return d.莱斯特Leicester
}

func (d *麦西亚HerefordDuke) CLincoln林肯() lincoln.LincolnCounty {
	return d.林肯Lincoln
}

func (d *麦西亚HerefordDuke) CWarwick沃里克() warwick.WarwickCounty {
	return d.沃里克Warwick
}

func (d *麦西亚HerefordDuke) CWorcester伍斯特() worcester.WorcesterCounty {
	return d.伍斯特Worcester
}

var DHereford麦西亚 HerefordDuke = &麦西亚HerefordDuke{}

func init() {
	f := DHereford麦西亚.(*麦西亚HerefordDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "hereford",
		TitleName: "麦西亚",
		TitleCode: "d_hereford",
		Counties:  map[string]feud.County{},
	}

	f.莱斯特Leicester = leicester.CLeicester莱斯特
	f.莱斯特Leicester.SetParent(f)

	f.林肯Lincoln = lincoln.CLincoln林肯
	f.林肯Lincoln.SetParent(f)

	f.沃里克Warwick = warwick.CWarwick沃里克
	f.沃里克Warwick.SetParent(f)

	f.伍斯特Worcester = worcester.CWorcester伍斯特
	f.伍斯特Worcester.SetParent(f)

}
