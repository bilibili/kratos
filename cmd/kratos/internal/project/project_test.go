package project

import (
	"testing"
)

func Test_getProjectPlaceDir(t *testing.T) {
	type args struct {
		projectName      string
		fallbackPlaceDir string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"absWindows", args{projectName: "c:\\kratos\\awesome\\go\\demo", fallbackPlaceDir: ""}, "c:\\kratos\\awesome\\go"},
		{"windowsHome", args{projectName: "~\\kratos\\awesome\\go\\demo", fallbackPlaceDir: ""}, "C:\\Users\\czyt\\kratos\\awesome\\go"},
		{"absLinux", args{projectName: "/home/kratos/awesome/go/demo", fallbackPlaceDir: ""}, "/home/kratos/awesome/go"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getProjectPlaceDir(tt.args.projectName, tt.args.fallbackPlaceDir); got != tt.want {
				t.Errorf("getProjectPlaceDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
