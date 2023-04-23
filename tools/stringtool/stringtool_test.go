package stringtool

import "testing"

func TestChineseNumStringToArabicNum(t *testing.T) {
	testConfig := []struct {
		input    string
		expected int64
	}{
		{"十一亿二百三十四万五千六百七十八", 1102345678},
		{"十六", 16},
		{"三千零八十四", 3084},
		{"三亿四千五百六十七万八千九百", 345678900},
		{"一亿五百万", 105000000},
		{"三百零五", 305},
		{"两万", 20000},
	}
	for _, s := range testConfig {
		out := ChineseNumStringToArabicNum(s.input)
		if out != s.expected {
			t.Errorf("Translate1(%s) != %d, out=%d", s.input, s.expected, out)
		}
	}
}

func TestGenerateUUID(t *testing.T) {
	out := GenerateUUID()
	if len(out) != 36 {
		t.Errorf("GenerateUUID() != 36, out=%d", len(out))
	}
	t.Logf("GenerateUUID() = %s", out)
}
