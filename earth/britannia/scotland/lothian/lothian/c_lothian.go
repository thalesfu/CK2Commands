package lothian

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LothianCounty interface {
	feud.County
	BAbercorn阿伯康() feud.Barony
	BEdinburgh爱丁堡() feud.Barony
	BFalkirk福尔柯克() feud.Barony
	BLeith利斯() feud.Barony
	BLinlithgow林利斯戈() feud.Barony
	BStirling斯特灵() feud.Barony
	BTorphichen托菲肯() feud.Barony
}

type 洛锡安LothianCounty struct {
	feud.BaseCounty
	阿伯康Abercorn    feud.Barony
	爱丁堡Edinburgh   feud.Barony
	福尔柯克Falkirk    feud.Barony
	利斯Leith        feud.Barony
	林利斯戈Linlithgow feud.Barony
	斯特灵Stirling    feud.Barony
	托菲肯Torphichen  feud.Barony
}

func (c *洛锡安LothianCounty) BAbercorn阿伯康() feud.Barony {
	return c.阿伯康Abercorn
}

func (c *洛锡安LothianCounty) BEdinburgh爱丁堡() feud.Barony {
	return c.爱丁堡Edinburgh
}

func (c *洛锡安LothianCounty) BFalkirk福尔柯克() feud.Barony {
	return c.福尔柯克Falkirk
}

func (c *洛锡安LothianCounty) BLeith利斯() feud.Barony {
	return c.利斯Leith
}

func (c *洛锡安LothianCounty) BLinlithgow林利斯戈() feud.Barony {
	return c.林利斯戈Linlithgow
}

func (c *洛锡安LothianCounty) BStirling斯特灵() feud.Barony {
	return c.斯特灵Stirling
}

func (c *洛锡安LothianCounty) BTorphichen托菲肯() feud.Barony {
	return c.托菲肯Torphichen
}

var CLothian洛锡安 LothianCounty = &洛锡安LothianCounty{}

func init() {
	f := CLothian洛锡安.(*洛锡安LothianCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "48",
		Title:     "lothian",
		TitleName: "洛锡安",
		TitleCode: "c_lothian",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿伯康Abercorn = BAbercorn阿伯康
	f.阿伯康Abercorn.SetParent(f)

	f.爱丁堡Edinburgh = BEdinburgh爱丁堡
	f.爱丁堡Edinburgh.SetParent(f)

	f.福尔柯克Falkirk = BFalkirk福尔柯克
	f.福尔柯克Falkirk.SetParent(f)

	f.利斯Leith = BLeith利斯
	f.利斯Leith.SetParent(f)

	f.林利斯戈Linlithgow = BLinlithgow林利斯戈
	f.林利斯戈Linlithgow.SetParent(f)

	f.斯特灵Stirling = BStirling斯特灵
	f.斯特灵Stirling.SetParent(f)

	f.托菲肯Torphichen = BTorphichen托菲肯
	f.托菲肯Torphichen.SetParent(f)

}
