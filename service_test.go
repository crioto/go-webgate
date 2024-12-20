package main

import (
	"net/http"
	"testing"

	"github.com/urfave/cli"
)

func TestConfig_ReadConfig(t *testing.T) {
	type fields struct {
		REST     *RESTConfig
		Services []*EndpointCategory
	}
	type args struct {
		ConfigFilepath string
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
			c := &Config{
				REST:     tt.fields.REST,
				Services: tt.fields.Services,
			}
			if err := c.ReadConfig(tt.args.ConfigFilepath); (err != nil) != tt.wantErr {
				t.Errorf("Config.ReadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildCache(t *testing.T) {
	type args struct {
		services []*EndpointCategory
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BuildCache(tt.args.services); (err != nil) != tt.wantErr {
				t.Errorf("BuildCache() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRunService(t *testing.T) {
	type args struct {
		c *cli.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RunService(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("RunService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandle(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Handle(tt.args.w, tt.args.r)
		})
	}
}
