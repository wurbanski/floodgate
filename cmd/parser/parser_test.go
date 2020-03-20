package parser

import (
	"reflect"
	"testing"

	"github.com/google/go-jsonnet"
)

func TestParser_loadFile(t *testing.T) {
	librariesPath := []string{"testdata/testlib.libsonnet"}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "load jsonnet file",
			args: args{
				filePath: "testdata/test.jsonnet",
			},
			want: map[string]interface{}{
				"name":     "test",
				"variable": false,
			},
			wantErr: false,
		},
		{
			name: "load json file",
			args: args{
				filePath: "testdata/test.json",
			},
			want: map[string]interface{}{
				"name":     "test",
				"variable": false,
			},
			wantErr: false,
		},
		{
			name: "load yaml file",
			args: args{
				filePath: "testdata/test.yaml",
			},
			want: map[string]interface{}{
				"name":     "test",
				"variable": false,
			},
			wantErr: false,
		},
		{
			name: "fail to load file with invalid extension",
			args: args{
				filePath: "testdata/test.invalid_ext",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "fail to load nonexistent file",
			args: args{
				filePath: "testdata/nonexistent.file",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := CreateParser(librariesPath)
			got, err := p.loadFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.loadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.loadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_loadDirectory(t *testing.T) {
	type fields struct {
		VM        *jsonnet.VM
		Resources SpinnakerResources
	}
	type args struct {
		entrypoint string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []map[string]interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				VM:        tt.fields.VM,
				Resources: tt.fields.Resources,
			}
			got, err := p.loadDirectory(tt.args.entrypoint)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.loadDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.loadDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_LoadObjectsFromDirectories(t *testing.T) {
	type fields struct {
		VM        *jsonnet.VM
		Resources SpinnakerResources
	}
	type args struct {
		directories []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				VM:        tt.fields.VM,
				Resources: tt.fields.Resources,
			}
			if err := p.LoadObjectsFromDirectories(tt.args.directories); (err != nil) != tt.wantErr {
				t.Errorf("Parser.LoadObjectsFromDirectories() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_readObjects(t *testing.T) {
	type fields struct {
		VM        *jsonnet.VM
		Resources SpinnakerResources
	}
	type args struct {
		objects []map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				VM:        tt.fields.VM,
				Resources: tt.fields.Resources,
			}
			if err := p.readObjects(tt.args.objects); (err != nil) != tt.wantErr {
				t.Errorf("Parser.readObjects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
