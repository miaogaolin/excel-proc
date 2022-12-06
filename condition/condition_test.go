package condition

import "testing"

func TestValidate(t *testing.T) {
	type args struct {
		data          map[string]interface{}
		conditionExpr string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"==", args{map[string]interface{}{"col1": 1},
			`{col1} == 1}`}, true, false},
		{"!=", args{map[string]interface{}{"col2": "aaa"},
			`{col2} != "aaa"}`}, false, false},
		{"in", args{map[string]interface{}{"col2": "aa"},
			`{col2} in ["aa", "bb"]}`}, true, false},
		{"not in", args{map[string]interface{}{"col2": "cc"},
			`{col2} not in ["aa", "bb"]`}, true, false},
		{"=~", args{map[string]interface{}{"col2": "aaaccaaa"},
			`{col2} =~ "cc"`}, true, false},
		{"!~", args{map[string]interface{}{"col2": "aaaccaaa"},
			`{col2} !~ "dd"`}, true, false},
		{">", args{map[string]interface{}{"col2": 11},
			`{col2} > 10`}, true, false},
		{"<", args{map[string]interface{}{"col2": -1},
			`{col2} < 0`}, true, false},
		{"and", args{map[string]interface{}{"col1": 111, "col2": "aaa"},
			`{col1} == 111 and {col2} == "aaa"`}, true, false},
		{"and", args{map[string]interface{}{"col1": 111, "col2": "aaa"},
			`{col1} == 111 and {col2} == "ccc"`}, false, false},
		{"or", args{map[string]interface{}{"col1": 111, "col2": "aaa"},
			`{col1} == 111 or {col2} == "ccc"`}, true, false},
		{"and or", args{map[string]interface{}{"col1": 111, "col2": "aaa", "col3": 333.21},
			`{col1} == 111 or ({col2} == "ccc" and {col3} in [333.21])`}, true, false},
		{"error", args{map[string]interface{}{"col1": "aaabbbb"},
			`{col1} =～ "aaa"`}, false, true},
		{"error", args{map[string]interface{}{"col4": "哈哈哈", "col3": "支付宝111"},
			`{col4} < 0 and {col3} =~ "支付宝"`}, false, true},
		{"error ==~", args{map[string]interface{}{"col2": "aaaccaaa"},
			`{col2} ==~ "cc"`}, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Validate(tt.args.data, tt.args.conditionExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
