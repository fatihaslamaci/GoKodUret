package main

import (

	"testing"
)

func TestSelectTable(t *testing.T) {
	type args struct {
		value string
		tmpl  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Kullanici",
			args:args{ value:sc, tmpl: SelectIdTmp },
			want: HedefSelectIdTmp,
		},

		{
			name: "KullaniciCreate",
			args:args{ value:sc, tmpl: TableTmp },
			want: HedefKullanicitable,
		},

		{
			name: "KullaniciStruct",
			args:args{ value:sc, tmpl: KullanicistructTmp },
			want: HedefKullanicistruct,
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SelectTable(tt.args.value, tt.args.tmpl); got != tt.want {
				t.Errorf("SelectTable() = %v, want %v", got, tt.want)
			}
		})
	}
}
