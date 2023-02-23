package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarurCounty interface {
	feud.County
	BAskhanda阿斯坎达() feud.Barony
	BChishtian支室帝延() feud.Barony
	BKarur伽卢罗() feud.Barony
	BMarot默罗德() feud.Barony
	BNyahalode若诃卢提() feud.Barony
	BOdalgaon乌陀罗伽罗摩() feud.Barony
	BVehari维哈里() feud.Barony
}

type 伽卢罗KarurCounty struct {
	feud.BaseCounty
	阿斯坎达Askhanda   feud.Barony
	支室帝延Chishtian  feud.Barony
	伽卢罗Karur       feud.Barony
	默罗德Marot       feud.Barony
	若诃卢提Nyahalode  feud.Barony
	乌陀罗伽罗摩Odalgaon feud.Barony
	维哈里Vehari      feud.Barony
}

func (c *伽卢罗KarurCounty) BAskhanda阿斯坎达() feud.Barony {
	return c.阿斯坎达Askhanda
}

func (c *伽卢罗KarurCounty) BChishtian支室帝延() feud.Barony {
	return c.支室帝延Chishtian
}

func (c *伽卢罗KarurCounty) BKarur伽卢罗() feud.Barony {
	return c.伽卢罗Karur
}

func (c *伽卢罗KarurCounty) BMarot默罗德() feud.Barony {
	return c.默罗德Marot
}

func (c *伽卢罗KarurCounty) BNyahalode若诃卢提() feud.Barony {
	return c.若诃卢提Nyahalode
}

func (c *伽卢罗KarurCounty) BOdalgaon乌陀罗伽罗摩() feud.Barony {
	return c.乌陀罗伽罗摩Odalgaon
}

func (c *伽卢罗KarurCounty) BVehari维哈里() feud.Barony {
	return c.维哈里Vehari
}

var CKarur伽卢罗 KarurCounty = &伽卢罗KarurCounty{}

func init() {
	f := CKarur伽卢罗.(*伽卢罗KarurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1178",
		Title:     "karur",
		TitleName: "伽卢罗",
		TitleCode: "c_karur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯坎达Askhanda = BAskhanda阿斯坎达
	f.阿斯坎达Askhanda.SetParent(f)

	f.支室帝延Chishtian = BChishtian支室帝延
	f.支室帝延Chishtian.SetParent(f)

	f.伽卢罗Karur = BKarur伽卢罗
	f.伽卢罗Karur.SetParent(f)

	f.默罗德Marot = BMarot默罗德
	f.默罗德Marot.SetParent(f)

	f.若诃卢提Nyahalode = BNyahalode若诃卢提
	f.若诃卢提Nyahalode.SetParent(f)

	f.乌陀罗伽罗摩Odalgaon = BOdalgaon乌陀罗伽罗摩
	f.乌陀罗伽罗摩Odalgaon.SetParent(f)

	f.维哈里Vehari = BVehari维哈里
	f.维哈里Vehari.SetParent(f)

}
