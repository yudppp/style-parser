package style

import (
	"testing"
)

func TestParseFromString(t *testing.T) {
	parser := Parser{}
	styleStr := `color: #666;`
	result := parser.ParseFromString(styleStr)
	color, ok := result["color"]
	if !ok || color != "#666" {
		t.Errorf("Not parsed to  %v", styleStr)
	}

	styleStr = `color:#666;`
	result = parser.ParseFromString(styleStr)
	color, ok = result["color"]
	if !ok || color != "#666" {
		t.Errorf("Not parsed to  %v", styleStr)
	}

	styleStr = `color:#666`
	result = parser.ParseFromString(styleStr)
	color, ok = result["color"]
	if !ok || color != "#666" {
		t.Errorf("Not parsed to  %v", styleStr)
	}

	styleStr = `color:blue; line-height:1.5;`
	result = parser.ParseFromString(styleStr)
	color, ok = result["color"]
	if !ok || color != "blue" {
		t.Errorf("Not parsed to  %v", styleStr)
	}
	lineHeight, ok := result["line-height"]
	if !ok || lineHeight != "1.5" {
		t.Errorf("Not parsed to  %v", styleStr)
	}

	styleStr = `
	color:blue;
	line-height:1.5;
	`
	result = parser.ParseFromString(styleStr)
	color, ok = result["color"]
	if !ok || color != "blue" {
		t.Errorf("Not parsed to  %v", styleStr)
	}
	lineHeight, ok = result["line-height"]
	if !ok || lineHeight != "1.5" {
		t.Errorf("Not parsed to  %v", styleStr)
	}
}
