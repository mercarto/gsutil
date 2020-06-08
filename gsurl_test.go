package gsutil

import (
	"fmt"
	"testing"
)

func TestParseGSURL(t *testing.T) {
	var v interface{} = float32(2387.23)
	var v3 = float64(v.(float32))
	fmt.Println(v, v3)
	type args struct {
		value string
	}
	tests := []struct {
		name       string
		args       args
		wantBucket string
		wantPath   string
		wantErr    bool
	}{
		{
			name:       "valid url with root path",
			args:       args{"gs://test-bucket/foo.txt"},
			wantBucket: "test-bucket",
			wantPath:   "foo.txt",
			wantErr:    false,
		},
		{
			name:       "valid url with nested path",
			args:       args{"gs://test-bucket/foo/bar/baz.txt"},
			wantBucket: "test-bucket",
			wantPath:   "foo/bar/baz.txt",
			wantErr:    false,
		},
		{
			name:       "valid bucket url (with no path)",
			args:       args{"gs://test-bucket"},
			wantBucket: "test-bucket",
			wantPath:   "",
			wantErr:    false,
		},
		{
			name:       "valid url with invalid scheme",
			args:       args{"s3://test-bucket/foo/bar/baz.txt"},
			wantBucket: "",
			wantPath:   "",
			wantErr:    true,
		},
		{
			name:       "invalid url",
			args:       args{"not_a+url"},
			wantBucket: "",
			wantPath:   "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBucket, gotPath, err := ParseGSURL(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseGSURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotBucket != tt.wantBucket {
				t.Errorf("ParseGSURL() gotBucket = %v, want %v", gotBucket, tt.wantBucket)
			}
			if gotPath != tt.wantPath {
				t.Errorf("ParseGSURL() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
		})
	}
}
