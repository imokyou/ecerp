package ecerp

import (
	"testing"
)

func TestGenerateSign(t *testing.T) {
	tests := []struct {
		name      string
		params    map[string]string
		appSecret string
		wantLen   int // MD5 hash 长度固定为 32
	}{
		{
			name: "基本签名测试",
			params: map[string]string{
				"app_key":          "test_app_key",
				"service_id":       "test_service",
				"interface_method": "getWarehouse",
				"timestamp":        "1609459200000",
				"nonce_str":        "abc123",
				"charset":          "UTF-8",
				"version":          "V1.0.0",
				"sign_type":        "MD5",
				"biz_content":      "{}",
			},
			appSecret: "test_secret",
			wantLen:   32,
		},
		{
			name: "过滤空值参数",
			params: map[string]string{
				"app_key":          "test_app_key",
				"service_id":       "",
				"interface_method": "getWarehouse",
				"timestamp":        "1609459200000",
				"nonce_str":        "",
				"charset":          "UTF-8",
				"version":          "V1.0.0",
				"sign_type":        "MD5",
				"biz_content":      "{}",
			},
			appSecret: "test_secret",
			wantLen:   32,
		},
		{
			name: "过滤sign参数",
			params: map[string]string{
				"app_key":          "test_app_key",
				"interface_method": "getWarehouse",
				"timestamp":        "1609459200000",
				"sign":             "should_be_excluded",
				"charset":          "UTF-8",
				"version":          "V1.0.0",
				"sign_type":        "MD5",
				"biz_content":      "{}",
			},
			appSecret: "test_secret",
			wantLen:   32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSign(tt.params, tt.appSecret)
			if len(got) != tt.wantLen {
				t.Errorf("GenerateSign() = %v (len=%d), want len=%d", got, len(got), tt.wantLen)
			}
		})
	}
}

func TestGenerateSignDeterministic(t *testing.T) {
	params := map[string]string{
		"app_key":          "mykey",
		"interface_method": "getOrderList",
		"timestamp":        "1609459200000",
		"nonce_str":        "random123",
		"charset":          "UTF-8",
		"version":          "V1.0.0",
		"sign_type":        "MD5",
		"biz_content":      `{"page":1,"page_size":10}`,
	}
	secret := "mysecret"

	sign1 := GenerateSign(params, secret)
	sign2 := GenerateSign(params, secret)

	if sign1 != sign2 {
		t.Errorf("签名不确定: %s != %s", sign1, sign2)
	}
}

func TestGenerateSignExcludesEmptyAndSign(t *testing.T) {
	params1 := map[string]string{
		"app_key":   "key1",
		"timestamp": "12345",
	}
	params2 := map[string]string{
		"app_key":   "key1",
		"timestamp": "12345",
		"sign":      "old_sign",
		"empty_val": "",
	}

	secret := "secret"
	sign1 := GenerateSign(params1, secret)
	sign2 := GenerateSign(params2, secret)

	if sign1 != sign2 {
		t.Errorf("sign 和空值应被过滤: %s != %s", sign1, sign2)
	}
}
