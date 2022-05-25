package morm

import "testing"

func Test_getTableName(t *testing.T) {
	type args struct {
		d interface{}
	}
	type Case struct {
		Field1    string ``
		TableName string `morm:"colName=case1"`
	}

	type Case2 struct {
		Field string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{d: &Case{Field1: "field"}}, want: "case1"},
		{args: args{d: &Case2{Field: "field"}}, want: "case2s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTableName(tt.args.d); got != tt.want {
				t.Errorf("getTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}
