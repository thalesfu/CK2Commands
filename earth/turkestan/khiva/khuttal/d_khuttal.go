package khuttal

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khuttal/badakhshan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khuttal/chaghaniyan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khuttal/khuttal"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khiva/khuttal/vakhan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhuttalDuke interface {
	feud.Duke
	CBadakhshan波多叉拏() badakhshan.BadakhshanCounty
	CChaghaniyan斫汗那() chaghaniyan.ChaghaniyanCounty
	CKhuttal珂咄罗() khuttal.KhuttalCounty
	CVakhan镬侃() vakhan.VakhanCounty
}

type 珂咄罗KhuttalDuke struct {
	feud.BaseDuke
	波多叉拏Badakhshan badakhshan.BadakhshanCounty
	斫汗那Chaghaniyan chaghaniyan.ChaghaniyanCounty
	珂咄罗Khuttal     khuttal.KhuttalCounty
	镬侃Vakhan       vakhan.VakhanCounty
}

func (d *珂咄罗KhuttalDuke) CBadakhshan波多叉拏() badakhshan.BadakhshanCounty {
	return d.波多叉拏Badakhshan
}

func (d *珂咄罗KhuttalDuke) CChaghaniyan斫汗那() chaghaniyan.ChaghaniyanCounty {
	return d.斫汗那Chaghaniyan
}

func (d *珂咄罗KhuttalDuke) CKhuttal珂咄罗() khuttal.KhuttalCounty {
	return d.珂咄罗Khuttal
}

func (d *珂咄罗KhuttalDuke) CVakhan镬侃() vakhan.VakhanCounty {
	return d.镬侃Vakhan
}

var DKhuttal珂咄罗 KhuttalDuke = &珂咄罗KhuttalDuke{}

func init() {
	f := DKhuttal珂咄罗.(*珂咄罗KhuttalDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "khuttal",
		TitleName: "珂咄罗",
		TitleCode: "d_khuttal",
		Counties:  map[string]feud.County{},
	}

	f.波多叉拏Badakhshan = badakhshan.CBadakhshan波多叉拏
	f.波多叉拏Badakhshan.SetParent(f)

	f.斫汗那Chaghaniyan = chaghaniyan.CChaghaniyan斫汗那
	f.斫汗那Chaghaniyan.SetParent(f)

	f.珂咄罗Khuttal = khuttal.CKhuttal珂咄罗
	f.珂咄罗Khuttal.SetParent(f)

	f.镬侃Vakhan = vakhan.CVakhan镬侃
	f.镬侃Vakhan.SetParent(f)

}
