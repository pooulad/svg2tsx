package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"my-icon_name", "MyIconName"},
		{"hello_world", "HelloWorld"},
		{"test-case", "TestCase"},
		{"singleword", "Singleword"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := toPascalCase(tt.input)
			if result != tt.expected {
				t.Errorf("For input '%s', expected '%s' but got '%s'", tt.input, tt.expected, result)
			}
		})
	}
}

func TestConvertSVGToTSX(t *testing.T) {
	inputDir := "./test_input"
	outputDir := "./test_output"

	os.RemoveAll(inputDir)
	os.RemoveAll(outputDir)

	err := os.MkdirAll(inputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create input directory: %v", err)
	}

	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create output directory: %v", err)
	}

	svgContent := "<svg>...</svg>"
	err = os.WriteFile(filepath.Join(inputDir, "sample.svg"), []byte(svgContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create SVG file: %v", err)
	}

	err = convertSVGToTSX(inputDir, outputDir)
	if err != nil {
		t.Fatalf("Error converting SVG to TSX: %v", err)
	}

	outputFile := filepath.Join(outputDir, "SampleIcon.tsx")
	_, err = os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output TSX file: %v", err)
	}

	os.RemoveAll(inputDir)
	os.RemoveAll(outputDir)
}
