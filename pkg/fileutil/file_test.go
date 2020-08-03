package fileutil

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestExists(t *testing.T) {
	cases := map[string]bool{
		"file.go":   true,
		"file.go.1": false,

		"file_test.go":   true,
		"file_test.go.1": false,

		"../fileutil":   true,
		"../fileutil_1": false,

		"../../": true,
	}

	for fpath, expect := range cases {
		if Exists(fpath) == expect {
			t.Logf("pass, fpath=%s, expect= %v", fpath, expect)
		} else {
			t.Errorf("fail, fpath=%s, expect= %v", fpath, expect)
		}

	}

}

func TestFilesWithSuffixInDir(t *testing.T) {
	here, _ := filepath.Abs("./testdata")
	files := map[string]string{
		"json1": here + "/file1.json",
		"json2": here + "/file2.json",
		"toml1": here + "/file1.toml",
		"toml2": here + "/file2.toml",
		"conf1": here + "/file1.conf",
		"conf2": here + "/file2.conf",
	}

	tests := []struct {
		name     string
		dir      string
		suffixes []string
		wantRes  []string
		wantErr  bool
	}{
		{"1", here, []string{".json"}, []string{files["json1"], files["json2"]}, false},
		{"2", here, []string{".toml"}, []string{files["toml1"], files["toml2"]}, false},
		{"3", here, []string{".conf"}, []string{files["conf1"], files["conf2"]}, false},
		{"4", here, []string{"conf"}, []string{files["conf1"], files["conf2"]}, false},
		{"5", here, []string{"conf", ".toml"}, []string{files["conf1"], files["toml1"], files["conf2"], files["toml2"]}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := FilesWithSuffixInDir(tt.dir, tt.suffixes)
			if (err != nil) != tt.wantErr {
				t.Errorf(" %v FilesWithSuffixInDir() error = %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("%v FilesWithSuffixInDir() = `%v`, want `%v`", tt.name, gotRes, tt.wantRes)
			}
		})
	}
}
