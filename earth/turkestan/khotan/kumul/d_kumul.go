package kumul

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/charkliq"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/kumtag"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/kumul"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/lopnor"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/loulan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/khotan/kumul/yuni"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KumulDuke interface {
	feud.Duke
	CCharkliq卡克里克() charkliq.CharkliqCounty
	CKumtag库姆塔格() kumtag.KumtagCounty
	CKumul哈密() kumul.KumulCounty
	CLopnor罗布泊() lopnor.LopnorCounty
	CLoulan楼兰() loulan.LoulanCounty
	CYuni扜泥() yuni.YuniCounty
}

type 哈密KumulDuke struct {
	feud.BaseDuke
	卡克里克Charkliq charkliq.CharkliqCounty
	库姆塔格Kumtag   kumtag.KumtagCounty
	哈密Kumul      kumul.KumulCounty
	罗布泊Lopnor    lopnor.LopnorCounty
	楼兰Loulan     loulan.LoulanCounty
	扜泥Yuni       yuni.YuniCounty
}

func (d *哈密KumulDuke) CCharkliq卡克里克() charkliq.CharkliqCounty {
	return d.卡克里克Charkliq
}

func (d *哈密KumulDuke) CKumtag库姆塔格() kumtag.KumtagCounty {
	return d.库姆塔格Kumtag
}

func (d *哈密KumulDuke) CKumul哈密() kumul.KumulCounty {
	return d.哈密Kumul
}

func (d *哈密KumulDuke) CLopnor罗布泊() lopnor.LopnorCounty {
	return d.罗布泊Lopnor
}

func (d *哈密KumulDuke) CLoulan楼兰() loulan.LoulanCounty {
	return d.楼兰Loulan
}

func (d *哈密KumulDuke) CYuni扜泥() yuni.YuniCounty {
	return d.扜泥Yuni
}

var DKumul哈密 KumulDuke = &哈密KumulDuke{}

func init() {
	f := DKumul哈密.(*哈密KumulDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kumul",
		TitleName: "哈密",
		TitleCode: "d_kumul",
		Counties:  map[string]feud.County{},
	}

	f.卡克里克Charkliq = charkliq.CCharkliq卡克里克
	f.卡克里克Charkliq.SetParent(f)

	f.库姆塔格Kumtag = kumtag.CKumtag库姆塔格
	f.库姆塔格Kumtag.SetParent(f)

	f.哈密Kumul = kumul.CKumul哈密
	f.哈密Kumul.SetParent(f)

	f.罗布泊Lopnor = lopnor.CLopnor罗布泊
	f.罗布泊Lopnor.SetParent(f)

	f.楼兰Loulan = loulan.CLoulan楼兰
	f.楼兰Loulan.SetParent(f)

	f.扜泥Yuni = yuni.CYuni扜泥
	f.扜泥Yuni.SetParent(f)

}
