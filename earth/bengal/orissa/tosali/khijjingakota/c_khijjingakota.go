package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhijjingakotaCounty interface {
	feud.County
	BAsanapat阿桑纳帕提() feud.Barony
	BBahalda伯赫尔达() feud.Barony
	BBaripada巴里帕达() feud.Barony
	BGhatagaon揭吒伽罗摩() feud.Barony
	BKhijjinga契承伽() feud.Barony
	BKirianagar吉里安纳加尔() feud.Barony
	BSitabhinji斯塔宾基() feud.Barony
}

type 契承伽拘吒KhijjingakotaCounty struct {
	feud.BaseCounty
	阿桑纳帕提Asanapat    feud.Barony
	伯赫尔达Bahalda      feud.Barony
	巴里帕达Baripada     feud.Barony
	揭吒伽罗摩Ghatagaon   feud.Barony
	契承伽Khijjinga     feud.Barony
	吉里安纳加尔Kirianagar feud.Barony
	斯塔宾基Sitabhinji   feud.Barony
}

func (c *契承伽拘吒KhijjingakotaCounty) BAsanapat阿桑纳帕提() feud.Barony {
	return c.阿桑纳帕提Asanapat
}

func (c *契承伽拘吒KhijjingakotaCounty) BBahalda伯赫尔达() feud.Barony {
	return c.伯赫尔达Bahalda
}

func (c *契承伽拘吒KhijjingakotaCounty) BBaripada巴里帕达() feud.Barony {
	return c.巴里帕达Baripada
}

func (c *契承伽拘吒KhijjingakotaCounty) BGhatagaon揭吒伽罗摩() feud.Barony {
	return c.揭吒伽罗摩Ghatagaon
}

func (c *契承伽拘吒KhijjingakotaCounty) BKhijjinga契承伽() feud.Barony {
	return c.契承伽Khijjinga
}

func (c *契承伽拘吒KhijjingakotaCounty) BKirianagar吉里安纳加尔() feud.Barony {
	return c.吉里安纳加尔Kirianagar
}

func (c *契承伽拘吒KhijjingakotaCounty) BSitabhinji斯塔宾基() feud.Barony {
	return c.斯塔宾基Sitabhinji
}

var CKhijjingakota契承伽拘吒 KhijjingakotaCounty = &契承伽拘吒KhijjingakotaCounty{}

func init() {
	f := CKhijjingakota契承伽拘吒.(*契承伽拘吒KhijjingakotaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1247",
		Title:     "khijjingakota",
		TitleName: "契承伽拘吒",
		TitleCode: "c_khijjingakota",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿桑纳帕提Asanapat = BAsanapat阿桑纳帕提
	f.阿桑纳帕提Asanapat.SetParent(f)

	f.伯赫尔达Bahalda = BBahalda伯赫尔达
	f.伯赫尔达Bahalda.SetParent(f)

	f.巴里帕达Baripada = BBaripada巴里帕达
	f.巴里帕达Baripada.SetParent(f)

	f.揭吒伽罗摩Ghatagaon = BGhatagaon揭吒伽罗摩
	f.揭吒伽罗摩Ghatagaon.SetParent(f)

	f.契承伽Khijjinga = BKhijjinga契承伽
	f.契承伽Khijjinga.SetParent(f)

	f.吉里安纳加尔Kirianagar = BKirianagar吉里安纳加尔
	f.吉里安纳加尔Kirianagar.SetParent(f)

	f.斯塔宾基Sitabhinji = BSitabhinji斯塔宾基
	f.斯塔宾基Sitabhinji.SetParent(f)

}
