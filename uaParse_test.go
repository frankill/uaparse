package ua

import (
	"testing"
)

const (
	tests = "Mozilla/5.0 (Linux; Android 7.1.2; vivo X9 Build/N2G47H; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/66.0.3359.126 MQQBrowser/6.2 TBS/044506 Mobile Safari/537.36 MMWEBID/134 MicroMessenger/7.0.3.1400(0x2700033C) Process/tools NetType/4G Language/zh_CN"
)

func TestParseDevice(t *testing.T) {

	if ParseDevice(tests).Brand != "vivo" {
		t.Errorf("品牌解析有问题, 解析结果 %s", ParseDevice(tests).Brand)
	}

	if ParseDevice(tests).Family != "vivo X9" {
		t.Errorf("设备解析有问题, 解析结果 %s", ParseDevice(tests).Family)
	}

	if ParseDevice(tests).Model != "X9" {
		t.Errorf("型号解析有问题, 解析结果 %s", ParseDevice(tests).Model)
	}

}

func TestParseOS(t *testing.T) {

	if ParseOS(tests).Family != "Android" {
		t.Errorf("解析有问题, 解析结果 %s", ParseOS(tests).Family)
	}

	if ParseOS(tests).Major != "7" {
		t.Errorf("解析有问题, 解析结果 %s", ParseOS(tests).Major)
	}

	if ParseOS(tests).Minor != "1" {
		t.Errorf("解析有问题, 解析结果 %s", ParseOS(tests).Minor)
	}

	if ParseOS(tests).Patch != "2" {
		t.Errorf("解析有问题, 解析结果 %s", ParseOS(tests).Patch)
	}

	if ParseOS(tests).PatchMinor != "" {
		t.Errorf("解析有问题, 解析结果 %s", ParseOS(tests).PatchMinor)
	}

}

func TestParseUA(t *testing.T) {

	if ParseUA(tests).Family != "Safari" {
		t.Errorf("解析有问题, 解析结果 %s", ParseUA(tests).Family)
	}

	if ParseUA(tests).Major != "4" {
		t.Errorf("解析有问题, 解析结果 %s", ParseUA(tests).Major)
	}

	if ParseUA(tests).Minor != "0" {
		t.Errorf("解析有问题, 解析结果 %s", ParseUA(tests).Minor)
	}

	if ParseUA(tests).Patch != "" {
		t.Errorf("解析有问题, 解析结果 %s", ParseUA(tests).Patch)
	}
}
