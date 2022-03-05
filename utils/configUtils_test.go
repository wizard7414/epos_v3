package utils

import "testing"

func Test_InitConfiguration(t *testing.T) {
	conf := InitConfiguration("../configuration/testConfig.json")

	if conf.Parse.ParseLimit != 0 || conf.Parse.SubFolders != false ||
		conf.Parse.ParseList != "/testPath" || conf.Parse.ParseListUrl != "/testURL" ||
		conf.Parse.ParseTarget != "/testPath" || conf.Db.DbPath != "/testPath" {
		t.Fail()
	}
}
