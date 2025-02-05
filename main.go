package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func convertSVGToTSX(inputPath, outputPath string) error {
	err := os.MkdirAll(outputPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	commentRegex := regexp.MustCompile(`<!--.*?-->`)
	xmlDeclRegex := regexp.MustCompile(`(?i)^\s*<\?xml.*\?>`)

	err = filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".svg" {
			fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
			pascalCaseFileName := toPascalCase(fileName)
			tsxFileName := pascalCaseFileName + "Icon.tsx"
			tsxFilePath := filepath.Join(outputPath, tsxFileName)

			tsxFile, err := os.Create(tsxFilePath)
			if err != nil {
				return err
			}
			defer tsxFile.Close()

			svgContent, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			
			svgContentStr := convertSVGAttributes(string(svgContent))

			cleanedSVGContent := commentRegex.ReplaceAllString(svgContentStr, "")

			cleanedSVGContent = xmlDeclRegex.ReplaceAllString(cleanedSVGContent, "")

			cleanedSVGContent = strings.TrimSpace(cleanedSVGContent)

			_, err = fmt.Fprintf(tsxFile, "import { SVGProps } from 'react';\n\n")
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(tsxFile, "export default function %sIcon(props: SVGProps<SVGSVGElement>) {\n", pascalCaseFileName)
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(tsxFile, "  return (\n    ")
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(tsxFile, "%s\n", cleanedSVGContent)
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(tsxFile, "  );\n}")
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func toPascalCase(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	words := strings.Fields(s)

	titleCaser := cases.Title(language.English)

	var result string
	for _, word := range words {
		result += titleCaser.String(word)
	}
	return result
}

func convertSVGAttributes(content string) string {
	replacements := map[string]string{
		`class=`:             `className=`,
		`stroke-width=`:      `strokeWidth=`,
		`stroke-linecap=`:    `strokeLinecap=`,
		`stroke-linejoin=`:   `strokeLinejoin=`,
		`fill-rule=`:         `fillRule=`,
		`clip-rule=`:         `clipRule=`,
		`xmlns:xlink=`:       `xmlnsXlink=`,
		`xlink:href=`:        `xlinkHref=`,
		`font-family=`:       `fontFamily=`,
		`font-size=`:         `fontSize=`,
		`text-anchor=`:       `textAnchor=`,
		`dominant-baseline=`: `dominantBaseline=`,
		`stop-color=`:        `stopColor=`,
		`stop-opacity=`:      `stopOpacity=`,
	}

	for oldAttr, newAttr := range replacements {
		content = strings.ReplaceAll(content, oldAttr, newAttr)
	}

	return content
}

func main() {
	inputPath := flag.String("input", "./input", "Path to the input directory => svg files")
	outputPath := flag.String("output", "./output", "Path to the output directory => tsx files")

	flag.Parse()

	err := convertSVGToTSX(*inputPath, *outputPath)
	if err != nil {
		log.Fatalf("Error converting SVG to TSX: %v", err)
	}
	fmt.Println("SVG to TSX conversion completed successfully")
}
